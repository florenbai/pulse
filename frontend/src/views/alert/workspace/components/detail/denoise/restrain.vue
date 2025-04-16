<template>
  <div :style="{ display: 'flex' }">
    <a-card title="告警抑制" :bordered="false" :style="{ width: '100%' }">
      <div>
        <a-grid
          :cols="{ xs: 1, sm: 2, md: 3, lg: 3, xl: 4, xxl: 4 }"
          :col-gap="12"
          :row-gap="16"
          class="grid"
        >
          <a-grid-item
            v-for="(item, index) of alertRestrainList"
            :key="index"
            class="grid-item"
            @click="editRestrain(item)"
          >
            <span class="right-link-hover">
              <a-button
                type="text"
                @click.stop="
                  (event) => {
                    deleteRestrain(item, event);
                  }
                "
              >
                <template #icon>
                  <icon-delete size="15" />
                </template>
              </a-button>
            </span>
            <div class="card-wrap">
              <a-space align="start">
                <a-card-meta>
                  <template #title>
                    <a-typography-text style="margin-right: 10px">
                      系统将
                      <a-tag v-if="item.restrainType === 'fingerprint'"
                        >告警指纹</a-tag
                      >
                      <a-tag v-else-if="item.restrainType === 'label'"
                        >告警标签和标签值</a-tag
                      >
                      <a-tag v-if="item.restrainType === 'annotation'"
                        >告警注解和注解值</a-tag
                      >
                      相同的未关闭的告警事件，抑制
                      {{ item.cumulativeTime }} 分钟
                    </a-typography-text>
                  </template>

                  <template #description>
                    <div
                      v-if="item.restrainType !== 'fingerprint'"
                      class="description"
                    >
                      <div style="margin-bottom: 10px">匹配字段如下：</div>
                      <a-tag v-for="(f, i) in item.fields" :key="i" bordered>{{
                        f
                      }}</a-tag>
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
              <span>新增告警抑制</span></div
            >
          </a-grid-item>
        </a-grid>
      </div>
    </a-card>
    <RestrainDrawer
      :id="props.id"
      ref="restrainDrawerRef"
      @refresh="fetchListData"
    ></RestrainDrawer>
  </div>
</template>

<script setup lang="ts">
  import { onMounted, ref, h } from 'vue';
  import {
    queryAllAlertRestrain,
    RestrainRecord,
    deleteRestrainRecord,
  } from '@/api/restrain';
  import { Message, Notification } from '@arco-design/web-vue';
  import RestrainDrawer from './restrain-drawer.vue';

  export interface Props {
    id: number | undefined;
  }
  const props = defineProps<Props>();
  const restrainDrawerRef = ref();
  const alertRestrainList = ref<RestrainRecord[]>([]);

  const fetchListData = async () => {
    try {
      const res = await queryAllAlertRestrain(props.id as number);
      alertRestrainList.value = res.data;
    } catch (err) {
      //
    }
  };

  const handleClick = () => {
    restrainDrawerRef.value.show();
  };

  const deleteRestrain = (item: RestrainRecord, event: any) => {
    const x = `${event.clientX - 340}px`;
    const y = `${event.clientY - 100}px`;
    const id = `${Date.now()}`;
    Notification.info({
      id,
      title: '告警抑制删除提醒',
      position: 'topLeft',
      duration: 0,
      content: '请确认是否要删除该抑制策略',
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
                  await deleteRestrainRecord(item.id);
                  Message.success('告警抑制删除成功');
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

  const editRestrain = (item: RestrainRecord) => {
    restrainDrawerRef.value.show(item);
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
