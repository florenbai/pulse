<template>
  <div class="container">
    <Breadcrumb
      :items="[{ label: 'menu.system' }, { label: 'menu.system.user' }]"
    />
    <a-card class="general-card" title="用户列表">
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
                <a-form-item field="nickname" label="用户名称">
                  <a-input
                    v-model="formModel.nickName"
                    placeholder="请输入用户名称"
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
              ><template #icon> <icon-plus /> </template>新增用户</a-button
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
            <a-button type="text" size="small" @click="editUser(record)">
              编辑
            </a-button>
            <a-popconfirm
              content="是否确认要删除用户?"
              @ok="onDelete(record.id)"
            >
              <a-button type="text" size="small"> 删除 </a-button>
            </a-popconfirm>
          </a-space>
        </template>
      </a-table>
    </a-card>
    <user-form-modal
      ref="userFormRef"
      @reload-change="reloadChange"
    ></user-form-modal>
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
    queryUserList,
    UserParams,
    deleteUser,
    UserRecord,
  } from '@/api/user';
  import userFormModal from './components/user-form-modal.vue';

  type SizeProps = 'mini' | 'small' | 'medium' | 'large';
  const size = ref<SizeProps>('medium');
  const reload = inject('reload') as any;
  type Column = TableColumnData & { checked?: true };

  const generateFormModel = () => {
    return {
      nickName: '',
    };
  };
  const { loading, setLoading } = useLoading(true);
  const renderData = ref<UserRecord[]>([]);
  const formModel = ref(generateFormModel());
  const cloneColumns = ref<Column[]>([]);
  const showColumns = ref<Column[]>([]);
  const userFormRef = ref();

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
      title: '用户名称',
      dataIndex: 'username',
    },
    {
      title: '显示名称',
      dataIndex: 'nickname',
    },
    {
      title: '用户邮箱',
      dataIndex: 'email',
    },
    {
      title: '用户手机',
      dataIndex: 'phone',
    },
    {
      title: '操作',
      dataIndex: 'operations',
      slotName: 'operations',
    },
  ]);

  const fetchData = async (
    params: UserParams = {
      page: 1,
      pageSize: 20,
      nickName: formModel.value.nickName,
    }
  ) => {
    setLoading(true);
    try {
      const { data } = await queryUserList(params);
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
    } as unknown as UserParams);
  };
  const onPageChange = (page: number) => {
    fetchData({ ...basePagination, ...formModel.value });
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

  const onDelete = async (id: number) => {
    setLoading(true);
    try {
      await deleteUser(id);
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
    userFormRef.value.show();
  };

  const editUser = (record: UserRecord) => {
    userFormRef.value.show(record);
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
