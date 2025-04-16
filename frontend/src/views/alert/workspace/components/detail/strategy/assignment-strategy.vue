<template>
  <a-card :title="`策略${props.index + 1}`" style="margin-bottom: 15px">
    <a-form-item
      style="margin-top: 10px"
      label="通知配置"
      :rules="[
        {
          required: true,
          message: '请选择通知人',
        },
      ]"
    >
      <a-input-group v-if="personShow" style="width: 100%; margin-bottom: 10px">
        <userMultipleSelect
          v-model="strategySet.members"
          :prefix="true"
          :multiple="true"
          :placeholder="'请选择通知人'"
        ></userMultipleSelect>

        <div>
          <a-button @click="personShowFunc">
            <template #icon>
              <icon-close />
            </template>
          </a-button>
        </div>
      </a-input-group>
      <a-input-group v-if="teamShow" style="width: 100%; margin-bottom: 10px">
        <teamMultipleSelect
          v-model="strategySet.teams"
          :prefix="true"
          :multiple="true"
          :placeholder="'请选择通知团队'"
        ></teamMultipleSelect>
        <div>
          <a-button @click="teamShowFunc">
            <template #icon>
              <icon-close />
            </template>
          </a-button>
        </div>
      </a-input-group>
      <a-input-group
        v-if="scheduleShow"
        style="width: 100%; margin-bottom: 10px"
      >
        <scheduleSelect
          v-model="strategySet.schedules"
          :multiple="true"
          :prefix="true"
          :placeholder="'请选择值班'"
        ></scheduleSelect>

        <div>
          <a-button @click="scheduleShowFunc">
            <template #icon>
              <icon-close />
            </template>
          </a-button>
        </div>
      </a-input-group>
      <div>
        <a-space>
          <a-button v-if="!personShow" @click="personShowFunc">
            <template #icon>
              <icon-user />
            </template>
            个人通知配置
          </a-button>
          <a-button v-if="!teamShow" @click="teamShowFunc">
            <template #icon>
              <icon-user-group />
            </template>
            团队通知配置
          </a-button>
          <a-button v-if="!scheduleShow" @click="scheduleShowFunc">
            <template #icon>
              <icon-schedule />
            </template>
            值班通知配置
          </a-button>
        </a-space>
      </div>
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
      <a-table
        :columns="columns"
        :data="levelList"
        :pagination="false"
        style="width: 100%"
        column-resizable
      >
        <template #levelName="{ rowIndex }">
          <a-tag :color="levelList[rowIndex].color">{{
            levelList[rowIndex].levelName
          }}</a-tag>
        </template>
        <template #notifyChannel="{ rowIndex }">
          <templateNotifySelect
            v-model="strategySet.alertChannels[rowIndex]"
            :channel-group="props.channelGroup"
            :multiple="true"
          />
        </template>
      </a-table>
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
      >
        <template #append> 分钟 </template>
      </a-input-number>
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
      >
        <template #append> 次 </template>
      </a-input-number>
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
      >
        <template #append> 分钟 </template>
      </a-input-number>
    </a-form-item>
  </a-card>
</template>

<script lang="ts" setup>
  import { onMounted, ref, watch } from 'vue';
  import teamMultipleSelect from '@/components/select/team-multiple-select.vue';
  import userMultipleSelect from '@/components/select/user-multiple-select.vue';
  import scheduleSelect from '@/components/select/schedule-select.vue';
  import templateNotifySelect from '@/components/select/template-notify-select.vue';
  import { StrategyItem } from '@/api/workspace';
  import { AlertLevelRecode, queryAllAlertLevel } from '@/api/alert-level';
  import { AlertChannelRecode } from '@/api/alert-template';

  export interface Props {
    index: number;
    strategySet: StrategyItem;
    channelGroup: AlertChannelRecode[];
  }
  const props = defineProps<Props>();
  const personShow = ref(false);
  const teamShow = ref(false);
  const scheduleShow = ref(false);
  const strategySet = ref<StrategyItem>(props.strategySet);
  const levelList = ref<AlertLevelRecode[]>([]);
  const emits = defineEmits(['strategySetChange', 'handleDelete', 'handleAdd']);
  const handleChangeIntercept = () => {
    if (strategySet.value.alertLoop.enable) {
      strategySet.value.alertLoop.enable = false;
    } else {
      strategySet.value.alertLoop.enable = true;
    }
    return true;
  };

  const columns = [
    {
      title: '告警等级',
      dataIndex: 'levelName',
      slotName: 'levelName',
      width: 90,
    },
    {
      title: '通知渠道',
      dataIndex: 'notifyChannel',
      slotName: 'notifyChannel',
    },
  ];
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

  const personShowFunc = () => {
    if (personShow.value) {
      personShow.value = false;
      strategySet.value.members = [];
    } else {
      personShow.value = true;
    }
  };

  const teamShowFunc = () => {
    if (teamShow.value) {
      teamShow.value = false;
      strategySet.value.teams = [];
    } else {
      teamShow.value = true;
    }
  };

  const scheduleShowFunc = () => {
    if (scheduleShow.value) {
      scheduleShow.value = false;
      strategySet.value.schedules = [];
    } else {
      scheduleShow.value = true;
    }
  };

  const getAllLevel = async () => {
    const { data } = await queryAllAlertLevel();
    levelList.value = data;
    if (levelList.value.length < strategySet.value.alertChannels.length) {
      for (
        let i = 0;
        i < strategySet.value.alertChannels.length - levelList.value.length;
        i += 1
      ) {
        strategySet.value.alertChannels.pop();
      }
    }
  };

  onMounted(() => {
    if (strategySet.value.members.length > 0) {
      personShow.value = true;
    }
    if (strategySet.value.teams.length > 0) {
      teamShow.value = true;
    }
    if (strategySet.value.schedules.length > 0) {
      scheduleShow.value = true;
    }
    getAllLevel();
  });
</script>

<style lang="less">
  .arco-form-item-content,
  .arco-form-item-content-flex {
    flex-wrap: wrap !important;
  }

  .arco-select-view-prefix {
    border-right: 1px solid var(--color-neutral-3) !important;
    margin-right: 12px;
  }
</style>
