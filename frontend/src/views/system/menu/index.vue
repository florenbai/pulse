<template>
  <div class="container">
    <Breadcrumb
      :items="[{ label: 'menu.system' }, { label: 'menu.system.menu' }]"
    />
    <a-card class="general-card" title="菜单列表">
      <a-row style="margin-bottom: 16px">
        <a-col :span="12">
          <a-space>
            <a-button type="primary" @click="handleClick"
              ><template #icon> <icon-plus /> </template>新增菜单</a-button
            >
          </a-space>
        </a-col>
      </a-row>
      <a-table
        ref="artable"
        row-key="id"
        :loading="loading"
        :columns="columns"
        :data="renderData"
        :bordered="false"
        :size="size"
        :pagination="false"
        :draggable="{ type: 'handle', width: 40 }"
        @change="handleChange"
      >
        <template #icon="{ record }">
          <Icon v-if="record.icon !== ''" :icon="record.icon" :size="20"></Icon>
        </template>
        <template #front="{ record }">
          <a-switch v-model="record.enabled" @change="changeMenu(record)">
            <template #checked> 启用 </template>
            <template #unchecked> 禁用 </template>
          </a-switch>
        </template>
        <template #operations="{ record }">
          <a-space>
            <a-button type="text" size="small" @click="editMenu(record)">
              编辑
            </a-button>
            <a-popconfirm
              content="删除菜单后，它的所有子菜单也将一并删除，请确认要删除?"
              @ok="onDelete(record.id)"
            >
              <a-button type="text" size="small"> 删除 </a-button>
            </a-popconfirm>
          </a-space>
        </template>
      </a-table>
    </a-card>
    <MenuForm ref="menuRef" @refresh="fetchData"></MenuForm>
  </div>
</template>

<script lang="ts" setup>
  import { computed, ref, inject, nextTick } from 'vue';
  import useLoading from '@/hooks/loading';
  import type { TableColumnData } from '@arco-design/web-vue/es/table/interface';
  import { Message } from '@arco-design/web-vue';
  import {
    changeMenuSort,
    changeMenuStatus,
    deleteMenu,
    getMenuList,
    MenuRecode,
  } from '@/api/menu';
  import { RouteRecordNormalized } from 'vue-router';
  import { Icon } from '@/components/Icon';
  import MenuForm from './components/menu-form.vue';

  const artable = ref();
  type SizeProps = 'mini' | 'small' | 'medium' | 'large';
  const size = ref<SizeProps>('medium');
  const reload = inject('reload') as any;

  const { loading, setLoading } = useLoading(true);
  const renderData = ref<RouteRecordNormalized[]>([]);

  const menuRef = ref();
  const columns = computed<TableColumnData[]>(() => [
    {
      title: '菜单名称',
      dataIndex: 'name',
    },
    {
      title: '图标',
      dataIndex: 'icon',
      slotName: 'icon',
    },
    {
      title: '状态',
      dataIndex: 'enable',
      slotName: 'front',
    },
    {
      title: '操作',
      dataIndex: 'operations',
      slotName: 'operations',
    },
  ]);

  const fetchData = async () => {
    setLoading(true);
    try {
      const { data } = await getMenuList();
      renderData.value = data;
      nextTick(() => {
        artable.value.expandAll();
      });
    } catch (err) {
      // you can report use errorHandler or other
    } finally {
      setLoading(false);
    }
  };

  const getSortData = (_data: any) => {
    const a: number[] = [];
    if (_data.length > 0) {
      for (let i = 0; i < _data.length; i += 1) {
        a.push(_data[i].id);
        if (_data[i].children !== undefined && _data[i].children.length > 0) {
          a.push(...getSortData(_data[i].children));
        }
      }
    }
    return a;
  };

  const handleChange = async (_data: any) => {
    let ids: number[] = [];
    ids = getSortData(_data);
    setLoading(true);
    try {
      await changeMenuSort(ids);
      Message.success({
        content: '排序完成',
        duration: 5 * 1000,
      });
      renderData.value = _data;
    } catch (err) {
      // you can report use errorHandler or other
    } finally {
      setLoading(false);
    }
  };

  const handleClick = () => {
    menuRef.value.show();
  };

  const editMenu = (item: MenuRecode) => {
    menuRef.value.show(item);
  };

  const changeMenu = async (item: MenuRecode) => {
    setLoading(true);
    try {
      await changeMenuStatus(item.id);
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
  fetchData();
  const onDelete = async (id: number) => {
    setLoading(true);
    try {
      await deleteMenu(id);
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
