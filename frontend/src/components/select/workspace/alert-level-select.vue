<template>
  <a-dropdown-button @select="changeFault">
    {{ fault }}
    <template #icon>
      <icon-down />
    </template>
    <template #content>
      <a-doption
        v-for="(item, index) in doptions"
        :key="index"
        :value="index"
        >{{ item }}</a-doption
      >
    </template>
  </a-dropdown-button>
</template>

<script lang="ts" setup>
  import { ref } from 'vue';

  const fault = ref('全部告警等级');
  const doptions = ['全部告警等级', 'A0', 'A1', 'A2', 'A3', 'A4', 'A5', 'A6'];
  const emits = defineEmits(['changeValue']);
  const changeFault = (value: number) => {
    if (doptions[value] === fault.value) {
      return;
    }
    fault.value = doptions[value];
    emits('changeValue', value);
  };
</script>
