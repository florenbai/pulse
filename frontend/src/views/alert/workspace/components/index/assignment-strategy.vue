<template>
  <a-card :title="`策略${props.index + 1}`" style="margin-bottom: 15px">
    <a-form-item
      style="margin-top: 10px"
      label="通知人"
      :field="`strategySet[${props.index}].members`"
      :rules="[
        {
          required: true,
          message: '请选择通知人',
        },
      ]"
    >
      <team-user-multiple-select
        v-model="strategySet.members"
        :team="props.team"
      />
    </a-form-item>
    <a-form-item
      style="margin-top: 10px"
      label="通知渠道"
      :field="`strategySet[${props.index}].alertChannels`"
      :rules="[
        {
          required: true,
          message: '请选择通知渠道',
        },
      ]"
    >
      <notify-type-select
        v-model="strategySet.alertChannels"
        :multiple="true"
      ></notify-type-select>
    </a-form-item>
    <a-form-item
      style="margin-top: 10px"
      label="告警等级"
      :field="`strategySet[${props.index}].alertLevel`"
      :rules="[
        {
          required: true,
          message: '请选择告警等级',
        },
      ]"
    >
      <alert-level-select
        v-model="strategySet.alertChannels"
        :multiple="true"
      ></alert-level-select>
    </a-form-item>
    <a-form-item style="margin-top: 10px" label="循环通知">
      <a-switch
        v-model="strategySet.alertLoop.enable"
        :before-change="handleChangeIntercept"
      />
    </a-form-item>
    <a-form-item
      v-if="strategySet.alertLoop.enable"
      style="margin-top: 10px"
      label="循环通知间隔"
    >
      <a-input-number
        v-model="strategySet.alertLoop.interval"
        placeholder="请输入循环通知间隔，单位分钟"
        :field="`strategySet[${props.index}].alertLoop.interval`"
        :rules="[
          {
            required: true,
            type: 'number',
            min: 1,
            message: '循环通知间隔最小为1分钟',
          },
        ]"
      />
      <span style="margin-left: 10px; width: 50px">分钟</span>
    </a-form-item>
    <a-form-item
      v-if="strategySet.alertLoop.enable"
      style="margin-top: 10px"
      label="循环通知次数"
    >
      <a-input-number
        v-model="strategySet.alertLoop.count"
        placeholder="请输入循环通知次数，开启后最小值为1次"
        :field="`strategySet[${props.index}].alertLoop.count`"
        :rules="[
          {
            required: true,
            type: 'number',
            min: 1,
            message: '循环通知次数最小为1次',
          },
        ]"
      />
    </a-form-item>
    <a-form-item style="margin-top: 10px" label="故障升级">
      <a-switch
        v-model="strategySet.progression.enable"
        :before-change="handleUpgrade"
      />
    </a-form-item>
    <a-form-item
      v-if="strategySet.progression.enable"
      style="margin-top: 10px"
      label="触发升级条件"
    >
      <a-radio-group
        v-model="strategySet.progression.condition"
        type="button"
        default-value="0"
      >
        <a-radio :value="0">未认领</a-radio>
        <a-radio :value="1">未关闭</a-radio>
      </a-radio-group>
    </a-form-item>
    <a-form-item
      v-if="strategySet.progression.enable"
      style="margin-top: 10px"
      label="触发升级时间"
    >
      <a-input-number
        v-model="strategySet.progression.duration"
        placeholder="请输入故障升级触发时间，单位为分钟。该值必须大于0"
      />
    </a-form-item>
  </a-card>
</template>

<script lang="ts" setup>
  import { ref, watch } from 'vue';
  import teamUserMultipleSelect from '@/components/select/team-user-multiple-select.vue';
  import alertLevelSelect from '@/components/select/alert-level-select.vue';
  import notifyTypeSelect from '@/components/select/notify-type-select.vue';
  import { StrategyItem } from '@/api/workspace';

  export interface Props {
    team: number[];
    index: number;
  }
  const props = defineProps<Props>();
  const strategySet = ref<StrategyItem>({
    teams: [],
    schedules: [],
    members: [],
    alertChannels: [],
    alertLoop: {
      enable: false,
      interval: 1,
      count: 1,
    },
    progression: {
      enable: false,
      condition: 0,
      duration: 1,
    },
  });
  const emits = defineEmits(['strategySetChange', 'handleDelete', 'handleAdd']);
  const handleChangeIntercept = () => {
    if (strategySet.value.alertLoop.enable) {
      strategySet.value.alertLoop.enable = false;
    } else {
      strategySet.value.alertLoop.enable = true;
    }
    return true;
  };

  watch(strategySet.value, () => {
    emits('strategySetChange', strategySet.value, props.index);
  });

  const handleUpgrade = () => {
    if (strategySet.value.progression.enable) {
      strategySet.value.progression.enable = false;
      emits('handleDelete', props.index);
    } else {
      emits('handleAdd');
      strategySet.value.progression.enable = true;
    }
    return true;
  };
</script>
