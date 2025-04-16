<template>
  <div class="card-wrap">
    <a-card v-if="loading" :bordered="false" hoverable>
      <slot name="skeleton"></slot>
    </a-card>
    <a-card v-else :bordered="false" hoverable @click="gotoDetail(props.id)">
      <a-space align="start">
        <a-card-meta style="width: 100%">
          <template #title>
            <a-typography-text style="margin-right: 10px">
              {{ title }}
            </a-typography-text>
            <a-tag v-if="open" size="small" color="green">
              <template #icon>
                <icon-check-circle-fill />
              </template>
              <span>已就绪</span>
            </a-tag>
            <a-tag v-else size="small" color="red">
              <template #icon>
                <icon-exclamation-circle-fill />
              </template>
              <span>已禁用</span>
            </a-tag>

            <a-popconfirm
              :content="switchContent"
              type="warning"
              @ok="switchChange(props.id)"
            >
              <a-switch
                v-model="open"
                style="position: absolute; right: 20px"
                :before-change="handleChangeIntercept"
                @click.stop=""
              />
            </a-popconfirm>
          </template>
          <template #description>
            <div class="description">
              {{ description }}
            </div>
            <div class="info">
              <a-statistic
                title="待处理告警"
                :value="pending"
                :value-from="0"
                :start="true"
                animation
              >
              </a-statistic>
              <a-statistic
                title="处理中告警"
                :value="processing"
                :value-from="0"
                :start="true"
                animation
                style="margin-left: 50px"
              >
              </a-statistic>
            </div>
          </template>
        </a-card-meta>
      </a-space>

      <template #actions>
        <span class="icon-hover">
          <a-tooltip content="编辑" position="top"
            ><IconEdit @click.stop="editWorkspace(props.id)" /> </a-tooltip
        ></span>
        <span class="icon-hover">
          <a-tooltip content="删除" position="top">
            <a-popconfirm
              content="删除工作区后，您将不再收到配置在该工作区源的所有告警信息，您确认删除吗？"
              type="warning"
              @ok="deleteWorkspace(props.id)"
            >
              <IconDelete @click.stop="" />
            </a-popconfirm>
          </a-tooltip>
        </span>
      </template>
    </a-card>
    <editWorkspaceComponent
      ref="workspaceRef"
      @refresh="refresh"
    ></editWorkspaceComponent>
  </div>
</template>

<script lang="ts" setup>
  import router from '@/router';
  import { useToggle } from '@vueuse/core';
  import { ref, watch } from 'vue';
  import { deleteWorkspaceItem, changeWorkspaceStatus } from '@/api/workspace';
  import { Message } from '@arco-design/web-vue';
  import editWorkspaceComponent from './index/edit.vue';

  const workspaceRef = ref();
  const emits = defineEmits(['refresh']);
  const props = defineProps({
    id: {
      type: Number,
      default: 0,
    },
    loading: {
      type: Boolean,
      default: false,
    },
    title: {
      type: String,
      default: '',
    },
    description: {
      type: String,
      default: '',
    },
    defaultValue: {
      type: Boolean,
      default: false,
    },
    sourceCount: {
      type: Number,
      default: 0,
    },
    pending: {
      type: Number,
      default: 0,
    },
    processing: {
      type: Number,
      default: 0,
    },
  });
  const [open, toggle] = useToggle(props.defaultValue);
  const switchContent = ref('工作区被禁用后，将无法接收新告警。');
  watch(open, () => {
    if (open) {
      switchContent.value = '工作区被禁用后，将无法接收新告警。';
    } else {
      switchContent.value = '工作区被启用后，将会接收新的告警。';
    }
  });
  const handleChangeIntercept = () => {
    return false;
  };
  const gotoDetail = (id: number) => {
    router.push(`/alert/workspace/detail/${id}`);
  };
  const switchChange = async (id: number) => {
    try {
      await changeWorkspaceStatus(id);
      if (open) {
        toggle();
        switchContent.value = '工作区被启用后，将会接收新的告警。';
      } else {
        toggle();
        switchContent.value = '工作区被禁用后，将无法接收新告警。';
      }
    } catch {
      //
    }
  };

  const editWorkspace = (id: number) => {
    workspaceRef.value.show(id);
  };

  const refresh = () => {
    emits('refresh');
  };

  const deleteWorkspace = async (id: number) => {
    try {
      await deleteWorkspaceItem(id);
      Message.success('删除成功');
      emits('refresh');
    } catch {
      //
    }
  };
</script>

<style scoped lang="less">
  .card-wrap {
    height: 100%;
    transition: all 0.3s;
    border: 1px solid var(--color-neutral-3);
    border-radius: 4px;
    &:hover {
      transform: translateY(-4px);
      // box-shadow: 4px 4px 10px rgba(0, 0, 0, 0.1);
    }
    :deep(.arco-card) {
      height: 100%;
      border-radius: 4px;
      .arco-card-body {
        height: 100%;
        .arco-space {
          width: 100%;
          height: 100%;
          .arco-space-item {
            height: 100%;
            &:last-child {
              flex: 1;
            }
            .arco-card-meta {
              height: 100%;
              display: flex;
              flex-flow: column;
              .arco-card-meta-content {
                flex: 1;
                .arco-card-meta-description {
                  margin-top: 8px;
                  color: rgb(var(--gray-6));
                  line-height: 20px;
                  font-size: 12px;
                }
              }
              .arco-card-meta-footer {
                margin-top: 0;
              }
            }
          }
        }
      }
    }
    :deep(.arco-card-meta-title) {
      display: flex;
      align-items: center;

      // To prevent the shaking
      line-height: 28px;
    }
    :deep(.arco-skeleton-line) {
      &:last-child {
        display: flex;
        justify-content: flex-end;
        margin-top: 20px;
      }
    }
  }

  .description {
    min-height: 60px;
    -webkit-line-clamp: 3;
    display: -webkit-inline-box;
    -webkit-box-orient: vertical;
    overflow: hidden;
    text-overflow: ellipsis;
    width: 85%;
  }
</style>
