<template>
  <a-select
    placeholder="请选择渠道类型"
    :options="channelTypeOptions"
    allow-clear
  >
  </a-select>
</template>

<script lang="ts" setup>
  const channelTypeOptions = [
    {
      label: '企业微信',
      value: '企业微信',
    },
    {
      label: '企微群机器人',
      value: '企微群机器人',
    },
    {
      label: '电话',
      value: '电话',
    },
    {
      label: '邮箱',
      value: '邮箱',
    },
  ];
</script>
