<template>
  <a-drawer
    :width="600"
    :visible="visible"
    unmount-on-close
    @ok="handleOk"
    @cancel="handleCancel"
  >
    <template #title> 配置告警抑制规则 </template>
    <div>
      <a-alert
        >根据您选择的抑制类型，根据标签、注解字段相同值或告警指纹相同的告警，在告警触发后的一段时间内进行抑制(只会抑制未关闭的，不会抑制已关闭的)。</a-alert
      >
    </div>
    <a-form
      ref="formRef"
      layout="vertical"
      :model="formData"
      style="margin-top: 20px"
    >
      <a-form-item
        style="margin-top: 10px"
        field="restrainType"
        label="告警抑制类型"
        :rules="[
          {
            required: true,
            message: '请选择告警抑制类型',
          },
        ]"
      >
        <div>
          <a-radio-group
            v-model="formData.restrainType"
            type="button"
            default-value="fingerprint"
          >
            <a-radio value="fingerprint">告警指纹</a-radio>
            <a-radio value="label">告警标签</a-radio>
            <a-radio value="annotation">告警注解</a-radio>
          </a-radio-group>
        </div>
      </a-form-item>
      <a-form-item
        v-if="formData.restrainType !== 'fingerprint'"
        style="margin-top: 10px"
        field="fields"
        label="匹配字段"
        validate-trigger="input"
        :rules="[
          {
            required: true,
            message: '请输入匹配字段',
          },
        ]"
      >
        <a-input-tag
          v-model="formData.fields"
          placeholder="请输入匹配字段"
          allow-clear
        >
        </a-input-tag>
      </a-form-item>
      <a-form-item
        style="margin-top: 10px"
        field="cumulativeTime"
        label="告警抑制时间"
        tooltip="该窗口为滑动窗口，会随新告警触发而延长。"
        validate-trigger="input"
        :rules="[
          {
            required: true,
            message: '请输入告警抑制时间',
          },
        ]"
      >
        <a-input-number
          v-model="formData.cumulativeTime"
          placeholder="请输入告警抑制时间"
          allow-clear
          hide-button
        >
          <template #suffix> 分钟 </template>
        </a-input-number>
      </a-form-item>
    </a-form>
  </a-drawer>
</template>

<script setup lang="ts">
  import {
    RestrainRequest,
    RestrainRecord,
    submitRestrainForm,
    editRestrainForm,
  } from '@/api/restrain';
  import { FormInstance, Message } from '@arco-design/web-vue';
  import { ref } from 'vue';

  export interface Props {
    id: number | undefined;
  }
  const emits = defineEmits(['refresh']);
  const props = defineProps<Props>();
  const visible = ref(false);
  const formRef = ref<FormInstance>();
  const itemId = ref();
  const formData = ref<RestrainRequest>({
    restrainType: 'fingerprint',
    cumulativeTime: 120,
  });
  const handleCancel = () => {
    visible.value = false;
    formData.value = {
      restrainType: 'fingerprint',
      cumulativeTime: 120,
    };
  };
  const handleOk = async () => {
    const res = await formRef.value?.validate();
    if (!res) {
      try {
        if (itemId.value !== undefined) {
          await editRestrainForm(itemId.value, formData.value);
          Message.success('告警抑制编辑成功');
        } else {
          await submitRestrainForm(props.id as number, formData.value);
          Message.success('告警抑制添加成功');
        }

        handleCancel();
        emits('refresh');
      } catch (err) {
        //
      }
    }
  };

  const show = (item: RestrainRecord | undefined) => {
    if (item !== undefined) {
      formData.value = { ...item };
      itemId.value = item.id;
    }
    visible.value = true;
  };

  defineExpose({
    show,
  });
</script>

<style scoped lang="less">
  .rule-condition-logic-group__logic {
    width: 27px;
    flex-grow: 0;
    flex-shrink: 0;
    overflow: hidden;
    position: relative;
    top: 16px;
  }

  .rule-condition-logic-group__logic-item {
    align-items: flex-start;
    box-sizing: initial;
    display: flex;
    flex-direction: column;
    overflow: hidden;
    pointer-events: none;
    position: absolute;
  }

  .rule-condition-logic-group__logic-item__start {
    border-top: 1px solid #d0d3d6;
    border-top-left-radius: 3px;
    border-left: 1px solid #d0d3d6;
    flex-grow: 1;
    margin-left: 7px;
    width: 12px;
  }

  .rule-condition-logic-group__logic-item__text {
    color: #646a73;
    line-height: 22px;
  }

  .rule-condition-logic-group__logic-item__end {
    width: 12px;
    border-bottom: 1px solid #d0d3d6;
    border-bottom-left-radius: 3px;
    border-left: 1px solid #d0d3d6;
    flex-grow: 1;
    margin-left: 7px;
  }
</style>
