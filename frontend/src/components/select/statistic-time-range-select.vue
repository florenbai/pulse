<template>
  <a-dropdown-button @select="changeFault">
    {{ fault.value }}
    <template #icon>
      <icon-down />
    </template>
    <template #content>
      <a-doption
        v-for="(name, index) of STATISTIC_TIME_RANGE_MAP"
        :key="index"
        :value="index"
        >{{ name }}</a-doption
      >
    </template>
  </a-dropdown-button>
</template>

<script lang="ts" setup>
  import { reactive } from 'vue';
  import { STATISTIC_TIME_RANGE_MAP } from '@/types/workspace';

  const props = defineProps({
    fault: {
      type: String,
      default: '最近7天',
    },
  });
  const fault = reactive({
    value: props.fault,
  });
  const emits = defineEmits(['changeValue']);

  const changeFault = (value: string) => {
    if (STATISTIC_TIME_RANGE_MAP[value] === fault.value) {
      return;
    }
    fault.value = STATISTIC_TIME_RANGE_MAP[value];
    emits('changeValue', value);
  };
</script>
