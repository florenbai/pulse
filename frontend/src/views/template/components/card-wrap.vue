<template>
  <div class="card-wrap">
    <a-card v-if="props.loading" :bordered="false" hoverable>
      <slot name="skeleton"></slot>
    </a-card>
    <a-card
      v-else
      :bordered="false"
      hoverable
      @click="gotoDetail(props.template.id)"
    >
      <a-space align="start">
        <a-card-meta style="width: 100%">
          <template #title>
            <a-typography-text style="margin-right: 10px">
              {{ props.template.templateName }}
            </a-typography-text>
          </template>
          <template #description>
            <div class="description">
              {{ props.template.templateDesc }}
            </div>
            <slot></slot>
          </template>
        </a-card-meta>
      </a-space>

      <template #actions>
        <span class="icon-hover">
          <a-tooltip
            content="编辑"
            position="top"
            @click.stop="emits('open-drawer', props.template)"
            ><IconEdit /> </a-tooltip
        ></span>
        <!--
        <span class="icon-hover">
          <a-tooltip
            content="复制"
            position="top"
            @click.stop="emits('open-drawer', props.template)"
            ><IconCopy /> </a-tooltip
        ></span>
        -->
        <span class="icon-hover">
          <a-popconfirm
            content="您确认要删除模板吗？"
            type="warning"
            @ok="deleteTemplateFunc(props.template.id)"
          >
            <IconDelete @click.stop="" />
          </a-popconfirm>
        </span>
      </template>
    </a-card>
  </div>
</template>

<script lang="ts" setup>
  import { AlertTemplateRecode, deleteTemplate } from '@/api/alert-template';
  import router from '@/router';
  import { Message } from '@arco-design/web-vue';

  export interface Props {
    template: AlertTemplateRecode;
    loading: boolean;
  }
  const props = defineProps<Props>();
  const emits = defineEmits(['open-drawer', 'refresh']);

  const deleteTemplateFunc = async (id: number) => {
    try {
      await deleteTemplate(id);
      Message.success('删除成功');
      emits('refresh');
    } catch (err) {
      // you can report use errorHandler or other
    }
  };

  const gotoDetail = (id: number) => {
    router.push(`detail/${id}`);
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
