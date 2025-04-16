<template>
  <div class="container">
    <Breadcrumb
      :items="[
        { label: 'menu.alert' },
        { label: 'menu.source.integration', path: '/owl/alert/integration' },
      ]"
    />
    <a-card class="general-card" title="告警集成管理">
      <a-grid
        :cols="{ xs: 1, sm: 2, md: 3, lg: 3, xl: 4, xxl: 4 }"
        :col-gap="12"
        :row-gap="16"
        class="grid"
      >
        <a-grid-item
          v-for="item of alertSource"
          :key="item.id"
          class="grid-item"
          @click.stop="gotoDetail(item.id)"
        >
          <div style="display: inline-flex">
            <div class="logo" style="margin-right: 8px">
              <icon-font :type="item.icon" :size="36"
            /></div>
            <div class="text-center text-l1 text-main">
              {{ item.name }}
              <div class="description">{{ item.description }}</div>
              <div v-if="item.eventCount === 0" class="text-tip">
                <span style="color: darkorange; display: inline-flex">
                  尚未收到该告警源的告警
                </span>
              </div>
              <div v-else class="text-tip">
                <span style="color: #2b912b; display: inline-flex">
                  累计收到 {{ item.eventCount }} 条告警
                </span>
              </div>
            </div>
            <div
              class="delete-icon"
              @click.stop="(event) => deleteSource(item.id, event)"
            >
              <icon-delete :size="20" />
            </div>
          </div>
        </a-grid-item>

        <alertSourceTemplate
          @update-data="getSystemSource"
        ></alertSourceTemplate>
      </a-grid>
    </a-card>
  </div>
</template>

<script lang="ts" setup>
  import { Message, Notification } from '@arco-design/web-vue';
  import { h, onMounted, ref } from 'vue';

  import {
    deleteSystemIntegration,
    querySystemIntegration,
    SystemIntegrationRecode,
  } from '@/api/integration';
  import router from '@/router';
  import alertSourceTemplate from './components/alert-source-relation.vue';

  const alertSource = ref<SystemIntegrationRecode[]>([]);

  const getSystemSource = async () => {
    try {
      const res = await querySystemIntegration();
      alertSource.value = res.data;
    } catch (err) {
      //
    }
  };
  onMounted(() => {
    getSystemSource();
  });

  const deleteSource = async (id: number, event: any) => {
    const x = `${event.clientX - 340}px`;
    const y = `${event.clientY - 100}px`;
    const nid = `${Date.now()}`;
    Notification.info({
      id: nid,
      title: '告警集成删除提醒',
      position: 'topLeft',
      duration: 0,
      content: '请确认是否要删除该告警集成',
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
                Notification.remove(nid);
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
                  await deleteSystemIntegration(id);
                  Message.success('告警集成删除成功');
                  getSystemSource();
                  Notification.remove(nid);
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

  const gotoDetail = (id: number) => {
    router.push(`/alert/integration/detail/${id}`);
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
  .grid {
    gap: 0.75rem;
  }
  .grid-item {
    display: flex;
    height: 100%;
    cursor: pointer;
    position: relative;
    flex-direction: column;
    overflow: hidden;
    border-radius: 0.25rem;
    border-width: 1px;
    border-style: solid;
    padding-top: 1rem !important;
    padding-left: 1.5rem;
    padding-bottom: 1rem !important;
    transition-property: color, background-color, border-color,
      text-decoration-color, fill, stroke, opacity, box-shadow, transform,
      filter, backdrop-filter, -webkit-text-decoration-color,
      -webkit-backdrop-filter !important;
    transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1) !important;
    transition-duration: 0.15s !important;
    border-color: #e5e7eb;
    background-color: rgb(var(--fc-fill-3) / var(--tw-bg-opacity)) !important;
  }

  .grid-item:hover {
    background-color: var(--color-secondary) !important;
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
</style>
