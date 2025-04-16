<template>
  <div class="container">
    <Breadcrumb
      :items="[{ label: 'menu.system' }, { label: 'menu.system.strategy' }]"
    />
    <a-card class="general-card" title="全局告警策略列表">
      <a-row style="margin-bottom: 16px">
        <a-col :span="12">
          <a-space>
            <a-button type="primary" @click="handleClick"
              ><template #icon> <icon-plus /> </template>新增告警策略</a-button
            >
          </a-space>
        </a-col>
      </a-row>
      <a-table
        row-key="id"
        :loading="loading"
        :pagination="pagination"
        :columns="(cloneColumns as TableColumnData[])"
        :data="renderData"
        :bordered="false"
        :size="size"
        @page-change="onPageChange"
      >
        <template #index="{ rowIndex }">
          {{ rowIndex + 1 + (pagination.page - 1) * pagination.pageSize }}
        </template>
        <template #operations="{ record }">
          <a-space>
            <a-button type="text" size="small" @click="editTeam(record)">
              编辑
            </a-button>
            <a-popconfirm
              content="是否确认要删除告警策略?"
              @ok="onDelete(record.id)"
            >
              <a-button type="text" size="small"> 删除 </a-button>
            </a-popconfirm>
          </a-space>
        </template>
      </a-table>
    </a-card>
    <strategyDrawer
      ref="strategyDrawerRef"
      @reload="reloadChange"
    ></strategyDrawer>
  </div>
</template>

<script lang="ts" setup>
  import { computed, ref, reactive, watch, inject } from 'vue';
  import useLoading from '@/hooks/loading';
  import { Pagination } from '@/types/global';
  import type { TableColumnData } from '@arco-design/web-vue/es/table/interface';
  import cloneDeep from 'lodash/cloneDeep';
  import { Message } from '@arco-design/web-vue';
  import {
    querySystemStrategyList,
    SystemStrategyRecord,
    SystemStrategyParams,
    deleteSystemStrategy,
  } from '@/api/system-strategy';
  import strategyDrawer from './components/strategy-drawer.vue';

  type SizeProps = 'mini' | 'small' | 'medium' | 'large';
  const size = ref<SizeProps>('medium');
  const reload = inject('reload') as any;
  type Column = TableColumnData & { checked?: true };

  const generateFormModel = () => {
    return {
      strategyName: '',
    };
  };
  const { loading, setLoading } = useLoading(true);
  const renderData = ref<SystemStrategyRecord[]>([]);
  const formModel = ref(generateFormModel());
  const cloneColumns = ref<Column[]>([]);
  const showColumns = ref<Column[]>([]);
  const strategyDrawerRef = ref();

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
      title: '告警策略名称',
      dataIndex: 'strategyName',
    },
    {
      title: '告警策略描述',
      dataIndex: 'strategyDesc',
    },
    {
      title: '更新人',
      dataIndex: 'nickname',
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

  const fetchData = async (
    params: SystemStrategyParams = { page: 1, pageSize: 20 }
  ) => {
    setLoading(true);
    try {
      const { data } = await querySystemStrategyList(params);
      renderData.value = data.list;
      pagination.page = params.page;
      pagination.total = data.total;
    } catch (err) {
      // you can report use errorHandler or other
    } finally {
      setLoading(false);
    }
  };

  const search = () => {
    fetchData({
      ...basePagination,
      ...formModel.value,
    } as unknown as SystemStrategyParams);
  };
  const onPageChange = (page: number) => {
    fetchData({ ...basePagination, page });
  };

  fetchData();
  const reset = () => {
    formModel.value = generateFormModel();
  };

  watch(
    () => columns.value,
    (val) => {
      cloneColumns.value = cloneDeep(val);
      cloneColumns.value.forEach((item, index) => {
        item.checked = true;
      });
      showColumns.value = cloneDeep(cloneColumns.value);
    },
    { deep: true, immediate: true }
  );
  const handleClick = () => {
    strategyDrawerRef.value.show();
  };

  const editTeam = (record: SystemStrategyRecord) => {
    strategyDrawerRef.value.show(record);
  };

  const onDelete = async (id: number) => {
    setLoading(true);
    try {
      await deleteSystemStrategy(id);
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

  const reloadChange = () => {
    reload();
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
@/api/team
