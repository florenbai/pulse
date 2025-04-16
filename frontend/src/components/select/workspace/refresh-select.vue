<template>
  <a-tooltip content="自动刷新设置">
    <a-dropdown-button @select="changeFault">
      {{ fault }}
      <template #icon>
        <icon-down />
      </template>

      <template #content>
        <a-doption
          v-for="(name, index) of REFRESH_TIME_MAP"
          :key="index"
          :value="index"
          >{{ name }}</a-doption
        >
      </template>
    </a-dropdown-button>
  </a-tooltip>
</template>

<script lang="ts" setup>
  import { ref } from 'vue';
  import { REFRESH_TIME_MAP } from '@/types/workspace';

  const emits = defineEmits(['refresh']);
  const fault = ref('10秒');
  const changeFault = (value: number) => {
    if (REFRESH_TIME_MAP[value] === fault.value) {
      return;
    }
    fault.value = REFRESH_TIME_MAP[value];
    emits('refresh', value);
  };
</script>
