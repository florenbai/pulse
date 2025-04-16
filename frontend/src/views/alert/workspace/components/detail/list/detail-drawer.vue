<template>
  <a-drawer
    :width="800"
    :visible="visible"
    unmount-on-close
    @ok="handleOk"
    @cancel="handleCancel"
  >
    <template #title> 告警详情 </template>
    <a-descriptions style="margin-top: 20px" :column="1">
      <a-descriptions-item label="告警标题">
        {{ alertEvent?.alertTitle }}
      </a-descriptions-item>
      <a-descriptions-item label="告警源">
        <a-tooltip :content="alertEvent?.sourceName">
          <icon-font :type="alertEvent?.icon" :size="24" />
        </a-tooltip>
      </a-descriptions-item>
      <a-descriptions-item label="告警等级">
        <a-tag :color="alertEvent?.color"> {{ alertEvent?.levelName }}</a-tag>
      </a-descriptions-item>
      <a-descriptions-item label="告警来源">
        {{ alertEvent?.sourceIp }}
      </a-descriptions-item>
      <a-descriptions-item label="告警描述">
        {{ alertEvent?.description }}
      </a-descriptions-item>
      <a-descriptions-item label="首次告警时间">
        {{ alertEvent?.firstTriggerTime }}
      </a-descriptions-item>
      <a-descriptions-item label="最新告警时间">
        {{ alertEvent?.triggerTime }}
      </a-descriptions-item>
      <a-descriptions-item label="告警次数">
        {{ alertEvent?.notifyCurNumber }}
      </a-descriptions-item>
      <a-descriptions-item label="是否恢复">
        <a-tag v-if="alertEvent?.isRecovered" color="green">是</a-tag>
        <a-tag v-else color="red">否</a-tag>
      </a-descriptions-item>
      <a-descriptions-item v-if="alertEvent?.isRecovered" label="恢复时间">
        {{ alertEvent?.recoverTime }}
      </a-descriptions-item>
      <a-descriptions-item label="告警标签">
        <a-tag
          v-for="(name, index) of alertEvent?.tags"
          :key="index"
          color="arcoblue"
          style="margin: 2px"
        >
          {{ index }} : {{ name }}
        </a-tag>
      </a-descriptions-item>
      <a-descriptions-item label="告警注解">
        <a-tag
          v-for="(name, index) of alertEvent?.annotations"
          :key="index"
          color="arcoblue"
        >
          {{ index }} : {{ name }}
        </a-tag>
      </a-descriptions-item>
      <a-descriptions-item label="处理进度">
        <a-tag :color="PROGRESS_COLOR_MAP[alertEvent?.progress as number]">{{
          PROGRESS_MAP[alertEvent?.progress as number]
        }}</a-tag>
      </a-descriptions-item>
      <a-descriptions-item label="最新操作人">
        {{ alertEvent?.nickname }}
      </a-descriptions-item>
    </a-descriptions>
    <a-tabs class="tabs-custom" type="line">
      <a-tab-pane key="1">
        <template #title> 告警日志 </template>
        <a-timeline>
          <a-timeline-item
            :label="alertEvent?.firstTriggerTime"
            dot-color="#F53F3F"
          >
            告警触发
            <a-tag color="arcoblue" style="margin-left: 20px">系统</a-tag>
          </a-timeline-item>
          <a-timeline-item
            v-for="(item, index) of alertLog"
            :key="index"
            :label="item.createdAt"
            :dot-color="ALERT_ACTION_COLOR_MAP[item.action]"
          >
            {{ ALERT_ACTION_MAP[item.action] }}
            <a-tag color="arcoblue" style="margin-left: 20px">{{
              item.nickname
            }}</a-tag>
          </a-timeline-item>
        </a-timeline>
      </a-tab-pane>
      <a-tab-pane key="2">
        <template #title> 告警通知日志 </template>
        <a-timeline>
          <a-timeline-item
            v-for="(item, index) of strategyLog"
            :key="index"
            :label="item.createdAt"
            :dot-color="item.notifyType === 2 ? '#00b42a' : '#F53F3F'"
            :style="{ marginTop: '6px' }"
          >
            匹配策略
            <a-tag style="margin-left: 20px">
              {{ item.strategyTitle }}
            </a-tag>
            <br />
            <div style="margin-top: 10px">
              告警渠道
              <a-tag
                v-for="(cv, ck) of item.channels"
                :key="ck"
                style="margin-left: 20px"
                color="arcoblue"
              >
                {{ cv }}
              </a-tag>
            </div>
            <div style="margin-top: 10px">
              通知人员
              <a-tag
                v-for="(cv, ck) of item.nicknames"
                :key="ck"
                style="margin-left: 20px"
                color="arcoblue"
              >
                {{ cv }}
              </a-tag>
            </div>
            <a-typography-text
              type="secondary"
              :style="{ fontSize: '12px', marginTop: '4px' }"
            >
              {{ item.errMessage }}
            </a-typography-text>
          </a-timeline-item>
        </a-timeline>
      </a-tab-pane>
    </a-tabs>
  </a-drawer>
</template>

<script setup lang="ts">
  import { ref } from 'vue';
  import { AlertEventRecode } from '@/api/alert-event';
  import {
    PROGRESS_MAP,
    PROGRESS_COLOR_MAP,
    ALERT_ACTION_MAP,
    ALERT_ACTION_COLOR_MAP,
  } from '@/types/workspace';
  import {
    AlertLogRecord,
    StrategyLogRecord,
    queryAlertLog,
    queryStrategyLog,
  } from '@/api/alert-log';

  const visible = ref(false);
  const alertEvent = ref<AlertEventRecode>();
  const alertLog = ref<AlertLogRecord[]>([]);
  const strategyLog = ref<StrategyLogRecord[]>([]);

  const fetchData = async () => {
    try {
      if (alertEvent.value?.id !== undefined) {
        const { data } = await queryAlertLog(alertEvent.value?.id);
        alertLog.value = data;
      }
    } catch (err) {
      //
    }
  };

  const fetchStrategyData = async () => {
    try {
      if (alertEvent.value?.id !== undefined) {
        const { data } = await queryStrategyLog(alertEvent.value?.id);
        strategyLog.value = data;
      }
    } catch (err) {
      //
    }
  };

  const show = (item: AlertEventRecode) => {
    alertEvent.value = { ...item };
    fetchData();
    fetchStrategyData();
    visible.value = true;
  };
  const handleOk = () => {
    visible.value = false;
  };
  const handleCancel = () => {
    visible.value = false;
  };

  defineExpose({
    show,
  });
</script>

<style scoped lang="less">
  .tabs-custom {
    padding-top: 20px;
  }
  .tabs-custom ::before {
    height: 0px !important;
  }
</style>
