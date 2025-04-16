<template>
  <div class="container">
    <Breadcrumb
      :items="[
        { label: 'menu.alert' },
        { label: 'menu.source.integration', path: '/owl/alert/integration' },
        {
          label: sourceData?.name,
          locale: false,
        },
      ]"
    />
    <a-card :bordered="false">
      <a-space direction="vertical" fill style="height: 100%">
        <a-descriptions title="基础信息" layout="horizontal">
          <a-descriptions-item label="集成名称">
            {{ sourceData?.name }}
          </a-descriptions-item>
          <a-descriptions-item label="接入地址">
            {{ axios.defaults.baseURL }}/api/v1/alert-event/push
          </a-descriptions-item>
          <a-descriptions-item label="告警等级字段">
            {{ sourceData?.levelField }}
          </a-descriptions-item>
          <a-descriptions-item label="集成描述">
            {{ sourceData?.description }}
          </a-descriptions-item>
          <a-descriptions-item label="token">
            {{ sourceData?.token }}
          </a-descriptions-item>
          <a-descriptions-item label="集成状态">
            <a-tag :color="sourceData?.disabled ? 'gray' : 'green'">{{
              sourceData?.disabled ? '禁用' : '启用'
            }}</a-tag>
          </a-descriptions-item>
        </a-descriptions>
      </a-space>
    </a-card>
    <tagRewrite :integration="id"></tagRewrite>
    <a-card :bordered="false">
      <a-descriptions title="告警路由" layout="horizontal"></a-descriptions>
      <a-grid
        :cols="{ xs: 1, sm: 2, md: 3, lg: 3, xl: 4, xxl: 4 }"
        :col-gap="12"
        :row-gap="16"
        class="grid"
      >
        <draggable
          :list="routerData"
          handle=".handle"
          item-key="id"
          style="display: contents"
          :component-data="{ tag: 'ul', name: 'flip-list', type: 'transition' }"
          @change="dragMoved"
        >
          <template #item="{ element }">
            <a-grid-item
              class="grid-item"
              style="border-style: dashed"
              @click="editRouter(element)"
            >
              <span class="right-link-hover">
                <a-button type="text" class="handle" @click.stop="">
                  <template #icon>
                    <icon-drag-dot-vertical size="15" />
                  </template>
                </a-button>

                <a-button
                  type="text"
                  @click.stop="
                    (event) => {
                      deleteRouter(element, event);
                    }
                  "
                >
                  <template #icon>
                    <icon-delete size="15" />
                  </template>
                </a-button>
              </span>
              <div class="card-wrap">
                <a-space align="start" direction="vertical" fill>
                  <div
                    v-if="
                      element.filters != undefined && element.filters.length > 0
                    "
                  >
                    <div style="margin-bottom: 10px">根据以下条件</div>
                    <a-row v-for="(f, i) in element.filters" :key="i">
                      <a-space fill direction="vertical">
                        <a-row
                          v-for="(r, k) in f.values"
                          :key="k"
                          justify="space-between"
                        >
                          <a-tag bordered>{{ r.tag }}</a-tag>
                          <a-tag bordered color="orangered">{{
                            r.operation === 'IN' ? '包含' : '不包含'
                          }}</a-tag>
                          <a-tag
                            v-for="(v, w) in r.value"
                            :key="w"
                            color="blue"
                            >{{ v }}</a-tag
                          >
                        </a-row>
                      </a-space>
                      <a-divider v-if="i < element.filters.length - 1" />
                    </a-row>
                  </div>
                  <div v-else>无需限定条件</div>
                  <div style="margin-top: 10px"
                    >将告警发送到
                    <a-tag
                      v-for="(v, k) in element.workspaceNames"
                      :key="k"
                      color="blue"
                      >{{ v }}</a-tag
                    >
                    工作区，并且
                    <a-tag v-if="element.next === 1">继续匹配</a-tag>
                    <a-tag v-else>退出匹配</a-tag>
                  </div>
                </a-space>
              </div>
            </a-grid-item>
          </template>
        </draggable>

        <a-grid-item
          class="grid-item"
          style="border-style: dashed"
          @click="handleClick"
        >
          <div style="text-align: center; margin: auto">
            <icon-plus style="margin-right: 1rem" /> <span>告警路由</span></div
          >
        </a-grid-item>
      </a-grid>
    </a-card>
    <div class="actions">
      <a-space>
        <a-popconfirm
          v-if="!sourceData?.disabled"
          content="禁用告警集成后，该集成将不能收到告警，请确认您的操作"
          @ok="changeIntegrationStatus"
          ><a-button> 禁用 </a-button></a-popconfirm
        >
        <a-popconfirm
          v-else
          content="启用告警集成后，该集成将收到告警，请确认您的操作"
          @ok="changeIntegrationStatus"
          ><a-button> 启用 </a-button></a-popconfirm
        >

        <a-button type="primary" @click="onSubmitClick"> 编辑 </a-button>
      </a-space>
    </div>
    <AlertRouter ref="alertRouterRef" @refresh="getAllRouters"></AlertRouter>
    <BaseInfo ref="baseInfoRef" @refresh="getSourceDetail"></BaseInfo>
  </div>
