<template>
  <a-row style="margin-bottom: 16px; padding-top: 20px">
    <a-col
      :span="12"
      style="display: flex; align-items: center; justify-content: start"
    >
      <a-space>
        <a-input
          v-model:model-value="listParams.alertTitle"
          :style="{ width: '320px' }"
          placeholder="请输入告警标题"
          allow-clear
        >
          <template #prepend> 告警标题 </template>
        </a-input>
        <faultSelect :value="4" @change-value="changeFault"></faultSelect>
        <timeSelect @change-value="changeTimeRange"> </timeSelect>
        <AlertLevelSelect @change-value="changeAlertLevel"></AlertLevelSelect>
      </a-space>
    </a-col>
    <a-col
      :span="12"
      style="display: flex; align-items: center; justify-content: end"
    >
      <a-space>
        <a-popconfirm content="请确认您要认领选中的告警?" @ok="claim">
          <a-button type="primary"> 认领 </a-button>
        </a-popconfirm>

        <closeModal
          type="primary"
          :selected="selectedKeys as []"
          @refresh="completeClosed"
        />
        <refreshSelect @refresh="changeInterval"></refreshSelect>
      </a-space>
    </a-col>
  </a-row>
  <a-table
    v-model:selectedKeys="selectedKeys"
    :columns="columns"
    :data="renderData"
    :row-selection="rowSelection"
    :loading="loading"
    :pagination="pagination"
    row-key="id"
    scrollbar
    :scroll="{ y: 700 }"
    @page-change="onPageChange"
  >
    <template
      #name-filter="{
        filterValue,
        setFilterValue,
        handleFilterConfirm,
        handleFilterReset,
      }"
    >
      <div class="custom-filter">
        <a-space direction="vertical">
          <a-input
            :model-value="filterValue[0]"
            @input="(value) => setFilterValue([value])"
          />
          <div class="custom-filter-footer">
            <a-button @click="handleFilterConfirm">Confirm</a-button>
            <a-button @click="handleFilterReset">Reset</a-button>
          </div>
        </a-space>
      </div>
    </template>
    <template #tags="{ record }">
      <TypographyTag :tags="record.tags"></TypographyTag>
    </template>
    <template #sourceName="{ record }">
      <a-tooltip :content="record.sourceName">
        <icon-font :type="record.icon" :size="24" />
      </a-tooltip>
    </template>
    <template #level="{ record }">
      <a-tag :color="record.color">{{ record.levelName }}</a-tag>
    </template>
    <template #progress="{ record }">
      <a-tag :color="PROGRESS_COLOR_MAP[record.progress]">{{
        PROGRESS_MAP[record.progress]
      }}</a-tag>
    </template>
    <template #duration="{ record }">
      {{ duration(record) }}
    </template>
  </a-table>
  <audio ref="audioRef" :autoplay="false">
    <source src="/src/assets/alert.mp3" type="audio/mpeg" />
  </audio>
</template>

