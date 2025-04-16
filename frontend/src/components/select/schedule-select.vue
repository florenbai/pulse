<template>
  <a-select
    :placeholder="props.placeholder"
    :options="scheduleOptions"
    :multiple="props.multiple"
  >
    <template v-if="props.prefix" #prefix> <icon-schedule /> 值班 </template>
  </a-select>
</template>

<script lang="ts" setup>
  import { SelectOptionData } from '@arco-design/web-vue/es/select/interface';
  import useRequest from '@/hooks/request';
  import { computed } from 'vue';
  import { queryAllSchedule, ScheduleRecord } from '@/api/schedule';

  export interface Props {
    multiple: boolean;
    prefix: boolean;
    placeholder: string;
  }
  const props = withDefaults(defineProps<Props>(), {
    multiple: true,
    prefix: false,
    placeholder: '请选择值班',
  });
  const { response: data } = useRequest<ScheduleRecord[]>(queryAllSchedule);

  const scheduleOptions = computed<SelectOptionData[]>(() => {
    if (!data.value) {
      return [];
    }

    const opts = data.value.map((item: ScheduleRecord) => ({
      label: `${item.scheduleName}`,
      value: item.id,
    }));
    return opts;
  });
</script>
