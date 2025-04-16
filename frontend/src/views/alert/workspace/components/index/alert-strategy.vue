<template>
  <a-form ref="formRef" :model="formData" layout="vertical">
    <a-collapse :default-active-key="['1', '2']" :bordered="false">
      <a-collapse-item key="1" header="基础信息" style="padding-left: 0px">
        <a-form-item
          style="margin-top: 10px"
          field="strategyName"
          label="策略名称"
          validate-trigger="input"
          :rules="[
            {
              required: true,
              message: '请输入策略名称',
            },
          ]"
        >
          <a-input
            v-model="formData.strategyName"
            placeholder="请输入策略名称"
          />
        </a-form-item>
        <a-form-item
          style="margin-top: 10px"
          field="strategyDesc"
          label="策略描述"
          validate-trigger="input"
        >
          <a-input
            v-model="formData.strategyDesc"
            placeholder="请输入策略描述"
          />
        </a-form-item>
        <a-form-item
          field="templateId"
          label="告警模板"
          :rules="[
            {
              required: true,
              message: '请选择告警模板',
            },
          ]"
        >
          <alert-template-select v-model="formData.templateId" />
        </a-form-item>
        <a-form-item
          style="margin-top: 10px"
          field="delay"
          label="通知延迟"
          validate-trigger="input"
          :rules="[
            {
              required: true,
              message: '请输入通知延迟',
            },
          ]"
        >
          <a-input-number
            v-model="formData.delay"
            placeholder="请输入通知延迟时间"
          >
            <template #append> 秒 </template>
          </a-input-number>

          <template #extra>
            <div style="line-height: 20px"
              >0表示实时通知。如果在等待期内故障自动恢复或被认领，则通知不会被发送。</div
            >
          </template>
        </a-form-item>
      </a-collapse-item>
      <a-collapse-item key="2" header="分派策略" style="padding-left: 0px">
        <assignment-strategy
          v-for="(post, index) of formData.strategySet"
          :key="index"
          :index="index"
          :team="props.team"
          @handle-delete="handleDelete"
          @strategy-set-change="strategySetChange"
          @handle-add="handleAdd"
        ></assignment-strategy>
      </a-collapse-item>
    </a-collapse>
  </a-form>
</template>

<script lang="ts" setup>
  import { ref } from 'vue';
  import { FormInstance } from '@arco-design/web-vue/es/form';
  import { AlertStrategyRequest, StrategyItem } from '@/api/workspace';
  import alertTemplateSelect from '@/components/select/alert-template-select.vue';
  import assignmentStrategy from './assignment-strategy.vue';

  export interface Props {
    team: number[];
  }
  const props = defineProps<Props>();
  const emits = defineEmits(['changeStep']);
  const formRef = ref<FormInstance>();
  const formData = ref<AlertStrategyRequest>({
    strategyName: '',
    delay: 0,
    strategyDesc: '',
    templateId: 1,
    timeSlot: {
      enable: false,
    },
    filters: [],
    strategySet: [
      {
        members: [],
        teams: [],
        schedules: [],
        alertChannels: [],
        alertLoop: {
          enable: false,
          interval: 0,
          count: 1,
        },
        progression: {
          enable: false,
          condition: 0,
          duration: 1,
        },
      },
    ],
  });

  const strategySetChange = (model: StrategyItem, index: number) => {
    formData.value.strategySet[index] = model;
  };

  const submitAlertStrategy = async (): Promise<boolean> => {
    const res = await formRef.value?.validate();
    if (!res) {
      emits('changeStep', { ...formData.value });
      return true;
    }
    return false;
  };

  const handleAdd = () => {
    formData.value.strategySet.push({
      members: [],
      teams: [],
      schedules: [],
      alertChannels: [],
      alertLoop: {
        enable: false,
        interval: 0,
        count: 1,
      },
      progression: {
        enable: false,
        condition: 0,
        duration: 1,
      },
    });
  };

  const handleDelete = (index: number) => {
    formData.value.strategySet.splice(index, 1);
  };

  defineExpose({
    submitAlertStrategy,
  });
</script>

<style>
  .arco-collapse-item-content {
    background-color: #fff !important;
  }

  .arco-collapse-item-header {
    border: none !important;
  }
</style>