<script lang="ts" setup>
  import { reactive, h, ref, watch, onBeforeUnmount } from 'vue';
  import { IconSearch } from '@arco-design/web-vue/es/icon';
  import { Message, TableRowSelection } from '@arco-design/web-vue';
  import { Pagination } from '@/types/global';
  import useLoading from '@/hooks/loading';
  import {
    AlertEventRecode,
    WorkspaceAlertEventParams,
    claimAlertEvent,
    queryAllAlertEventList,
  } from '@/api/alert-event';
  import faultSelect from '@/components/select/workspace/fault-select.vue';
  import timeSelect from '@/components/select/workspace/time-select.vue';
  import refreshSelect from '@/components/select/workspace/refresh-select.vue';
  import { PROGRESS_MAP, PROGRESS_COLOR_MAP } from '@/types/workspace';
  import AlertLevelSelect from '@/components/select/workspace/alert-level-select.vue';
  import closeModal from '@/views/alert/workspace/components/detail/list/components/close-modal.vue';
  import TypographyTag from './typography-tag.vue';

  const intervalId = ref<number | null>(null);
  const abortController = new AbortController();
  const interval = ref(10);
  const rowSelection = ref<TableRowSelection>({
    type: 'checkbox',
    showCheckedAll: true,
  });

  const selectedKeys = ref<(string | number)[] | undefined>([]);
  const renderData = ref<AlertEventRecode[]>([]);
  const maxId = ref(0);
  const isInit = ref(true);
  const audioRef = ref();
  const { loading, setLoading } = useLoading(true);

  const columns = [
    {
      title: '告警标题',
      dataIndex: 'alertTitle',
      filterable: {
        filter: (value: any, record: any) => record.name.includes(value),
        slotName: 'name-filter',
        icon: () => h(IconSearch),
      },
      ellipsis: true,
      tooltip: { position: 'left' },
      width: 150,
    },
    {
      title: '标签',
      slotName: 'tags',
      width: 250,
    },
    {
      title: '数据源',
      slotName: 'sourceName',
      width: 30,
    },
    {
      title: '告警等级',
      slotName: 'level',
      width: 40,
    },
    {
      title: '处理进度',
      slotName: 'progress',
      width: 40,
    },
    {
      title: '告警次数',
      dataIndex: 'notifyCurNumber',
      width: 40,
    },
    {
      title: '首次告警时间',
      dataIndex: 'firstTriggerTime',
      width: 65,
    },
    {
      title: '最新告警时间',
      dataIndex: 'triggerTime',
      width: 65,
    },
    {
      title: '持续时间',
      slotName: 'duration',
      width: 50,
    },
  ];
  const basePagination: Pagination = {
    page: 1,
    pageSize: 20,
    total: 1,
  };

  const pagination = reactive({
    ...basePagination,
  });

  const listParams: WorkspaceAlertEventParams = reactive({
    ...pagination,
    timeRange: '6month',
    progress: 4,
  });

  const fetchData = async (params: WorkspaceAlertEventParams) => {
    setLoading(true);
    try {
      const { data } = await queryAllAlertEventList(params);
      if (data.list.length > 0) {
        if (maxId.value === 0) {
          if (!isInit.value) {
            audioRef.value.play();
          }
          maxId.value = data.list[0].id;
        } else {
          if (maxId.value < data.list[0].id) {
            audioRef.value.play();
          }
          maxId.value = data.list[0].id;
        }
      }
      renderData.value = data.list;
      pagination.page = params.page;
      pagination.total = data.total;
      if (isInit.value) {
        isInit.value = false;
      }
    } catch (err) {
      // you can report use errorHandler or other
    } finally {
      setLoading(false);
    }
  };

  const onPageChange = (page: number) => {
    listParams.page = page;
    fetchData(listParams);
  };
  const changeFault = (value: number) => {
    if (value === 0) {
      listParams.progress = undefined;
    } else {
      listParams.progress = value;
    }
  };

  const changeTimeRange = (value: string) => {
    listParams.timeRange = value;
  };

  const start = () => {
    if (intervalId.value !== null) {
      return;
    }
    intervalId.value = window.setInterval(() => {
      fetchData(listParams);
    }, interval.value * 1000);
  };

  const stop = () => {
    if (intervalId.value === null) {
      return;
    }
    clearInterval(intervalId.value);
    intervalId.value = null;
  };

  const changeInterval = (newInterval: number) => {
    if (intervalId.value === null) {
      return;
    }
    stop();
    if (newInterval <= 0) {
      return;
    }
    interval.value = newInterval;
    start();
  };

  fetchData(listParams);
  start();

  const duration = (record: AlertEventRecode) => {
    let startTime = new Date();
    let endTime = new Date();
    if (record.isRecovered) {
      startTime = new Date(record.firstTriggerTime);
      endTime = new Date(record.recoverTime);
    } else {
      startTime = new Date(record.firstTriggerTime);
      endTime = new Date();
    }

    const usedTime = endTime.getTime() - startTime.getTime();
    const days = Math.floor(usedTime / (24 * 3600 * 1000));
    const leavel = usedTime % (24 * 3600 * 1000);
    const hours = Math.floor(leavel / (3600 * 1000));
    const leavel2 = leavel % (3600 * 1000);
    const minutes = Math.floor(leavel2 / (60 * 1000));
    if (days > 0) {
      return `${days}天${hours}时`;
    }
    if (days === 0 && hours > 0) {
      return `${hours}时${minutes}分`;
    }
    return `${minutes}分`;
  };

  watch(listParams, () => {
    if (listParams.alertTitle === '') {
      listParams.alertTitle = undefined;
    }
    stop();
    fetchData(listParams);
    start();
  });

  const completeClosed = () => {
    selectedKeys.value = [];
    fetchData(listParams);
  };

  const claim = async () => {
    if (selectedKeys.value === undefined || selectedKeys.value.length === 0) {
      Message.error('请您至少选中一个要认领的告警');
      return;
    }
    try {
      await claimAlertEvent(selectedKeys.value);
      Message.success('领取成功');
      completeClosed();
    } catch (err) {
      // you can report use errorHandler or other
    } finally {
      setLoading(false);
    }
  };

  const changeAlertLevel = (value: string) => {
    listParams.alertLevel = value;
  };

  onBeforeUnmount(() => {
    abortController.abort();
    if (intervalId.value !== null) {
      clearInterval(intervalId.value);
    }
  });
</script>

<style>
  .custom-filter {
    padding: 20px;
    background: var(--color-bg-5);
    border: 1px solid var(--color-neutral-3);
    border-radius: var(--border-radius-medium);
    box-shadow: 0 2px 5px rgb(0 0 0 / 10%);
  }

  .custom-filter-footer {
    display: flex;
    justify-content: space-between;
  }
</style>
