<template>
  <div class="container">
    <Breadcrumb :items="[{ label: 'menu.schedule' }]" />
    <a-card class="general-card" title="值班列表">
      <a-divider style="margin-top: 0" />
      <a-row style="margin-bottom: 16px">
        <a-col :span="12">
          <a-space>
            <a-button type="primary" @click="handleClick"
              ><template #icon> <icon-plus /> </template>新增值班</a-button
            >
          </a-space>
        </a-col>
      </a-row>
      <a-table
        row-key="id"
        :loading="loading"
        :pagination="pagination"
        :columns="columns"
        :data="renderData"
        :bordered="false"
        :size="size"
        @page-change="onPageChange"
      >
        <template #index="{ rowIndex }">
          {{ rowIndex + 1 + (pagination.page - 1) * pagination.pageSize }}
        </template>
        <template #period="{ record }">
          {{ record.period }} {{ PERIOD_MAP[record.periodType] }}
        </template>
        <template #timeRange="{ record }">
          {{ record.startRange }} - {{ record.endRange }}
        </template>
        <template #operations="{ record }">
          <a-space>
            <a-button
              type="text"
              size="small"
              @click="gotoScheduleDetail(record.id)"
            >
              排班
            </a-button>
            <a-popconfirm
              content="是否确认要删除值班?"
              @ok="onDelete(record.id)"
            >
              <a-button type="text" size="small"> 删除 </a-button>
            </a-popconfirm>
          </a-space>
        </template>
      </a-table>
    </a-card>
    <scheduleFormModal
      ref="scheduleFormRef"
      @reload-change="reloadChange"
    ></scheduleFormModal>
  </div>
</template>

<script lang="ts" setup>
  import { computed, ref, reactive, inject } from 'vue';
  import useLoading from '@/hooks/loading';
  import { Pagination } from '@/types/global';
  import type { TableColumnData } from '@arco-design/web-vue/es/table/interface';
  import { Message } from '@arco-design/web-vue';
  import {
    queryScheduleList,
    deleteSchedule,
    ScheduleParams,
    ScheduleRecord,
  } from '@/api/schedule';
  import PERIOD_MAP from '@/types/schedule';
  import router from '@/router';
  import scheduleFormModal from './components/schedule-form-modal.vue';

  type SizeProps = 'mini' | 'small' | 'medium' | 'large';
  const size = ref<SizeProps>('medium');
  const scheduleFormRef = ref();
  const reload = inject('reload') as any;

  const { loading, setLoading } = useLoading(true);
  const renderData = ref<ScheduleRecord[]>([]);

  const basePagination: Pagination = {
    page: 1,
    pageSize: 20,
  };
  const pagination = reactive({
    ...basePagination,
  });

  const columns = computed<TableColumnData[]>(() => [
    {
      title: '编号',
      dataIndex: 'id',
    },
    {
      title: '值班名称',
      dataIndex: 'scheduleName',
    },
    {
      title: '值班描述',
      dataIndex: 'scheduleDesc',
    },
    {
      title: '所属团队',
      dataIndex: 'teamName',
    },
    {
      title: '值班轮换周期',
      dataIndex: 'period',
      slotName: 'period',
    },
    {
      title: '值班时间范围',
      dataIndex: 'timeRange',
      slotName: 'timeRange',
    },
    {
      title: '操作',
      dataIndex: 'operations',
      slotName: 'operations',
    },
  ]);

  const fetchData = async (
    params: ScheduleParams = { page: 1, pageSize: 20 }
  ) => {
    setLoading(true);
    try {
      const { data } = await queryScheduleList(params);
      renderData.value = data.list;
      pagination.page = params.page;
      pagination.total = data.total;
    } catch (err) {
      // you can report use errorHandler or other
    } finally {
      setLoading(false);
    }
  };

  const onPageChange = (page: number) => {
    fetchData({ ...basePagination, page });
  };

  fetchData();

  const onDelete = async (id: number) => {
    setLoading(true);
    try {
      await deleteSchedule(id);
      Message.success({
        content: '操作成功',
        duration: 5 * 1000,
      });
      reload();
    } catch (err) {
      // you can report use errorHandler or other
    } finally {
      setLoading(false);
    }
  };

  const handleClick = () => {
    scheduleFormRef.value.show();
  };

  const reloadChange = () => {
    reload();
  };

  const gotoScheduleDetail = (id: number) => {
    router.push(`/schedule/detail/${id}`);
  };
</script>

<script lang="ts">
  export default {
    name: 'SearchTable',
  };
</script>

<style scoped lang="less">
  .container {
    padding: 0 20px 20px 20px;
  }
  :deep(.arco-table-th) {
    &:last-child {
      .arco-table-th-item-title {
        margin-left: 16px;
      }
    }
  }
  .action-icon {
    margin-left: 12px;
    cursor: pointer;
  }
  .active {
    color: #0960bd;
    background-color: #e3f4fc;
  }
  .setting {
    display: flex;
    align-items: center;
    width: 200px;
    .title {
      margin-left: 12px;
      cursor: pointer;
    }
  }
  .arco-btn.arco-btn-text {
    padding: 0;
    height: auto;
  }
</style>
