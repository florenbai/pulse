<template>
  <div :style="{ display: 'flex' }">
    <a-card title="告警静默" :bordered="false" :style="{ width: '100%' }">
      <div>
        <a-grid
          :cols="{ xs: 1, sm: 2, md: 3, lg: 3, xl: 4, xxl: 4 }"
          :col-gap="12"
          :row-gap="16"
          class="grid"
        >
          <a-grid-item
            v-for="(item, index) of alertSilenceList"
            :key="index"
            class="grid-item"
          >
            <div class="card-wrap" @click="editSilence(item)">
              <a-space align="start">
                <a-card-meta>
                  <template #title>
                    <a-typography-text style="margin-right: 10px">
                      {{ item.silenceName }}
                    </a-typography-text>
                    <a-tag
                      v-if="checkSilence(item.silenceType, item.silenceTime)"
                      size="small"
                      color="green"
                    >
                      <span>当前时间静默已生效</span>
                    </a-tag>
                    <a-tag v-else size="small" color="orange">
                      <span>当前时间静默未生效</span>
                    </a-tag>
                    <span class="right-link-hover">
                      <a-button
                        type="text"
                        @click.stop="(event) => deleteSilence(item, event)"
                      >
                        <template #icon>
                          <icon-delete size="15" />
                        </template>
                      </a-button>
                    </span>
                  </template>

                  <template #description>
                    <div class="description">
                      {{ item.silenceDesc }}
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
              <span>新增告警静默</span></div
            >
          </a-grid-item>
        </a-grid>
      </div>
    </a-card>
    <silenceDrawer
      :id="props.id"
      ref="silenceDrawerRef"
      @refresh="fetchListData"
    ></silenceDrawer>
  </div>
</template>

<script setup lang="ts">
  import { onMounted, ref, h } from 'vue';
  import {
    queryAllAlertSilence,
    SilenceRecord,
    SilenceTime,
    deleteSilenceRecord,
  } from '@/api/alert-denoise';
  import { Message, Notification } from '@arco-design/web-vue';
  import silenceDrawer from './silence-drawer.vue';

  export interface Props {
    id: number | undefined;
  }
  const props = defineProps<Props>();
  const silenceDrawerRef = ref();
  const alertSilenceList = ref<SilenceRecord[]>([]);

  const fetchListData = async () => {
    try {
      const res = await queryAllAlertSilence(props.id as number);
      alertSilenceList.value = res.data;
    } catch (err) {
      //
    }
  };

  const handleClick = () => {
    silenceDrawerRef.value.show();
  };

  const deleteSilence = (item: SilenceRecord, event: any) => {
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
                  await deleteSilenceRecord(item.id);
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

  function formatDate(date: Date) {
    const year = date.getFullYear();

    // 注意月份是从0开始的，所以需要+1
    const month = (1 + date.getMonth()).toString().padStart(2, '0');

    // 日期也需要填充前导零
    const day = date.getDate().toString().padStart(2, '0');

    return `${year}-${month}-${day}`;
  }

  const checkSilence = (type: string, times: SilenceTime) => {
    const now = new Date();
    if (type === 'once') {
      if (times.rangeTime !== undefined) {
        const startTime = new Date(times.rangeTime[0]);
        const endTime = new Date(times.rangeTime[1]);

        return now >= startTime && now <= endTime;
      }
    } else if (times.weeks?.includes(now.getDay())) {
      if (times.times !== undefined) {
        const ymd = formatDate(now);
        const startTime = new Date(`${ymd} ${times.times[0]}`);
        const endTime = new Date(`${ymd} ${times.times[1]}`);
        return now >= startTime && now <= endTime;
      }
    }
    return false;
  };

  const editSilence = (item: SilenceRecord) => {
    silenceDrawerRef.value.show(item);
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
