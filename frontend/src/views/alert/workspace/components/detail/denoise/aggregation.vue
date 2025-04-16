<template>
  <div :style="{ display: 'flex' }">
    <a-card title="告警聚合" :bordered="false" :style="{ width: '100%' }">
      <div>
        <a-grid
          :cols="{ xs: 1, sm: 2, md: 3, lg: 3, xl: 4, xxl: 4 }"
          :col-gap="12"
          :row-gap="16"
          class="grid"
        >
          <a-grid-item
            v-for="(item, index) of alertAggregationList"
            :key="index"
            class="grid-item"
          >
            <div class="card-wrap" @click="editAggregation(item)">
              <a-space align="start">
                <a-card-meta>
                  <template #title>
                    <a-typography-text style="margin-right: 10px">
                      {{ item.aggregationName }}
                    </a-typography-text>

                    <a-tag
                      v-if="item.status === 'enabled'"
                      size="small"
                      color="green"
                    >
                      <template #icon>
                        <icon-check-circle-fill />
                      </template>
                      <span>启用</span>
                    </a-tag>
                    <a-tag
                      v-else-if="item.status === 'disabled'"
                      size="small"
                      color="red"
                    >
                      <template #icon>
                        <icon-stop />
                      </template>
                      <span>禁用</span>
                    </a-tag>

                    <span class="right-link-hover">
                      <a-button
                        type="text"
                        @click.stop="aggregationSwitch(item, index)"
                      >
                        <template #icon>
                          <icon-stop
                            v-if="item.status === 'enabled'"
                            size="15"
                          />
                          <icon-check-circle
                            v-else-if="item.status === 'disabled'"
                            size="15"
                          />
                        </template>
                      </a-button>
                      <a-button
                        v-if="item.status === 'disabled'"
                        type="text"
                        @click.stop="(event) => deleteAggregation(item, event)"
                      >
                        <template #icon>
                          <icon-delete size="15" />
                        </template>
                      </a-button>
                    </span>
                  </template>

                  <template #description>
                    <div class="description">
                      {{ item.aggregationDesc }}
                    </div>
                    <div class="info">
                      {{ item.nickname }} {{ item.updatedAt }}
                    </div>
                  </template>
                </a-card-meta>
              </a-space>
            </div>
          </a-grid-item>
          <a-grid-item
            class="grid-item"
            style="border-style: dashed"
            @click="handleClick"
          >
            <div style="text-align: center; margin: auto">
              <icon-plus style="margin-right: 1rem" />
              <span>新增告警聚合</span></div
            >
          </a-grid-item>
        </a-grid>
      </div>
    </a-card>
    <aggregationDrawer
      :id="props.id"
      ref="aggregationDrawerRef"
      @refresh="fetchListData"
    ></aggregationDrawer>
  </div>
</template>

<script setup lang="ts">
  import { onMounted, ref, h } from 'vue';
  import {
    queryAllAlertAggregation,
    AlertAggregationRecord,
    deleteAggregationRecord,
    changeAggregationStatus,
  } from '@/api/alert-denoise';
  import { Message, Notification } from '@arco-design/web-vue';
  import aggregationDrawer from './aggregation-drawer.vue';

  export interface Props {
    id: number | undefined;
  }
  const props = defineProps<Props>();
  const aggregationDrawerRef = ref();
  const alertAggregationList = ref<AlertAggregationRecord[]>([]);

  const fetchListData = async () => {
    try {
      const res = await queryAllAlertAggregation(props.id as number);
      alertAggregationList.value = res.data;
    } catch (err) {
      //
    }
  };

  const handleClick = () => {
    aggregationDrawerRef.value.show();
  };

  const aggregationSwitch = async (
    item: AlertAggregationRecord,
    index: number
  ) => {
    if (item.status === 'enabled') {
      item.status = 'disabled';
      alertAggregationList.value[index] = item;
    } else {
      item.status = 'enabled';
      alertAggregationList.value[index] = item;
    }

    try {
      await changeAggregationStatus(item.id);
    } catch (err) {
      //
    }
  };

  const deleteAggregation = (item: AlertAggregationRecord, event: any) => {
    const x = `${event.clientX - 340}px`;
    const y = `${event.clientY - 100}px`;
    const id = `${Date.now()}`;
    Notification.info({
      id,
      title: '策略删除提醒',
      position: 'topLeft',
      duration: 0,
      content: '请确认是否要删除该策略',
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
                Notification.remove(id);
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
                  await deleteAggregationRecord(item.id);
                  Message.success('策略删除成功');
                  Notification.remove(id);
                  fetchListData();
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

  const editAggregation = (item: AlertAggregationRecord) => {
    aggregationDrawerRef.value.show(item);
  };

  onMounted(() => {
    fetchListData();
  });
</script>

<style scoped>
  .grid {
    gap: 0.75rem !important;
  }
  .grid-item {
    position: relative !important;
    display: flex;
    height: 100% !important;
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
  .description {
    min-height: 40px;
    -webkit-line-clamp: 3;
    display: -webkit-inline-box;
    -webkit-box-orient: vertical;
    overflow: hidden;
    text-overflow: ellipsis;
    width: 85%;
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
</style>
