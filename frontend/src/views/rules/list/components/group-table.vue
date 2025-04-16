<template>
  <a-card class="general-card" :title="title">
    <a-row style="margin-bottom: 16px">
      <a-col :span="12">
        <a-space>
          <a-button type="primary" @click="handleClick"
            ><template #icon> <icon-plus /> </template>新增告警规则</a-button
          >
        </a-space>
      </a-col>
      <a-col :span="12">
        <a-space style="float: right">
          <a-popconfirm
            content="是否确定要重新加载Prometheus配置?"
            @ok="prometheusReload"
          >
            <a-button type="primary">重新加载Prometheus配置</a-button>
          </a-popconfirm>
        </a-space>
      </a-col>
    </a-row>
    <a-table
      row-key="id"
      :loading="loading"
      :columns="columns"
      :data="renderData"
      :bordered="false"
      :pagination="false"
    >
      <template #operations="{ record }">
        <a-button type="text" @click="editRuleClick(record)">编辑</a-button>
        <a-button type="text" @click="showDetailDrawer(record)">详情</a-button>
        <a-popconfirm
          content="您的操作将同步到远程yaml文件"
          @ok="deleteRule(record)"
        >
          <a-button type="text">删除</a-button>
        </a-popconfirm>
      </template>
      <template #wd="{ record }">
        {{ secToFormat(record.duration) }}
      </template>
      <template #status="{ record }">
        <a-tag v-if="record.health === 'ok'" color="green">{{
          record.health
        }}</a-tag>
        <a-tag v-else color="red">{{ record.health }}</a-tag>
      </template>
    </a-table>
  </a-card>
  <DetailDrawer ref="detailDrawerRef"></DetailDrawer>
  <RuleDrawer ref="ruleDrawer" @refresh="fetchData()"></RuleDrawer>
</template>

<script lang="ts" setup>
  import {
    AlertRuleRecode,
    deleteAlertRule,
    queryAlertRuleList,
    reloadAlertRuleSource,
  } from '@/api/alert-rule';
  import useLoading from '@/hooks/loading';
  import { TableColumnData } from '@arco-design/web-vue';
  import { computed, ref, watch } from 'vue';
  import { secToFormat } from '@/utils/times';
  import DetailDrawer from './detail-drawer.vue';
  import RuleDrawer from './rule-drawer.vue';

  const props = defineProps({
    selectKey: {
      type: Number,
    },
    selectGroup: {
      type: String,
    },
  });
  const title = ref();
  const detailDrawerRef = ref();
  const ruleDrawer = ref();
  const { loading, setLoading } = useLoading(true);
  const renderData = ref<AlertRuleRecode[]>([]);
  const columns = computed<TableColumnData[]>(() => [
    {
      title: '数据源类型',
      dataIndex: 'sourceType',
    },
    {
      title: '标题',
      dataIndex: 'name',
    },
    {
      title: '状态',
      dataIndex: 'health',
      slotName: 'status',
    },
    {
      title: '评估等待时间',
      dataIndex: 'duration',
      slotName: 'wd',
    },
    {
      title: '更新时间',
      dataIndex: 'updatedAt',
    },
    {
      title: '操作',
      dataIndex: 'operations',
      slotName: 'operations',
    },
  ]);

  const fetchData = async () => {
    title.value = `${props.selectGroup} 告警规则列表`;
    setLoading(true);
    if (props.selectKey === undefined) {
      return;
    }
    try {
      const { data } = await queryAlertRuleList(props.selectKey as number);
      renderData.value = data;
    } catch (err) {
      // you can report use errorHandler or other
    } finally {
      setLoading(false);
    }
  };

  watch(
    props,
    () => {
      fetchData();
    },
    {
      deep: true,
      immediate: true,
    }
  );

  const showDetailDrawer = (val: AlertRuleRecode) => {
    detailDrawerRef.value.show(val);
  };

  const handleClick = () => {
    ruleDrawer.value.show(props.selectKey, props.selectGroup);
  };

  const deleteRule = async (record: AlertRuleRecode) => {
    setLoading(true);
    try {
      await deleteAlertRule(record.id);
    } catch (err) {
      // you can report use errorHandler or other
    } finally {
      fetchData();
      setLoading(false);
    }
  };

  const editRuleClick = (record: AlertRuleRecode) => {
    try {
      ruleDrawer.value.show(props.selectKey, props.selectGroup, record);
    } catch (error) {
      //
    }
  };

  const prometheusReload = async () => {
    setLoading(true);
    try {
      await reloadAlertRuleSource(props.selectKey as number);
    } catch (err) {
      // you can report use errorHandler or other
    } finally {
      fetchData();
      setLoading(false);
    }
  };
</script>

<style scoped lang="less"></style>
