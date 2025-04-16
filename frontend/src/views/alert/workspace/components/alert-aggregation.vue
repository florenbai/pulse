<template>
  <a-form ref="formRef" :model="formData">
    <a-collapse
      :default-active-key="['1', '2']"
      :bordered="false"
      destroy-on-hide
    >
      <a-collapse-item
        key="1"
        header="告警聚合设置"
        style="padding-left: 0px; border: none"
      >
        <template #extra>
          <a-tooltip
            content="当告警的标题或者属性与某个未关闭故障完全相同时，将告警聚合，以达到减少通知的目的。"
          >
            <a-switch
              :before-change="handleChangeAlertAggregation"
              :model-value="formData.alertAggregation.enable"
            />
          </a-tooltip>
        </template>
        <div v-if="formData.alertAggregation.enable">
          <a-form-item
            style="margin-top: 10px"
            field="alertAggregation.dimension"
            label="聚合维度"
            :rules="[
              {
                required: true,
                message: '请选中聚合维度',
              },
            ]"
          >
            <dimension-select
              v-model="formData.alertAggregation.dimension"
            ></dimension-select>
            <template #extra>
              <div style="line-height: 20px">
                当告警维度完全匹配时，才会进行告警聚合操作。
              </div>
            </template>
          </a-form-item>
          <a-form-item
            style="margin-top: 10px"
            field="alertAggregation.window"
            label="聚合窗口"
          >
            <a-input-number
              v-model="formData.alertAggregation.window"
              placeholder="请输入聚合窗口，单位分钟"
            />
            <span style="margin-left: 10px; width: 30px">分钟</span>
            <template #extra>
              <div style="line-height: 20px">
                超出时间范围后的告警将被聚合到新的故障，并进行通知。默认为0，表示不设置聚合窗口。
              </div>
            </template>
          </a-form-item>
          <a-form-item
            style="margin-top: 10px"
            field="alertAggregation.stormWatch"
            label="风暴预警"
            :rules="[
              {
                required: true,
                message: '请输入风暴预警',
              },
            ]"
          >
            <a-input-number
              v-model="formData.alertAggregation.stormWatch"
              placeholder="请输入风暴预警条数"
            />
            <template #extra>
              <div style="line-height: 20px">
                当聚合的告警达到设定次数，则进行再次通知。默认为0，表示不进行风暴预警
              </div>
            </template>
          </a-form-item>
        </div>
      </a-collapse-item>
      <a-collapse-item
        key="2"
        header="告警静默设置"
        style="padding-left: 0px; border: none"
      >
        <template #extra>
          <a-tooltip
            content="设置告警静默后，告警信息将被标记为静默，不计入告警统计，并且不会受到告警通知。"
          >
            <a-switch
              :before-change="handleChangeIntercept"
              :model-value="loop"
          /></a-tooltip>
        </template>
        <div v-if="loop">
          <a-form-item
            style="margin-top: 10px"
            field="strategyName"
            label="静默时间"
          >
            <a-radio-group type="button" default-value="按日期">
              <a-radio value="按日期">按日期</a-radio>
              <a-radio value="按周期">按周期</a-radio>
            </a-radio-group>
          </a-form-item>
          <a-form-item style="margin-top: 10px">
            <a-range-picker
              style="width: 360px; margin: 0 24px 24px 0"
              show-time
              :time-picker-props="{ defaultValue: ['00:00:00', '09:09:06'] }"
              format="YYYY-MM-DD HH:mm"
              @ok="onOk"
            />
          </a-form-item>
          <a-form-item
            style="margin-top: 10px"
            field="delay"
            label="静默条件设置"
            validate-trigger="input"
          >
            <a-switch
              :before-change="handleChangeIntercept"
              :model-value="loop"
            />
          </a-form-item>
        </div>
        <div v-if="loop">
          <a-card hoverable :style="{ marginBottom: '20px' }">
            <a-form-item
              v-for="(post, index) of form"
              :key="index"
              style="display: block; margin-bottom: 20px"
            >
              <a-select
                v-model="post.name"
                :style="{ width: '150px' }"
                placeholder="请选择条件"
                allow-clear
              >
                <a-option>告警等级</a-option>
                <a-option>告警标签</a-option>
                <a-option>告警内容</a-option>
              </a-select>
              <a-dropdown>
                <a-button style="border-color: rgb(var(--primary-6))"
                  >等于</a-button
                >
                <template #content>
                  <a-doption>等于</a-doption>
                  <a-doption>不等于</a-doption>
                  <a-doption>正则匹配</a-doption>
                </template>
              </a-dropdown>

              <a-input
                v-model="post.value"
                placeholder="请输入匹配项，将采用正则表达式进行匹配"
                :style="{ width: '400px' }"
              />
              <icon-minus-circle
                size="20"
                style="margin-left: 20px"
                @click="deleteConditionItem"
              />
            </a-form-item>
            <div style="margin-top: 20px">
              <a-button @click="handleAddConditionItem">新增条件项</a-button>
            </div>
          </a-card>
          <a-button type="dashed" long>新增条件</a-button>
        </div>
      </a-collapse-item>
    </a-collapse>
  </a-form>
</template>

<script lang="ts" setup>
  import { ref } from 'vue';
  import { FormInstance } from '@arco-design/web-vue/es/form';
  import { DownFrequencyRequest } from '@/api/workspace';
  import dimensionSelect from '@/components/select/dimension-select.vue';

  const emits = defineEmits(['changeStep']);
  const form = ref([{ value: '', name: '' }]);
  const formRef = ref<FormInstance>();
  const formData = ref<DownFrequencyRequest>({
    alertAggregation: {
      enable: true,
      dimension: [],
      window: 1,
      stormWatch: 1,
    },
    alertQuiet: {
      enable: true,
      quietType: '',
      quietStartTime: '',
      quietEndTime: '',
      condition: {
        enable: false,
        gather: [],
      },
    },
  });

  const loop = ref(true);
  const upgrade = ref(false);
  const handleChangeIntercept = () => {
    if (loop.value) {
      loop.value = false;
    } else {
      loop.value = true;
    }
    return true;
  };

  const handleChangeAlertAggregation = () => {
    if (formData.value.alertAggregation.enable) {
      formData.value.alertAggregation.enable = false;
    } else {
      formData.value.alertAggregation.enable = true;
    }
    return true;
  };

  const submitBaseInfo = async (): Promise<boolean> => {
    const res = await formRef.value?.validate();
    if (!res) {
      emits('changeStep', { ...formData.value });
      return true;
    }
    return false;
  };

  const onOk = (dateString: any, date: (Date | undefined)[]) => {
    console.log('onSelect', dateString, date);
  };

  const handleAddConditionItem = () => {
    form.value.push({
      name: '',
      value: '',
    });
  };

  const deleteConditionItem = (index: any) => {
    form.value.splice(index, 1);
  };
  defineExpose({
    submitBaseInfo,
  });
</script>

<style>
  .arco-collapse-item-content {
    background-color: #fff !important;
  }
</style>
