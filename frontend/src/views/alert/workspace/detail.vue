<template>
  <div class="container">
    <Breadcrumb
      :items="[
        { label: 'menu.alert' },
        { label: 'menu.alert.workspace', path: '/owl/alert/workspace' },
        {
          label: formData.workspaceName,
          locale: false,
        },
      ]"
    />
    <a-space direction="vertical" :size="12" fill>
      <a-space direction="vertical" :size="16" fill>
        <div class="space-unit">
          <PublicOpinion :id="id" />
        </div>
        <div>
          <a-card class="general-card">
            <a-tabs class="tabs-custom" type="line">
              <a-tab-pane key="1">
                <template #title> 故障列表 </template>
                <incidentList :id="id"></incidentList>
              </a-tab-pane>
              <!--
              <a-tab-pane key="2">
                <template #title> 告警源配置 </template>
                <alertSourceConfiguration :id="id"></alertSourceConfiguration>
              </a-tab-pane>
              -->
              <a-tab-pane key="2">
                <template #title> 告警通知策略 </template>
                <strategyList :id="id"></strategyList>
              </a-tab-pane>
              <a-tab-pane key="3">
                <template #title> 降噪配置 </template>
                <denoiseLayout :id="id"></denoiseLayout>
              </a-tab-pane>
            </a-tabs>
          </a-card>
        </div>
      </a-space>
    </a-space>
  </div>
</template>

<script lang="ts" setup>
  import { useRoute } from 'vue-router';
  import { ref, watch } from 'vue';
  import { BaseInfoRequest, queryWorkspaceBase } from '@/api/workspace';
  import PublicOpinion from './components/detail/statistic/public-opinion.vue';
  import incidentList from './components/detail/incident-list.vue';
  import strategyList from './components/detail/strategy/alert-strategy.vue';
  import denoiseLayout from './components/detail/denoise/denoise-layout.vue';

  const id = ref<number>();
  const route = useRoute();
  const formData = ref<BaseInfoRequest>({
    workspaceName: '',
    workspaceDesc: '',
    teams: [],
    strategy: 0,
  });
  const getWorkspace = async (did: number) => {
    try {
      const res = await queryWorkspaceBase(did);
      formData.value = { ...res.data };
    } catch {
      //
    }
  };

  watch(
    route,
    () => {
      if (route.params.id) {
        id.value = Number(route.params?.id);
        getWorkspace(id.value);
      }
    },
    {
      deep: true,
      immediate: true,
    }
  );
</script>

<style scoped lang="less">
  .container {
    padding: 0 20px 20px 20px;
    :deep(.arco-list-content) {
      overflow-x: hidden;
    }

    :deep(.arco-card-meta-title) {
      font-size: 14px;
    }
  }
  .tabs-custom {
    padding-top: 20px;
  }
  .tabs-custom ::before {
    height: 0px !important;
  }
</style>
