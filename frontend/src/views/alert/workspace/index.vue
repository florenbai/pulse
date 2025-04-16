<template>
  <div class="container">
    <Breadcrumb
      :items="[{ label: 'menu.alert' }, { label: 'menu.alert.workspace' }]"
    />
    <a-card class="general-card" title="工作区">
      <a-space>
        <a-row justify="space-between">
          <a-col :span="24">
            <a-input-search
              placeholder="请输入"
              style="width: 240px; position: absolute; top: 60px; left: 20px"
            />
          </a-col>
          <a-col :span="24">
            <a-button
              type="primary"
              style="position: absolute; top: 60px; right: 20px"
              @click="handleClick"
            >
              <template #icon> <icon-plus /> </template>
              创建工作区
            </a-button>
          </a-col>
        </a-row>
      </a-space>
      <div class="list-wrap">
        <a-row class="list-row" :gutter="24">
          <a-col
            v-for="item in renderData"
            :key="item.id"
            :xs="12"
            :sm="12"
            :md="12"
            :lg="6"
            :xl="6"
            :xxl="6"
            class="list-col"
            style="min-height: 162px"
          >
            <workspace-card
              :id="item.id"
              :loading="loading"
              :title="item.workspaceName"
              :description="item.workspaceDesc"
              :default-value="item.enable"
              :source-count="item.sourceCount"
              :pending="item.pending"
              :processing="item.processing"
              @refresh="fetchData"
            >
              <template #skeleton>
                <a-skeleton :animation="true">
                  <a-skeleton-line
                    :widths="['100%', '40%', '100%']"
                    :rows="3"
                  />
                  <a-skeleton-line :widths="['40%']" :rows="1" />
                </a-skeleton>
              </template>
            </workspace-card>
          </a-col>
        </a-row>
      </div>
      <a-pagination
        v-if="
          pagination.total != undefined &&
          pagination.total - pagination.pageSize > 0
        "
        :total="(pagination.total as number)"
      />
    </a-card>
    <createWorkspace ref="workspaceRef" @refresh="fetchData"></createWorkspace>
  </div>
</template>

<script lang="ts" setup>
  import { reactive, ref } from 'vue';
  import { Pagination } from '@/types/global';
  import useLoading from '@/hooks/loading';
  import {
    WorkspaceRecode,
    WorkspaceParams,
    queryWorkspaceList,
  } from '@/api/workspace';
  import createWorkspace from './components/index/create.vue';
  import workspaceCard from './components/workspace-card.vue';

  const { loading, setLoading } = useLoading(true);
  const renderData = ref<WorkspaceRecode[]>([]);
  const workspaceRef = ref();
  const basePagination: Pagination = {
    page: 1,
    pageSize: 20,
    total: 1,
  };
  const pagination = reactive({
    ...basePagination,
  });
  const fetchData = async (
    params: WorkspaceParams = { page: 1, pageSize: 20 }
  ) => {
    setLoading(true);
    try {
      const { data } = await queryWorkspaceList(params);
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
    workspaceRef.value.show();
  };
</script>

<script lang="ts">
  export default {
    name: 'Workspace',
  };
</script>

<style scoped lang="less">
  .container {
    padding: 0 20px 20px 20px;
    height: 100vh;
    overflow-y: auto;
    :deep(.arco-list-content) {
      overflow-x: hidden;
    }

    :deep(.arco-card-meta-title) {
      font-size: 14px;
    }
  }
  :deep(.arco-list-col) {
    display: flex;
    flex-direction: row;
    flex-wrap: wrap;
    justify-content: space-between;
  }

  .list-wrap {
    margin-top: 30px;
  }

  :deep(.arco-list-item) {
    width: 33%;
  }

  :deep(.block-title) {
    margin: 0 0 12px 0;
    font-size: 14px;
  }
  :deep(.list-wrap) {
    // min-height: 140px;
    .list-row {
      align-items: stretch;
      .list-col {
        margin-bottom: 16px;
      }
    }
    :deep(.arco-space) {
      width: 100%;
      .arco-space-item {
        &:last-child {
          flex: 1;
        }
      }
    }
  }
</style>
