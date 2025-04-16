<template>
  <a-drawer
    :width="600"
    :visible="visible"
    unmount-on-close
    @ok="handleOk"
    @cancel="handleCancel"
  >
    <template #title> 配置告警静默规则 </template>
    <div>
      <a-alert
        >您可以配置多个告警静默规则，但系统只会按顺序匹配一个执行。当告警被静默时，仍然会在告警列表中展示和处理，只是不会触发通知。</a-alert
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
        field="silenceName"
        label="静默规则名称"
        validate-trigger="input"
        :rules="[
          {
            required: true,
            message: '请输入静默规则名称',
          },
        ]"
      >
        <a-input
          v-model="formData.silenceName"
          placeholder="请输入静默规则名称"
        />
      </a-form-item>
      <a-form-item
        style="margin-top: 10px"
        field="aggregationDesc"
        label="静默规则描述"
        validate-trigger="input"
      >
        <a-input
          v-model="formData.silenceDesc"
          placeholder="请输入静默规则描述"
        />
      </a-form-item>
      <a-form-item
        style="margin-top: 10px"
        field="aggregationDesc"
        label="静默时间"
      >
        <div>
          <a-radio-group
            v-model="formData.silenceType"
            type="button"
            default-value="once"
            @change="changeType"
          >
            <a-radio value="once">单次静默</a-radio>
            <a-radio value="period">周期静默</a-radio>
          </a-radio-group>
          <div v-if="formData.silenceType === 'once'" style="margin-top: 10px">
            <a-range-picker
              v-model="formData.silenceTime.rangeTime"
              show-time
              :style="{
                width: '480px',
                marginRight: '24px',
              }"
            />
          </div>
          <div
            v-if="formData.silenceType === 'period'"
            style="margin-top: 10px"
          >
            <a-select
              v-model="formData.silenceTime.weeks"
              :style="{ width: '200px' }"
              :max-tag-count="1"
              placeholder="请选择通知时间周期"
              multiple
              allow-clear
            >
              <a-option :value="1">星期一</a-option>
              <a-option :value="2">星期二</a-option>
              <a-option :value="3">星期三</a-option>
              <a-option :value="4">星期四</a-option>
              <a-option :value="5">星期五</a-option>
              <a-option :value="6">星期六</a-option>
              <a-option :value="0">星期日</a-option>
            </a-select>
            <a-time-picker
              v-model="formData.silenceTime.times"
              :style="{ width: '360px' }"
              type="time-range"
            />
          </div>
        </div>
      </a-form-item>
      <a-form-item
        field="aggregationDesc"
        label="静默条件"
        validate-trigger="input"
        tooltip="您可以使用标签以及‘告警标题’来设置静默条件"
      >
        <div>
          <a-card>
            <div style="display: flex">
              <div
                v-if="
                  formData.filters !== undefined && formData.filters.length > 1
                "
                class="rule-condition-logic-group__logic"
              >
                <div
                  class="rule-condition-logic-group__logic-item"
                  :style="{
                    height: `${tagHeight}px`,
                  }"
                >
                  <div
                    class="rule-condition-logic-group__logic-item__start"
                    style="width: 12px"
                  ></div>
                  <span class="rule-condition-logic-group__logic-item__text"
                    >且</span
                  >
                  <div
                    class="rule-condition-logic-group__logic-item__end"
                    style="width: 12px"
                  ></div>
                </div>
              </div>
              <div>
                <a-input-group
                  v-for="(item, index) of formData.filters"
                  :key="index"
                  style="margin-bottom: 8px"
                >
                  <a-input
                    v-model="item.tag"
                    :style="{ width: '140px' }"
                    placeholder="标签或告警标题"
                  />
                  <a-select
                    v-model="item.operation"
                    style="width: 100px; margin-left: 5px; margin-right: 5px"
                  >
                    <a-option value="IN">包含</a-option>
                    <a-option value="NOT IN">不包含</a-option>
                  </a-select>
                  <a-input-tag
                    v-model="item.value"
                    allow-clear
                    :style="{ width: '230px' }"
                    placeholder="值"
                  />
                  <a-button
                    v-if="formData.filters.length > 1"
                    style="margin-left: 5px; background-color: #fff"
                    @click="deleteFilterItem(index)"
                  >
                    <template #icon> <icon-delete size="15" /> </template
                  ></a-button>
                </a-input-group>
              </div>
            </div>

            <a-button type="text" @click="createCondition">+ 添加条件</a-button>
          </a-card>
        </div>
      </a-form-item>
    </a-form>
  </a-drawer>
</template>

<script setup lang="ts">
  import {
    SilenceRequest,
    submitSilenceForm,
    editSilenceForm,
    SilenceRecord,
  } from '@/api/alert-denoise';
  import { FormInstance, Message } from '@arco-design/web-vue';
  import { ref } from 'vue';

  export interface Props {
    id: number | undefined;
  }
  const emits = defineEmits(['refresh']);
  const props = defineProps<Props>();
  const visible = ref(false);
  const tagHeight = ref(40);
  const formRef = ref<FormInstance>();
  const itemId = ref();
  const formData = ref<SilenceRequest>({
    workspaceId: props.id as number,
    silenceName: '',
    silenceDesc: '',
    silenceType: 'once',
    silenceTime: {
      rangeTime: [],
    },
    filters: [
      {
        tag: '',
        operation: 'IN',
        value: [],
      },
    ],
  });
  const handleCancel = () => {
    visible.value = false;
    formData.value = {
      workspaceId: props.id as number,
      silenceName: '',
      silenceDesc: '',
      silenceType: 'once',
      silenceTime: {
        rangeTime: [],
      },
      filters: [
        {
          tag: '',
          operation: 'IN',
          value: [],
        },
      ],
    };
  };
  const handleOk = async () => {
    const res = await formRef.value?.validate();
    if (!res) {
      try {
        if (itemId.value !== undefined) {
          await editSilenceForm(itemId.value, formData.value);
          Message.success('静默编辑成功');
        } else {
          await submitSilenceForm(formData.value);
          Message.success('静默添加成功');
        }

        handleCancel();
        emits('refresh');
      } catch (err) {
        //
      }
    }
  };

  const changeType = (value: string | number | boolean) => {
    if (value === 'once') {
      formData.value.silenceTime.weeks = [];
      formData.value.silenceTime.times = [];
    } else {
      formData.value.silenceTime.rangeTime = [];
    }
  };

  const show = (item: SilenceRecord | undefined) => {
    if (item !== undefined) {
      formData.value = { ...item };
      itemId.value = item.id;
      if (formData.value.filters.length > 2) {
        tagHeight.value = (formData.value.filters.length - 2) * 40 + 40;
      }
    }
    visible.value = true;
  };

  const createCondition = () => {
    formData.value.filters.push({
      tag: '',
      operation: 'IN',
      value: [],
    });
    if (formData.value.filters.length > 2) {
      tagHeight.value = (formData.value.filters.length - 2) * 40 + 40;
    }
  };

  const deleteFilterItem = (index: number) => {
    if (index === 0) {
      formData.value.filters.shift();
    } else {
      formData.value.filters.splice(index, 1);
    }
    if (formData.value.filters.length >= 2) {
      tagHeight.value -= 40;
    }
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
