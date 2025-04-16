<template>
  <a-drawer
    :width="600"
    :visible="visible"
    unmount-on-close
    @ok="handleOk"
    @cancel="handleCancel"
  >
    <template #title> 配置告警聚合策略 </template>
    <div>
      <a-alert
        >您可以配置多个警聚合策略，但系统只会按顺序匹配一个策略执行。您可以通过配置聚合维度来将具有相同属性以及标签的告警聚合。被聚合的告警只会按照配置（聚合窗口和预警值）进行通知。</a-alert
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
        field="aggregationName"
        label="聚合策略名称"
        validate-trigger="input"
        :rules="[
          {
            required: true,
            message: '请输入聚合策略名称',
          },
        ]"
      >
        <a-input
          v-model="formData.aggregationName"
          placeholder="请输入聚合策略名称"
        />
      </a-form-item>
      <a-form-item
        style="margin-top: 10px"
        field="aggregationDesc"
        label="聚合策略描述"
        validate-trigger="input"
      >
        <a-input
          v-model="formData.aggregationDesc"
          placeholder="请输入聚合策略描述"
        />
      </a-form-item>

      <a-alert
        >当告警的属性或标签与某个未关闭故障完全相同时，将告警进行聚合以达到减少通知的目的。</a-alert
      >
      <a-form-item field="dimension" label="属性维度" style="margin-top: 10px">
        <a-checkbox v-model="formData.titleDimension">告警标题</a-checkbox>
        <a-checkbox v-model="formData.levelDimension" style="margin-left: 10px"
          >告警等级</a-checkbox
        >
      </a-form-item>
      <a-form-item field="tagsDimension">
        <template #label>
          <div>
            <span>标签维度</span>
            <a-switch
              size="small"
              style="margin-left: 20px"
              :default-checked="showTagFilter"
              @change="changeTagFilter"
            />
          </div>
        </template>
        <div v-if="showTagFilter" style="width: -webkit-fill-available">
          <div
            v-for="(filter, filterIndex) of formData.tagsDimension"
            :key="filterIndex"
          >
            <a-card :style="{ 'margin-top': filterIndex > 0 ? '10px' : '0' }">
              <div style="display: flex">
                <div
                  v-if="filter.values !== null && filter.values.length > 1"
                  class="rule-condition-logic-group__logic"
                >
                  <div
                    class="rule-condition-logic-group__logic-item"
                    :style="{
                      height: `${tagHeight[filterIndex]}px`,
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
                    v-for="(item, index) of filter.values"
                    :key="index"
                    style="margin-bottom: 8px"
                  >
                    <a-input v-model="item.tag" placeholder="标签" />
                    <a-select
                      v-model="item.operation"
                      style="margin-left: 5px; margin-right: 5px"
                    >
                      <a-option value="IN">包含</a-option>
                      <a-option value="NOT IN">不包含</a-option>
                    </a-select>
                    <a-input-tag
                      v-model="item.value"
                      allow-clear
                      style="min-width: 250px"
                      placeholder="值"
                    />
                    <a-button
                      style="margin-left: 5px; background-color: #fff"
                      @click="deleteFilterItem(filterIndex, index)"
                    >
                      <template #icon> <icon-delete size="15" /> </template
                    ></a-button>
                  </a-input-group>
                </div>
              </div>

              <a-button type="text" @click="createCondition(filterIndex)"
                >+ 添加条件</a-button
              >
            </a-card>
          </div>

          <a-button
            style="width: -webkit-fill-available; margin-top: 8px"
            type="dashed"
            @click="createFilter"
            >+ 或者</a-button
          >
        </div>
      </a-form-item>
      <a-form-item
        field="windows"
        label="聚合窗口"
        :rules="[
          {
            required: true,
            message: '请输入聚合窗口',
          },
        ]"
      >
        <a-input-number
          v-model="formData.windows"
          :precision="0"
          :step="1"
          :min="0"
          :max="1440"
        >
          <template #suffix> 分钟 </template>
        </a-input-number>
        <template #extra>
          <div
            >聚合设定时间内的告警（0表示无聚合窗口），超出时间范围后告警将被聚合到新的故障中。</div
          >
        </template>
      </a-form-item>
      <a-form-item
        field="storm"
        label="风暴预警"
        :rules="[
          {
            required: true,
            message: '请输入风暴预警',
          },
        ]"
      >
        <a-input-number
          v-model="formData.storm"
          :precision="0"
          :step="1"
          :min="0"
          :max="1440"
        >
          <template #suffix> 条 </template>
        </a-input-number>
        <template #extra>
          <div
            >聚合告警达到设定条数后（0表示不进行风暴预警），将再次发送消息通知</div
          >
        </template>
      </a-form-item>
    </a-form>
  </a-drawer>
