<template>
  <a-select
    :placeholder="props.placeholder"
    :options="teamOptions"
    :multiple="props.multiple"
  >
    <template v-if="props.prefix" #prefix>
      <icon-user-group /> 告警策略
    </template>
  </a-select>
</template>

<script lang="ts" setup>
  import { SelectOptionData } from '@arco-design/web-vue/es/select/interface';
  import useRequest from '@/hooks/request';
  import { computed } from 'vue';
  import {
    SystemStrategyRecord,
    queryAllSystemStrategy,
  } from '@/api/system-strategy';

  export interface Props {
    multiple?: boolean;
    prefix?: boolean;
    placeholder?: string;
  }
  const props = withDefaults(defineProps<Props>(), {
    multiple: true,
    prefix: false,
    placeholder: '请选择全局告警策略',
  });
  const { response: data } = useRequest<SystemStrategyRecord[]>(
    queryAllSystemStrategy
  );

  const teamOptions = computed<SelectOptionData[]>(() => {
    if (!data.value) {
      return [];
    }

    const opts = data.value.map((item: SystemStrategyRecord) => ({
      label: `${item.strategyName}`,
      value: item.id,
    }));
    return opts;
  });
</script>
