<template>
  <div class="container">
    <Breadcrumb :items="[{ label: 'menu.template.list' }]" />
    <a-row :gutter="20" align="stretch">
      <a-col :span="24">
        <a-card class="general-card" title="告警模板列表">
          <a-row justify="space-between">
            <a-col :span="24">
              <a-space>
                <a-input-search
                  placeholder="请输入告警模板名称"
                  style="
                    width: 240px;
                    position: absolute;
                    top: 60px;
                    left: 20px;
                  "
                />
                <a-button
                  type="primary"
                  style="position: absolute; top: 60px; right: 20px"
                  @click="handleClick"
                  ><template #icon> <icon-plus /> </template
                  >新增告警模板</a-button
                >
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
                    <CardWrap
                      :template="item"
                      :loading="loading"
                      @open-drawer="openDrawer"
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
                    </CardWrap>
                  </a-col>
                </a-row>
              </div>
            </a-col>
          </a-row>
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
    <templateDrawer
      ref="templateDrawerRef"
      @refresh="fetchData"
    ></templateDrawer>
  </div>
</template>

<script lang="ts" setup>
  import { reactive, ref } from 'vue';
  import { Pagination } from '@/types/global';
  import useLoading from '@/hooks/loading';
  import {
    AlertTemplateRecode,
    AlertTemplateParams,
    queryAlertTemplateList,
  } from '@/api/alert-template';
  import CardWrap from './components/card-wrap.vue';
  import templateDrawer from './components/template-drawer.vue';

  const { loading, setLoading } = useLoading(true);
  const renderData = ref<AlertTemplateRecode[]>([]);
  const templateDrawerRef = ref();
  const basePagination: Pagination = {
    page: 1,
    pageSize: 20,
    total: 1,
  };
  const pagination = reactive({
    ...basePagination,
  });
  const fetchData = async (
    params: AlertTemplateParams = { page: 1, pageSize: 20 }
  ) => {
    setLoading(true);
    try {
      const { data } = await queryAlertTemplateList(params);
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
    templateDrawerRef.value.show();
  };

  const openDrawer = (template: AlertTemplateRecode) => {
    templateDrawerRef.value.show(template);
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

  .list-wrap {
    margin-top: 30px;
  }
</style>
