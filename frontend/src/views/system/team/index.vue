<template>
  <div class="container">
    <Breadcrumb
      :items="[{ label: 'menu.system' }, { label: 'menu.system.team' }]"
    />
    <a-card class="general-card" title="团队列表">
      <a-row>
        <a-col :flex="1">
          <a-form
            :model="formModel"
            :label-col-props="{ span: 6 }"
            :wrapper-col-props="{ span: 18 }"
            label-align="left"
          >
            <a-row :gutter="16">
              <a-col :span="8">
                <a-form-item field="teamName" label="团队名称">
                  <a-input
                    v-model="formModel.teamName"
                    placeholder="请输入团队名称"
                  />
                </a-form-item>
              </a-col>
            </a-row>
          </a-form>
        </a-col>
        <a-col :flex="'86px'" style="text-align: right">
          <a-space :size="18">
            <a-button type="primary" @click="search">
              <template #icon>
                <icon-search />
              </template>
              查询
            </a-button>
            <a-button @click="reset">
              <template #icon>
                <icon-refresh />
              </template>
              重置
            </a-button>
          </a-space>
        </a-col>
      </a-row>
      <a-divider style="margin-top: 0" />
      <a-row style="margin-bottom: 16px">
        <a-col :span="12">
          <a-space>
            <a-button type="primary" @click="handleClick"
              ><template #icon> <icon-plus /> </template>新增团队</a-button
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
        <template #status="{ record }">
          <a-switch
            v-model:model-value="record.status"
            @click="handleChangeIntercept(record.id)"
          >
            <template #checked> 启用 </template>
            <template #unchecked> 禁用 </template>
          </a-switch>
        </template>
        <template #operations="{ record }">
          <a-space>
            <a-button
              type="text"
              size="small"
              @click="showPermission(record.id)"
            >
              权限授权
            </a-button>
            <a-button type="text" size="small" @click="editTeam(record.id)">
              编辑
            </a-button>
            <a-popconfirm
              content="是否确认要删除团队?"
              @ok="onDelete(record.id)"
            >
              <a-button type="text" size="small"> 删除 </a-button>
            </a-popconfirm>
          </a-space>
        </template>
      </a-table>
    </a-card>
    <team-form-modal
      :id="teamId"
      :visible="modalVisible"
      @visible-change="visibleChange"
      @reload-change="reloadChange"
    ></team-form-modal>
    <PermissionModal ref="permissionRef"></PermissionModal>
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
    queryTeamList,
    TeamParams,
    deleteTeam,
    TeamRecord,
    changeTeamStatus,
  } from '@/api/team';
  import teamFormModal from './components/team-form-modal.vue';
  import PermissionModal from './components/permission-modal.vue';

  const modalVisible = ref(false);
  type SizeProps = 'mini' | 'small' | 'medium' | 'large';
  const size = ref<SizeProps>('medium');
  const reload = inject('reload') as any;
  type Column = TableColumnData & { checked?: true };

  const generateFormModel = () => {
    return {
      teamName: '',
    };
  };
  const { loading, setLoading } = useLoading(true);
  const renderData = ref<TeamRecord[]>([]);
  const formModel = ref(generateFormModel());
  const permissionRef = ref();
  const cloneColumns = ref<Column[]>([]);
  const showColumns = ref<Column[]>([]);
  const teamId = ref<number>(0);

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
      title: '团队名称',
      dataIndex: 'teamName',
    },
    {
      title: '团队描述',
      dataIndex: 'teamDesc',
    },
    {
      title: '更新人',
      dataIndex: 'nickname',
    },
    {
      title: '状态',
      dataIndex: 'status',
      slotName: 'status',
    },
    {
      title: '操作',
      dataIndex: 'operations',
      slotName: 'operations',
    },
  ]);

  const fetchData = async (params: TeamParams = { page: 1, pageSize: 20 }) => {
    setLoading(true);
    try {
      const { data } = await queryTeamList(params);
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
    } as unknown as TeamParams);
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
    modalVisible.value = true;
  };

  const visibleChange = () => {
    modalVisible.value = false;
  };

  const editTeam = (id: number) => {
    modalVisible.value = true;
    teamId.value = id;
  };

  const onDelete = async (id: number) => {
    setLoading(true);
    try {
      await deleteTeam(id);
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

  const handleChangeIntercept = async (id: number) => {
    setLoading(true);
    try {
      await changeTeamStatus(id);
    } catch (err) {
      //
    } finally {
      setLoading(false);
    }
  };

  const reloadChange = () => {
    reload();
  };

  const showPermission = (id: number) => {
    permissionRef.value.show(id);
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
