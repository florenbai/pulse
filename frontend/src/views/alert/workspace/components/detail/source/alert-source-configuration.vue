<template>
  <div>
    <a-grid
      :cols="{ xs: 1, sm: 2, md: 3, lg: 3, xl: 4, xxl: 4 }"
      :col-gap="12"
      :row-gap="16"
      class="grid"
    >
      <alertSourceList
        v-for="item of alertSource"
        :key="item.id"
        :record="item"
        @refresh="getWorkspaceAlertSource"
        @handle-click-warning="handleClickWarning"
      />
      <alertSourceTemplate
        :id="props.id"
        @update-data="getWorkspaceAlertSource"
      ></alertSourceTemplate>
    </a-grid>
  </div>
</template>

<script lang="ts" setup>
  import {
    WorkspaceAlertSourceRecode,
    queryWorkspaceAlertSource,
    deleteAlertSource,
  } from '@/api/alert-source';
  import { Modal } from '@arco-design/web-vue';
  import { onMounted, ref } from 'vue';
  import { AlertSourceRequest } from '@/api/workspace';

  import alertSourceTemplate from './alert-source-relation.vue';
  import alertSourceList from './alert-source-list.vue';

  export interface Props {
    id: number | undefined;
  }
  const props = defineProps<Props>();
  const alertSource = ref<WorkspaceAlertSourceRecode[]>([]);
  const alertSourceModel = ref<AlertSourceRequest>({ alertSource: [] });
  onMounted(async () => {
    try {
      const res = await queryWorkspaceAlertSource(props.id);
      alertSource.value = res.data;
    } catch (err) {
      //
    }
  });

  const submitAlertSource = async (): Promise<boolean> => {
    if (alertSourceModel.value.alertSource.length < 1) {
      return false;
    }
    return true;
  };

  const handleClickWarning = (id: number) => {
    Modal.warning({
      hideCancel: false,
      title: '告警源删除提醒',
      content:
        '删除告警源后，在当前工作区内，您将不再收到通过该源发出的告警！您是否确认删除？',
      onOk: async () => {
        try {
          await deleteAlertSource(id);
          const res = await queryWorkspaceAlertSource(props.id);
          alertSource.value = res.data;
        } catch (err) {
          //
        }
      },
    });
  };

  const getWorkspaceAlertSource = async () => {
    try {
      const res = await queryWorkspaceAlertSource(props.id);
      alertSource.value = res.data;
    } catch (err) {
      //
    }
  };

  defineExpose({
    submitAlertSource,
    handleClickWarning,
  });
</script>

<style scoped>
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
