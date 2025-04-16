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

  const fault = ref('全部故障');
  const doptions = [
    '全部故障',
    '待处理故障',
    '处理中故障',
    '未关闭故障',
    '已关闭故障',
  ];

  const changeFault = (value: number) => {
    if (doptions[value] === fault.value) {
      return;
    }
    fault.value = doptions[value];
  };
</script>
