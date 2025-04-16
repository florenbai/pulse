<template>
  <a-drawer
    :width="1000"
    :visible="visible"
    unmount-on-close
    :ok-loading="loading"
    @ok="handleOk"
    @cancel="handleCancel"
  >
    <template #title> agent列表 </template>
    <a-table
      row-key="id"
      :loading="loading"
      :columns="columns"
      :data="agentData"
      :bordered="false"
      :pagination="false"
    ></a-table>
  </a-drawer>
</template>

<script lang="ts" setup>
  import { ref, computed } from 'vue';
  import { TableColumnData } from '@arco-design/web-vue';
  import {
    Agent,
    AlertRuleSourceRecode,
    queryAllAlertRuleSourceAgent,
  } from '@/api/alert-rule';
  import useLoading from '@/hooks/loading';

  const { loading, setLoading } = useLoading(false);

  const visible = ref(false);
  const agentData = ref<Agent[]>([]);

  const columns = computed<TableColumnData[]>(() => [
    {
      title: '客户端IP',
      dataIndex: 'clientIp',
    },
    {
      title: '心跳时间',
      dataIndex: 'heartbeatTime',
    },
  ]);
  const handleCancel = () => {
    visible.value = false;
  };
  const handleOk = () => {
    visible.value = false;
  };

  const getAgentList = async (id: number) => {
    try {
      setLoading(true);
      const res = await queryAllAlertRuleSourceAgent(id);
      agentData.value = res.data;
    } catch (err) {
      //
    } finally {
      setLoading(false);
    }
  };
  const show = (item: AlertRuleSourceRecode) => {
    getAgentList(item.id);
    visible.value = true;
  };

  defineExpose({
    show,
  });
</script>

<style scoped></style>