</template>

<script setup lang="ts">
  import {
    AlertAggregationRequest,
    submitAggregationForm,
    editAggregationForm,
    AlertAggregationRecord,
  } from '@/api/alert-denoise';
  import { FormInstance, Message } from '@arco-design/web-vue';
  import { ref, watch } from 'vue';

  export interface Props {
    id: number | undefined;
  }
  const emits = defineEmits(['refresh']);
  const activeKeyRef = ref<(string | number)[]>(['1']);
  const props = defineProps<Props>();
  const visible = ref(false);
  const showTagFilter = ref<boolean>(false);
  const dimension = ref(false);
  const formRef = ref<FormInstance>();
  const itemId = ref();
  const tagHeight = [40];
  const formData = ref<AlertAggregationRequest>({
    workspaceId: props.id as number,
    aggregationName: '',
    aggregationDesc: '',
    titleDimension: false,
    levelDimension: false,
    tagsDimension: [],
    windows: 0,
    storm: 0,
  });
  const handleOk = async () => {
    const res = await formRef.value?.validate();
    if (!res) {
      try {
        if (itemId.value !== undefined) {
          await editAggregationForm(itemId.value, formData.value);
          Message.success('策略编辑成功');
        } else {
          await submitAggregationForm(formData.value);
          Message.success('策略添加成功');
        }

        visible.value = false;
        formData.value = {} as AlertAggregationRequest;
        emits('refresh');
      } catch (err) {
        //
      }
    }
  };
  const handleCancel = () => {
    visible.value = false;
    formData.value = {} as AlertAggregationRequest;
  };

  const show = (item: AlertAggregationRecord | undefined) => {
    if (item !== undefined) {
      formData.value = { ...item };
      itemId.value = item.id;
    }
    visible.value = true;
  };

  const createCondition = (filterIndex: number) => {
    formData.value.tagsDimension[filterIndex].values.push({
      tag: '',
      operation: 'IN',
      value: [],
    });
    if (formData.value.tagsDimension[filterIndex].values.length > 2) {
      tagHeight[filterIndex] =
        (formData.value.tagsDimension[filterIndex].values.length - 2) * 40 + 40;
    }
  };

  const createFilter = () => {
    if (formData.value.tagsDimension.length === undefined) {
      tagHeight[0] = 40;
    } else {
      tagHeight[formData.value.tagsDimension.length] = 40;
    }

    formData.value.tagsDimension.push({
      values: [
        {
          tag: '',
          operation: 'IN',
          value: [],
        },
      ],
    });
  };

  const changeTagFilter = (value: boolean | string | number) => {
    if (value) {
      showTagFilter.value = true;
      createFilter();
    } else {
      formData.value.tagsDimension = [];
      showTagFilter.value = false;
    }
  };
  const deleteFilterItem = (filterIndex: number, index: number) => {
    console.log(index);
    if (index === 0) {
      formData.value.tagsDimension[filterIndex].values.shift();
    } else {
      formData.value.tagsDimension[filterIndex].values.splice(index, 1);
    }
    if (formData.value.tagsDimension[filterIndex].values.length >= 2) {
      tagHeight[filterIndex] -= 40;
    }
    if (
      formData.value.tagsDimension[filterIndex].values.length === 0 &&
      formData.value.tagsDimension.length === 0
    ) {
      changeTagFilter(false);
    } else if (
      formData.value.tagsDimension[filterIndex].values.length === 0 &&
      formData.value.tagsDimension.length !== 0
    ) {
      if (filterIndex === 0) {
        formData.value.tagsDimension.shift();
      } else {
        formData.value.tagsDimension.splice(filterIndex);
      }
    }
  };

  watch(dimension, () => {
    if (dimension.value && !activeKeyRef.value.includes('2')) {
      activeKeyRef.value.push('2');
    } else {
      activeKeyRef.value = ['1'];
    }
  });
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
