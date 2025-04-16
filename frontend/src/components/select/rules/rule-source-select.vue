<template>
  <a-select placeholder="请选择数据源" :options="sourceOptions"> </a-select>
</template>

<script lang="ts" setup>
  import { SelectOptionData } from '@arco-design/web-vue/es/select/interface';
  import useRequest from '@/hooks/request';
  import { computed } from 'vue';
  import {
    AlertRuleSourceRecode,
    queryAllAlertRuleSource,
  } from '@/api/alert-rule';

  const { response: data } = useRequest<AlertRuleSourceRecode[]>(
    queryAllAlertRuleSource
  );

  const sourceOptions = computed<SelectOptionData[]>(() => {
    if (data.value === undefined) {
      return [];
    }

    const opts = data.value.map((item: AlertRuleSourceRecode) => ({
      label: `${item.sourceName}`,
      value: item.id,
    }));
    return opts;
  });
</script>
