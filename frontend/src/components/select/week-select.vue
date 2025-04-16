<template>
  <a-select placeholder="请选择周期" :options="weekOptions" allow-clear>
  </a-select>
</template>

<script lang="ts" setup>
  const weekOptions = [
    {
      label: '周一',
      value: '周一',
    },
    {
      label: '周二',
      value: '周二',
    },
    {
      label: '周三',
      value: '周三',
    },
    {
      label: '周四',
      value: '周四',
    },
    {
      label: '周五',
      value: '周五',
    },
    {
      label: '周六',
      value: '周六',
    },
    {
      label: '周日',
      value: '周日',
    },
  ];
</script>
