<template>
  <a-select
    placeholder="请选中聚合维度"
    :options="dimensionOptions"
    multiple
    allow-clear
  >
  </a-select>
</template>

<script lang="ts" setup>
  const dimensionOptions = [
    {
      label: '故障标题',
      value: 1,
    },
    {
      label: '故障等级',
      value: 2,
    },
    {
      label: '故障标签',
      value: 3,
    },
  ];
</script>
