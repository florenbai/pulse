<template>
  <a-drawer
    :visible="visible"
    title="配置告警等级"
    :width="600"
    @ok="handleOk"
    @cancel="handleCancel"
  >
    <a-form ref="formRef" :model="formData" layout="vertical">
      <a-form-item
        field="levelName"
        label="告警等级名称"
        :rules="[
          {
            required: true,
            message: '请输入告警等级名称',
          },
        ]"
      >
        <a-input
          v-model="formData.levelName"
          placeholder="请输入告警等级名称"
        />
      </a-form-item>
      <a-form-item field="levelDesc" label="告警等级描述">
        <a-input
          v-model="formData.levelDesc"
          placeholder="请输入告警等级描述"
        />
      </a-form-item>
      <a-form-item
        field="color"
        label="告警等级颜色"
        :rules="[
          {
            required: true,
            message: '请选择告警等级颜色',
          },
        ]"
      >
        <input v-model="formData.color" type="color" />
      </a-form-item>
    </a-form>
  </a-drawer>
</template>

<script lang="ts" setup>
  import { ref } from 'vue';
  import { FormInstance } from '@arco-design/web-vue/es/form';
  import { Message } from '@arco-design/web-vue';
  import {
    AlertLevelRecode,
    AlertLevelRequest,
    editLevelForm,
    submitLevelForm,
  } from '@/api/alert-level';

  const visible = ref(false);
  const itemId = ref();
  const emits = defineEmits(['refresh']);
  const formRef = ref<FormInstance>();
  const formData = ref<AlertLevelRequest>({
    levelName: '',
    levelDesc: '',
    color: '#000000',
  });
  const show = (item: AlertLevelRecode | undefined) => {
    if (item !== undefined) {
      formData.value = { ...item };
      itemId.value = item.id;
    }
    visible.value = true;
  };

  const handleCancel = () => {
    visible.value = false;
    formData.value = {
      levelName: '',
      levelDesc: '',
      color: '#000000',
    };
    itemId.value = undefined;
  };

  const handleOk = async () => {
    const res = await formRef.value?.validate();
    if (!res) {
      try {
        if (itemId.value !== undefined) {
          await editLevelForm(itemId.value, formData.value);
          Message.success('告警等级编辑成功');
        } else {
          await submitLevelForm(formData.value);
          Message.success('告警等级添加成功');
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
