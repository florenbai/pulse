<template>
  <a-drawer
    :visible="visible"
    title="配置告警渠道"
    :width="600"
    @ok="handleOk"
    @cancel="handleCancel"
  >
    <a-form ref="formRef" :model="formData" layout="vertical">
      <a-form-item
        field="channelName"
        label="告警渠道名称"
        :rules="[
          {
            required: true,
            message: '请输入告警渠道名称',
          },
        ]"
      >
        <a-input
          v-model="formData.channelName"
          placeholder="请输入告警渠道名称"
        />
      </a-form-item>
      <a-form-item
        field="channelType"
        label="告警渠道类型"
        :rules="[
          {
            required: true,
            message: '请选择渠道类型',
          },
        ]"
      >
        <ChannelTypeSelect
          v-model="formData.channelType"
          placeholder="请选择渠道类型"
        />
      </a-form-item>
      <a-form-item
        field="channelSign"
        label="告警渠道标识"
        tooltip="UMDP类型请输入渠道标识，机器人类型请输入机器人Key"
        :rules="[
          {
            required: true,
            message: '请输入告警渠道标识',
          },
        ]"
      >
        <a-input
          v-model="formData.channelSign"
          placeholder="请输入告警渠道标识"
        />
      </a-form-item>
      <a-form-item
        field="channelGroup"
        label="告警渠道分组"
        tooltip="UMDP类型请输入渠道标识，机器人类型请输入机器人Key"
        :rules="[
          {
            required: true,
            message: '请输入告警渠道分组',
          },
        ]"
      >
        <a-input
          v-model="formData.channelGroup"
          placeholder="请输入告警渠道分组"
        />
      </a-form-item>
    </a-form>
  </a-drawer>
</template>

<script lang="ts" setup>
  import { ref } from 'vue';
  import { FormInstance } from '@arco-design/web-vue/es/form';
  import { Message } from '@arco-design/web-vue';

  import {
    AlertChannelRecode,
    AlertChannelRequest,
    editChannelForm,
    submitChannelForm,
  } from '@/api/alert-template';
  import ChannelTypeSelect from './channel-type-select.vue';

  const visible = ref(false);
  const itemId = ref();
  const emits = defineEmits(['refresh']);
  const formRef = ref<FormInstance>();
  const formData = ref<AlertChannelRequest>({
    channelName: '',
    channelType: '',
    channelSign: '',
    channelGroup: '',
  });
  const show = (item: AlertChannelRecode | undefined) => {
    if (item !== undefined) {
      formData.value = { ...item };
      itemId.value = item.id;
    }
    visible.value = true;
  };

  const handleCancel = () => {
    visible.value = false;
    formData.value = {} as AlertChannelRequest;
    formData.value.channelType = '';
    itemId.value = undefined;
  };

  const handleOk = async () => {
    const res = await formRef.value?.validate();
    if (!res) {
      try {
        if (itemId.value !== undefined) {
          await editChannelForm(itemId.value, formData.value);
          Message.success('告警渠道编辑成功');
        } else {
          await submitChannelForm(formData.value);
          Message.success('告警渠道添加成功');
        }

        handleCancel();
        emits('refresh');
      } catch (err) {
        //
      }
    }
  };

  defineExpose({
    show,
  });
</script>
