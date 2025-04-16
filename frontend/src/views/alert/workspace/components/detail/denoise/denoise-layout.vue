<template>
  <a-space direction="vertical" fill style="height: 100%">
    <aggregation :id="props.id" />
    <silence :id="props.id" />
    <restrain :id="props.id"></restrain>
  </a-space>
</template>

<script lang="ts" setup>
  import aggregation from './aggregation.vue';
  import silence from './silence.vue';
  import restrain from './restrain.vue';

  export interface Props {
    id: number | undefined;
  }
  const props = defineProps<Props>();
</script>
