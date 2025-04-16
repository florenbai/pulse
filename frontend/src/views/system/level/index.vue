<template>
  <div class="container">
    <Breadcrumb
      :items="[{ label: 'menu.system' }, { label: 'menu.system.level' }]"
    />
    <a-card class="general-card" title="告警等级列表">
      <a-row style="margin-bottom: 16px">
        <a-col :span="12">
          <a-space>
            <a-button type="primary" @click="handleClick"
              ><template #icon> <icon-plus /> </template>新增告警等级</a-button
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
        :size="size"
        :pagination="false"
      >
        <template #color="{ record }">
          <a-tag :color="record.color">{{ record.color }}</a-tag>
        </template>
        <template #isDefault="{ record }">
          <a-tag v-if="record.isDefault" color="green">
            <icon-check-circle size="15" />
          </a-tag>
          <a-tag v-if="!record.isDefault" color="red">
            <icon-close-circle size="15" />
          </a-tag>
        </template>
        <template #operations="{ record }">
          <a-space>
            <a-button type="text" size="small" @click="editLevel(record)">
              编辑
            </a-button>
            <a-popconfirm
              v-if="!record.isDefault"
              content="告警等级删除后，已配置的等级映射将失效。
              是否确认要删除告警等级?"
              @ok="onDelete(record.id)"
            >
              <a-button type="text" size="small"> 删除 </a-button>
            </a-popconfirm>
          </a-space>
        </template>
      </a-table>
    </a-card>
    <LevelFormDrawer ref="levelFormRef" @refresh="fetchData" />
  </div>
</template>

<script lang="ts" setup>
  import { computed, ref, inject } from 'vue';
  import useLoading from '@/hooks/loading';
  import type { TableColumnData } from '@arco-design/web-vue/es/table/interface';
  import { Message } from '@arco-design/web-vue';
  import {
    AlertLevelRecode,
    deleteLevel,
    queryAllAlertLevel,
  } from '@/api/alert-level';
  import LevelFormDrawer from './components/level-form-drawer.vue';

  type SizeProps = 'mini' | 'small' | 'medium' | 'large';
  const size = ref<SizeProps>('medium');
  const reload = inject('reload') as any;

  const { loading, setLoading } = useLoading(true);
  const renderData = ref<AlertLevelRecode[]>([]);

  const levelFormRef = ref();
  const columns = computed<TableColumnData[]>(() => [
    {
      title: '编号',
      dataIndex: 'id',
    },
    {
      title: '告警等级',
      dataIndex: 'levelName',
    },
    {
      title: '告警描述',
      dataIndex: 'levelDesc',
    },
    {
      title: '颜色',
      dataIndex: 'color',
      slotName: 'color',
    },
    {
      title: '系统内置',
      dataIndex: 'isDefault',
      slotName: 'isDefault',
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
      const { data } = await queryAllAlertLevel();
      renderData.value = data;
    } catch (err) {
      // you can report use errorHandler or other
    } finally {
      setLoading(false);
    }
  };

  const handleClick = () => {
    levelFormRef.value.show();
  };

  const editLevel = (item: AlertLevelRecode) => {
    levelFormRef.value.show(item);
  };
  fetchData();
  const onDelete = async (id: number) => {
    setLoading(true);
    try {
      await deleteLevel(id);
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
