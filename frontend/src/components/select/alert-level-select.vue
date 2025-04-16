<template>
  <a-select
    placeholder="请选择告警等级"
    :options="levelOptions"
    :multiple="props.multiple"
    allow-clear
  >
  </a-select>
</template>

<script lang="ts" setup>
  export interface Props {
    multiple: boolean;
  }
  const props = defineProps<Props>();
  const levelOptions = [
    {
      label: 'A0',
      value: 'A0',
    },
    {
      label: 'A1',
      value: 'A1',
    },
    {
      label: 'A2',
      value: 'A2',
    },
    {
      label: 'A3',
      value: 'A3',
    },
    {
      label: 'A4',
      value: 'A4',
    },
    {
      label: 'A5',
      value: 'A5',
    },
  ];
</script>
