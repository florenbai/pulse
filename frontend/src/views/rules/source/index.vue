<template>
  <div class="container">
    <Breadcrumb :items="[{ label: 'menu.alert.rule.dataSource' }]" />
    <a-row :gutter="20" align="stretch">
      <a-col :span="24">
        <a-card class="general-card" title="数据源列表">
          <a-row style="margin-bottom: 16px">
            <a-col :span="12">
              <a-space>
                <a-button type="primary" @click="handleClick"
                  ><template #icon> <icon-plus /> </template
                  >新增数据源</a-button
                >
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
              <a-button type="text" @click="gotoAgent(record)"
                >agent列表</a-button
              >
              <a-button type="text" @click="editHandle(record)">编辑</a-button>
              <a-popconfirm
                content="数据源删除后，该数据源的规则也将删除，请确认要删除数据源?"
                @ok="deleteSource(record)"
              >
                <a-button type="text">删除</a-button>
              </a-popconfirm>
            </template>
          </a-table>
          <a-pagination
            v-if="
              pagination.total != undefined &&
              pagination.total - pagination.pageSize > 0
            "
            :total="(pagination.total as number)"
          />
        </a-card>
      </a-col>
    </a-row>
    <sourceDrawer ref="sourceDrawerRef" @refresh="fetchData"></sourceDrawer>
    <agentDrawer ref="agentDrawerRef" />
  </div>
</template>

<script lang="ts" setup>
  import { computed, reactive, ref } from 'vue';
  import { Pagination } from '@/types/global';
  import useLoading from '@/hooks/loading';
  import { TableColumnData } from '@arco-design/web-vue';
  import {
    AlertRuleSourceRecode,
    deleteAlertRuleSource,
    PageParams,
    queryAlertRuleSource,
  } from '@/api/alert-rule';
  import sourceDrawer from './components/source-drawer.vue';
  import agentDrawer from './components/agent-drawer.vue';

  const { loading, setLoading } = useLoading(true);
  const renderData = ref<AlertRuleSourceRecode[]>([]);
  const sourceDrawerRef = ref();
  const agentDrawerRef = ref();
  const columns = computed<TableColumnData[]>(() => [
    {
      title: '编号',
      dataIndex: 'id',
    },
    {
      title: '数据源',
      dataIndex: 'sourceName',
    },
    {
      title: '数据源类型',
      dataIndex: 'sourceType',
    },
    {
      title: '连接地址',
      dataIndex: 'address',
    },
    {
      title: '备注',
      dataIndex: 'mark',
    },
    {
      title: '数据源标识',
      dataIndex: 'sign',
    },
    {
      title: '操作',
      dataIndex: 'operations',
      slotName: 'operations',
    },
  ]);
  const basePagination: Pagination = {
    page: 1,
    pageSize: 20,
    total: 1,
  };
  const pagination = reactive({
    ...basePagination,
  });
  const fetchData = async (params: PageParams = { page: 1, pageSize: 20 }) => {
    setLoading(true);
    try {
      const { data } = await queryAlertRuleSource(params);
      renderData.value = data.list;
      pagination.page = params.page;
      pagination.total = data.total;
    } catch (err) {
      // you can report use errorHandler or other
    } finally {
      setLoading(false);
    }
  };
  fetchData();
  const handleClick = () => {
    sourceDrawerRef.value.show();
  };

  const editHandle = (item: AlertRuleSourceRecode) => {
    sourceDrawerRef.value.show(item);
  };

  const gotoAgent = (item: AlertRuleSourceRecode) => {
    agentDrawerRef.value.show(item);
  };

  const deleteSource = async (item: AlertRuleSourceRecode) => {
    setLoading(true);
    try {
      await deleteAlertRuleSource(item.id);
      fetchData();
    } catch (err) {
      // you can report use errorHandler or other
    } finally {
      setLoading(false);
    }
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

  .layout {
    display: flex;

    &-left-side {
      flex-basis: 300px;
    }

    &-content {
      flex: 1;
      padding: 0 16px;
    }

    &-right-side {
      flex-basis: 280px;
    }
  }
</style>