</template>

<script lang="ts" setup>
  import {
    changeRouterSort,
    deleteIntegrationRouterRecord,
    IntegrationRouter,
    queryIntegrationAllRouters,
    queryIntegrationDetail,
    SystemIntegrationRecode,
    changeIntegrationStatusById,
  } from '@/api/integration';
  import { h, ref, watch } from 'vue';
  import { useRoute } from 'vue-router';
  import axios from 'axios';
  import draggable from 'vuedraggable';
  import { Message, Notification } from '@arco-design/web-vue';
  import AlertRouter from './components/alert-router.vue';
  import tagRewrite from './components/tag-rewrite.vue';
  import BaseInfo from './components/base-info.vue';

  const id = ref<number>();
  const route = useRoute();
  const alertRouterRef = ref();
  const baseInfoRef = ref();
  const sourceData = ref<SystemIntegrationRecode>();
  const routerData = ref<IntegrationRouter[]>([]);

  const getSourceDetail = async (did: number) => {
    try {
      const res = await queryIntegrationDetail(did);
      sourceData.value = res.data;
    } catch {
      //
    }
  };

  const getAllRouters = async (did: number) => {
    try {
      const res = await queryIntegrationAllRouters(did);
      routerData.value = res.data as IntegrationRouter[];
    } catch {
      //
    }
  };

  const changeIntegrationStatus = async () => {
    try {
      await changeIntegrationStatusById(sourceData.value?.id as number);
      getSourceDetail(sourceData.value?.id as number);
    } catch {
      //
    }
  };

  watch(
    route,
    () => {
      if (route.params.id) {
        id.value = Number(route.params?.id);
        getSourceDetail(id.value);
        getAllRouters(id.value);
      }
    },
    {
      deep: true,
      immediate: true,
    }
  );

  const handleClick = () => {
    alertRouterRef.value.show(id.value);
  };

  const editRouter = (item: IntegrationRouter) => {
    alertRouterRef.value.show(id.value, item);
  };

  const onSubmitClick = () => {
    baseInfoRef.value.show(sourceData.value);
  };

  const dragMoved = async () => {
    const ids = [];
    for (let i = 0; i < routerData.value.length; i += 1) {
      ids.push(routerData.value[i].id);
    }
    try {
      await changeRouterSort(ids);
      Message.success('路由排序更新成功');
    } catch {
      //
    }
  };

  const deleteRouter = (item: IntegrationRouter, event: any) => {
    const x = `${event.clientX - 340}px`;
    const y = `${event.clientY - 100}px`;
    const now = `${Date.now()}`;
    Notification.info({
      id: now,
      title: '路由删除提醒',
      position: 'topLeft',
      duration: 0,
      content: '请确认是否要删除该路由',
      style: { left: x, top: y },
      footer: () =>
        h('Space', [
          h(
            'Button',
            {
              type: 'secondary',
              size: 'small',
              style: {
                color: '#var(--color-text-1)',
                backgroundColor: 'transparent',
                border: '1px solid var(--color-border-3)',
                height: '28px',
                padding: '0 15px',
                borderRadius: 'var(--border-radius-small)',
                fontSize: '14px',
                marginRight: '8px',
              },
              onClick: () => {
                Notification.remove(now);
              },
            },
            '取消'
          ),
          h(
            'Button',
            {
              type: 'primary',
              size: 'small',
              style: {
                color: '#fff',
                backgroundColor: 'rgb(var(--primary-6))',
                border: '1px solid transparent',
                height: '28px',
                padding: '0 15px',
                borderRadius: 'var(--border-radius-small)',
                fontSize: '14px',
              },
              onClick: async () => {
                try {
                  await deleteIntegrationRouterRecord(item.id);
                  Message.success('路由删除成功');
                  Notification.remove(now);
                  getAllRouters(id.value as number);
                } catch (err) {
                  //
                }
              },
            },
            '确认'
          ),
        ]),
    });
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

  .grid {
    gap: 0.75rem !important;
  }
  .grid-item {
    position: relative !important;
    display: flex;
    height: 100% !important;
    min-height: 150px;
    cursor: pointer !important;
    flex-direction: column !important;
    align-items: center !important;
    overflow: hidden !important;
    border-radius: 0.25rem !important;
    border-width: 1px !important;
    border-style: solid !important;
    padding-top: 1rem !important;
    padding-bottom: 1rem !important;
    transition-property: color, background-color, border-color,
      text-decoration-color, fill, stroke, opacity, box-shadow, transform,
      filter, backdrop-filter, -webkit-text-decoration-color,
      -webkit-backdrop-filter !important;
    transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1) !important;
    transition-duration: 0.15s !important;
    border-color: #e5e7eb !important;
    background-color: rgb(var(--fc-fill-3) / var(--tw-bg-opacity)) !important;
  }

  .grid-item:hover {
    border-color: #4f7fe1 !important;
  }

  .grid-item:hover .right-link-hover {
    display: inline-flex;
  }

  .grid-item:hover .right-link {
    display: none;
  }

  .grid-item-checked {
    right: 0 !important;
    top: -5px !important;
    position: absolute !important;
    color: rgb(var(--primary-6));
  }

  .grid-item-checked-bg {
    rotate: 45deg !important;
    opacity: 1 !important;
    background-color: var(--color-primary-light-1);
    width: 26px !important;
    height: 26px !important;
    top: -11px !important;
    right: -11px !important;
    position: absolute !important;
  }

  .text-tip {
    font-size: 12px;
    margin-top: 5px;
  }

  .delete-icon {
    position: absolute;
    right: 10px;
    top: 50%;
    transform: translate(-50%, -50%);
    display: none;
  }

  .grid-item:hover .delete-icon {
    display: flex;
  }

  .title {
    font-size: 14px;
    font-weight: 600;
  }

  .source-input {
    width: 500px;
  }

  .grid-demo {
    margin-bottom: 10px;
  }

  .description {
    font-size: 13px;
    margin-top: 5px;
    color: darkgray;
  }

  .card-wrap {
    width: -webkit-fill-available;
    padding-left: 20px;
  }

  .right-link {
    position: absolute;
    right: 0;
    margin-right: 20px;
  }

  .right-link-hover {
    position: absolute;
    right: 0;
    margin-right: 20px;
    display: none;
  }
  .actions {
    position: fixed;
    left: 0;
    right: 0;
    bottom: 0;
    height: 60px;
    padding: 14px 20px 14px 0;
    background: var(--color-bg-2);
    text-align: right;
  }
</style>
