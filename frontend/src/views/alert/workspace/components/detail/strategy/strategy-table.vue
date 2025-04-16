<template>
  <a-table
    :columns="columns"
    :data="alertStrategyList"
    :draggable="{}"
  ></a-table>
</template>

<script setup lang="ts">
  import { ref, onMounted, reactive } from 'vue';
  import {
    StrategyItem,
    AlertStrategyRequest,
    AlertStrategyRecord,
    submitStrategyForm,
    queryAllAlertStrategy,
    changeStrategyStatus,
    editStrategyForm,
    deleteStrategyRecord,
  } from '@/api/alert-strategy';

  export interface Props {
    id: number | undefined;
  }
  const props = defineProps<Props>();
  const alertStrategyList = ref<AlertStrategyRecord[]>([]);
  const columns = reactive([
    {
      title: '告警策略名称',
      dataIndex: 'strategyName',
    },
    {
      title: '告警策略描述',
      dataIndex: 'strategyDesc',
    },
    {
      title: '最后修改时间',
      dataIndex: 'updatedAt',
    },
    {
      title: '修改人',
      dataIndex: 'nickname',
    },
  ]);
  const fetchListData = async () => {
    try {
      const res = await queryAllAlertStrategy(props.id as number);
      alertStrategyList.value = res.data;
    } catch (err) {
      //
    }
  };
  onMounted(() => {
    fetchListData();
  });
</script>
