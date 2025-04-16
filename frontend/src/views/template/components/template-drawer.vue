<template>
  <a-drawer
    :width="600"
    :visible="visible"
    unmount-on-close
    @ok="handleOk"
    @cancel="handleCancel"
  >
    <template #title> 告警模板设置 </template>

    <a-form
      ref="formRef"
      layout="vertical"
      :model="formData"
      style="margin-top: 20px"
    >
      <a-form-item
        style="margin-top: 10px"
        field="templateName"
        label="告警模板名称"
        validate-trigger="input"
        :rules="[
          {
            required: true,
            message: '请输入告警模板名称',
          },
        ]"
      >
        <a-input
          v-model="formData.templateName"
          placeholder="请输入告警模板名称"
        />
      </a-form-item>
      <a-form-item
        style="margin-top: 10px"
        field="templateDesc"
        label="告警模板描述"
        validate-trigger="input"
      >
        <a-input
          v-model="formData.templateDesc"
          placeholder="请输入告警模板描述"
        />
      </a-form-item>
      <a-form-item
        style="margin-top: 10px"
        field="channels"
        label="通知渠道配置"
        validate-trigger="input"
        :rules="[
          {
            required: true,
            message: '请选择通知渠道',
          },
        ]"
      >
        <NotifyTypeSelect
          v-model="formData.channels"
          :multiple="true"
        ></NotifyTypeSelect>
      </a-form-item>
    </a-form>
  </a-drawer>
</template>

<script setup lang="ts">
  import {
    AlertTemplateRecode,
    AlertTemplateRequest,
    editTemplateForm,
    queryChannelsByTemplateId,
    submitTemplateForm,
  } from '@/api/alert-template';
  import NotifyTypeSelect from '@/components/select/notify-type-select.vue';
  import { FormInstance, Message } from '@arco-design/web-vue';
  import { ref } from 'vue';

  const emits = defineEmits(['refresh']);
  const visible = ref(false);
  const editId = ref(0);
  const formRef = ref<FormInstance>();
  const formData = ref<AlertTemplateRequest>({
    templateName: '',
    templateDesc: '',
    channels: [],
  });

  const handleOk = async () => {
    const res = await formRef.value?.validate();
    if (!res) {
      try {
        if (editId.value > 0) {
          await editTemplateForm(editId.value, formData.value);
          Message.success('模板编辑成功');
        } else {
          await submitTemplateForm(formData.value);
          Message.success('模板添加成功');
        }
        emits('refresh');
        visible.value = false;
      } catch (err) {
        //
      }
    }
  };
  const handleCancel = () => {
    visible.value = false;
    editId.value = 0;
    formData.value = {} as AlertTemplateRequest;
  };

  const getChannels = async () => {
    const { data } = await queryChannelsByTemplateId(editId.value);
    formData.value.channels = data;
  };

  const show = (item: AlertTemplateRecode | undefined) => {
    if (item !== undefined) {
      formData.value.templateName = item.templateName;
      formData.value.templateDesc = item.templateDesc;
      editId.value = item.id;
      getChannels();
    }
    visible.value = true;
  };

  defineExpose({
    show,
  });
</script>
