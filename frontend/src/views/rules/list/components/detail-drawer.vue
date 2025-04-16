<template>
  <a-drawer
    :width="800"
    :visible="visible"
    unmount-on-close
    @ok="handleOk"
    @cancel="handleCancel"
  >
    <template #title> 规则详情 </template>
    <a-descriptions
      style="margin-top: 20px"
      :column="1"
      :value-style="{
        'display': 'block',
        'white-space': 'normal',
        'word-wrap': 'break-word',
      }"
    >
      <a-descriptions-item label="告警规则标题">
        {{ alertRule?.name }}
      </a-descriptions-item>
      <a-descriptions-item label="数据源类型">
        {{ alertRule?.sourceType }}
      </a-descriptions-item>
      <a-descriptions-item label="查询语句">
        {{ alertRule?.query }}
      </a-descriptions-item>
      <a-descriptions-item label="异常触发告警时长">
        {{ secToFormat(alertRule?.duration as number) }}
      </a-descriptions-item>
      <a-descriptions-item label="状态">
        <a-tag v-if="alertRule?.health === 'ok'" color="green">{{
          alertRule?.health
        }}</a-tag>
        <a-tag v-else color="red">{{ alertRule?.health }}</a-tag>
      </a-descriptions-item>
      <a-descriptions-item label="标签">
        <a-tag
          v-for="(name, index) of alertRule?.labels"
          :key="index"
          color="arcoblue"
          style="margin: 2px"
        >
          {{ index }} : {{ name }}
        </a-tag>
      </a-descriptions-item>
      <a-descriptions-item label="注解">
        <a-tag
          v-for="(name, index) of alertRule?.annotations"
          :key="index"
          color="arcoblue"
          style="
            margin: 2px;
            white-space: normal;
            overflow-wrap: break-word;
            height: auto;
          "
        >
          {{ index }} : {{ name }}
        </a-tag>
      </a-descriptions-item>
    </a-descriptions>
  </a-drawer>
</template>

<script setup lang="ts">
  import { ref } from 'vue';
  import { AlertRuleRecode } from '@/api/alert-rule';
  import { secToFormat } from '@/utils/times';

  const visible = ref(false);
  const alertRule = ref<AlertRuleRecode>();

  const show = (item: AlertRuleRecode) => {
    alertRule.value = { ...item };
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
