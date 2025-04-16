<template>
  <a-drawer
    :width="600"
    :visible="visible"
    unmount-on-close
    :footer="!isReadonly"
    @ok="handleOk"
    @cancel="handleCancel"
  >
    <template #title> 全局告警通知策略 </template>

    <a-form
      ref="formRef"
      layout="vertical"
      :model="formData"
      style="margin-top: 20px"
    >
      <a-collapse :bordered="false" :active-key="['1', '2']">
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
            <alert-template-select
              v-model="formData.templateId"
              @change="fetchChannelGroup(formData.templateId)"
            />
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
            v-for="(item, index) of formData.strategySet"
            :key="index"
            :index="index"
            :strategy-set="item"
            :channel-group="channelGroupData"
            @handle-delete="handleDelete"
            @strategy-set-change="strategySetChange"
            @handle-add="handleAdd"
          ></assignment-strategy>
        </a-collapse-item>
      </a-collapse>
    </a-form>
  </a-drawer>
</template>

<script lang="ts" setup>
  import { ref } from 'vue';
  import {
    SystemStrategyRequest,
    submitSystemStrategyForm,
    editSystemStrategyForm,
    SystemStrategyRecord,
  } from '@/api/system-strategy';
  import { FormInstance, Message } from '@arco-design/web-vue';
  import { StrategyItem } from '@/api/alert-strategy';
  import assignmentStrategy from '@/views/alert/workspace/components/detail/strategy/assignment-strategy.vue';
  import alertTemplateSelect from '@/components/select/alert-template-select.vue';
  import {
    AlertChannelRecode,
    queryTemplateEnableChannel,
  } from '@/api/alert-template';

  const visible = ref(false);
  const editId = ref(0);
  const isReadonly = ref<boolean>(false);
  const formRef = ref<FormInstance>();
  const emits = defineEmits(['reload']);
  const channelGroupData = ref<AlertChannelRecode[]>([]);

  const formData = ref<SystemStrategyRequest>({
    strategyName: '',
    delay: 0,
    strategyDesc: '',
    templateId: 1,
    strategySet: [
      {
        teams: [],
        schedules: [],
        members: [],
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

  const fetchChannelGroup = async (id: number) => {
    try {
      const res = await queryTemplateEnableChannel(id);
      channelGroupData.value = res.data;
    } catch (err) {
      //
    }
  };

  const handleOk = async () => {
    const res = await formRef.value?.validate();
    if (!res) {
      try {
        if (editId.value !== 0) {
          await editSystemStrategyForm(editId.value, formData.value);
          Message.success('编辑成功');
        } else {
          await submitSystemStrategyForm(formData.value);
          Message.success('添加成功');
        }
        visible.value = false;
        emits('reload');
      } catch (err) {
        //
      }
    }
  };

  const handleCancel = () => {
    visible.value = false;
    formData.value = {
      strategyName: '',
      delay: 0,
      strategyDesc: '',
      templateId: 1,
      strategySet: [
        {
          teams: [],
          schedules: [],
          members: [],
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
    };
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

  const strategySetChange = (model: StrategyItem, index: number) => {
    formData.value.strategySet[index] = model;
  };

  const handleDelete = (index: number) => {
    formData.value.strategySet.splice(index, 1);
  };

  const show = (
    record: SystemStrategyRecord | undefined,
    readonly: boolean | undefined
  ) => {
    if (record !== undefined) {
      editId.value = record.id;
      formData.value = { ...record };
    }
    if (readonly !== undefined) {
      isReadonly.value = readonly;
    }
    visible.value = true;
    fetchChannelGroup(formData.value.templateId);
  };

  defineExpose({
    show,
  });
</script>

<style lang="less">
  .arco-collapse-item-content {
    background-color: #fff !important;
  }
</style>
