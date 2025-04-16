<template>
  <a-button type="text" @click.stop="open">
    <template #icon>
      <icon-sort size="15" />
    </template>
  </a-button>
  <a-drawer
    :width="600"
    :visible="visible"
    unmount-on-close
    @ok="handleOk"
    @cancel="handleCancel"
  >
    <template #title> 告警通知策略排序 </template>
    <div style="margin-bottom: 20px">
      <a-alert>您可以拖动告警通知策略，进行排序。</a-alert>
    </div>
    <a-table
      :columns="columns"
      :data="data"
      :draggable="{}"
      :pagination="false"
      @change="handleChange"
    >
      <template #sort="{ rowIndex }">
        {{ rowIndex + 1 }}
      </template>
    </a-table>
  </a-drawer>
</template>

<script lang="ts" setup>
  import { reactive, ref } from 'vue';
  import {
    AlertStrategyRecord,
    AlertStrategySort,
    editStrategySort,
  } from '@/api/alert-strategy';
  import { Message, TableData } from '@arco-design/web-vue';

  const visible = ref(false);
  interface Props {
    alertStrategyList: AlertStrategyRecord[];
  }
  const props = defineProps<Props>();
  const emits = defineEmits(['fetchListData']);
  const sortData = ref<AlertStrategySort[]>([]);
  const data = ref<TableData[]>(props.alertStrategyList);
  const columns = reactive([
    {
      title: '策略名称',
      dataIndex: 'strategyName',
    },
    {
      title: '策略排序',
      slotName: 'sort',
    },
    {
      title: '最后修改用户',
      dataIndex: 'nickname',
    },
    {
      title: '最后更新时间',
      dataIndex: 'updatedAt',
    },
  ]);

  const open = () => {
    visible.value = true;
  };
  const handleOk = async () => {
    if (sortData.value.length > 0) {
      try {
        await editStrategySort(sortData.value);
        Message.success('策略排序成功');
        visible.value = false;
        emits('fetchListData');
      } catch {
        //
      }
    } else {
      visible.value = false;
    }
  };
  const handleCancel = () => {
    visible.value = false;
  };

  const handleChange = (tableData: TableData[]) => {
    data.value = tableData;
    sortData.value = [];
    for (let i = 0; i < tableData.length; i += 1) {
      sortData.value.push({
        id: tableData[i].id,
        weight: i + 1,
      });
    }
  };
</script>
