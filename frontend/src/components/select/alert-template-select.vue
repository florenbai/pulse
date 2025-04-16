<template>
  <a-select placeholder="请选择告警模板" :options="templateOptions"> </a-select>
</template>

<script lang="ts" setup>
  import { SelectOptionData } from '@arco-design/web-vue/es/select/interface';
  import useRequest from '@/hooks/request';
  import { computed } from 'vue';
  import { queryAllTemplate, AlertTemplateRecode } from '@/api/alert';

  const { response: data } =
    useRequest<AlertTemplateRecode[]>(queryAllTemplate);

  const templateOptions = computed<SelectOptionData[]>(() => {
    if (data.value === undefined) {
      return [];
    }

    const opts = data.value.map((item: AlertTemplateRecode) => ({
      label: `${item.templateName}`,
      value: item.id,
    }));
    return opts;
  });
</script>
