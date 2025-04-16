<template>
  <a-grid-item
    class="grid-item"
    style="border-style: dashed"
    @click="handleClick"
  >
    <div style="text-align: center; margin: auto">
      <icon-plus style="margin-right: 1rem" /> <span>关联告警源</span></div
    >
  </a-grid-item>
  <a-drawer
    :width="600"
    :visible="visible"
    unmount-on-close
    @ok="handleOk"
    @cancel="handleCancel"
  >
    <template #title> 关联告警源 </template>
    <div>
      <a-alert>关联的告警源只应用于当前工作区。</a-alert>
    </div>
    <div style="margin-top: 40px">
      <a-grid
        :cols="{ xs: 1, sm: 2, md: 3, lg: 3, xl: 4, xxl: 4 }"
        :col-gap="12"
        :row-gap="16"
        class="grid"
      >
        <a-grid-item
          v-for="item in alertSource"
          :key="item.id"
          class="grid-item"
          @click="checkSource(item)"
        >
          <div class="logo"> <icon-font :type="item.icon" :size="32" /></div>
          <div class="text-center text-l1 text-main">
            {{ item.sourceName }}
          </div>
          <div v-if="checkedSource === item.id">
            <div class="grid-item-checked-bg"></div>
            <div class="grid-item-checked"><icon-check /> </div
          ></div>
        </a-grid-item>
      </a-grid>
    </div>
    <a-form
      ref="formRef"
      layout="vertical"
      :model="form"
      style="margin-top: 20px"
    >
      <a-form-item
        field="sourceName"
        label="告警源名称"
        :rules="[
          {
            required: true,
            message: '请输入告警源名称',
          },
        ]"
      >
        <a-input v-model="form.sourceName" placeholder="您可以自定义告警源名称">
        </a-input>
      </a-form-item>
      <a-form-item field="description" label="告警源描述">
        <a-input
          v-model="form.description"
          placeholder="您可以自定义告警源描述"
        >
        </a-input>
      </a-form-item>
    </a-form>
  </a-drawer>
</template>

<script lang="ts" setup>
  import { ref, onMounted } from 'vue';
  import { FormInstance, Message } from '@arco-design/web-vue';
  import { AlertSourceRecode, queryAllAlertSource } from '@/api/alert-source';
  import { relatedWorkspaceAlertSourceForm } from '@/api/workspace';

  export interface Props {
    id: number | undefined;
  }
  const props = defineProps<Props>();
  const emits = defineEmits(['updateData']);

  const visible = ref(false);
  const checkedSource = ref(0);
  const formRef = ref<FormInstance>();
  const form = ref({
    sourceId: 0,
    sourceName: '',
    description: '',
  });
  const handleClick = () => {
    visible.value = true;
  };
  const handleOk = async () => {
    const res = await formRef.value?.validate();
    if (!res) {
      try {
        await relatedWorkspaceAlertSourceForm(props.id as number, form.value);

        Message.success({
          content: '操作成功',
          duration: 5 * 1000,
        });
        emits('updateData');
      } catch (err) {
        return;
      }
    } else {
      return;
    }
    visible.value = false;
  };
  const handleCancel = () => {
    visible.value = false;
  };

  const alertSource = ref<AlertSourceRecode[]>([]);
  onMounted(async () => {
    try {
      const res = await queryAllAlertSource();
      alertSource.value = res.data;
    } catch (err) {
      //
    }
  });

  const checkSource = (record: AlertSourceRecode) => {
    checkedSource.value = record.id;
    form.value.sourceId = record.id;
    form.value.sourceName = record.sourceName;
  };
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
</style>
