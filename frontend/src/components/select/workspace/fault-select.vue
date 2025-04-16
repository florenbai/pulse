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

  export interface Props {
    value: number;
  }
  const props = withDefaults(defineProps<Props>(), {
    value: 0,
  });

  const doptions = ['全部告警状态', '未认领', '已认领', '已关闭', '未关闭'];
  const fault = ref(doptions[props.value]);
  const emits = defineEmits(['changeValue']);
  const changeFault = (value: number) => {
    if (doptions[value] === fault.value) {
      return;
    }
    fault.value = doptions[value];
    emits('changeValue', value);
  };
</script>
