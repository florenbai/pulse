<template>
  <a-select v-model="opd[0]" :options="options" style="width: 150px">
  </a-select>
  <a-select v-model="opd[1]" :options="periodOptions" style="width: 150px">
  </a-select>
</template>

<script lang="ts" setup>
  import { computed, onMounted, ref, watch } from 'vue';
  import { SelectOptionData } from '@arco-design/web-vue';

  export interface Props {
    modelValue: string[];
  }
  const opd = ref<string[]>(['1', 'day']);
  const emits = defineEmits(['update']);
  const props = defineProps<Props>();
  const options = computed(() => {
    const arr: SelectOptionData[] = [];
    /*
    for (let i = 1; i <= 60; i += 1) {
      arr.push({ value: String(i), label: String(i) });
    }
      */
    return arr;
  });

  watch(opd.value, () => {
    emits('update', opd.value);
  });

  onMounted(() => {
    if (props.modelValue.length > 0) {
      opd.value = props.modelValue;
    }
  });

  const periodOptions = [
    {
      label: '天',
      value: 'day',
    },
    /*
    {
      label: '周',
      value: 'week',
    },
    {
      label: '月',
      value: 'month',
    },
    */
  ];
</script>
