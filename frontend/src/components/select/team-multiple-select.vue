<template>
  <a-select
    :placeholder="props.placeholder"
    :options="teamOptions"
    :multiple="props.multiple"
  >
    <template v-if="props.prefix" #prefix> <icon-user-group /> 团队 </template>
  </a-select>
</template>

<script lang="ts" setup>
  import { SelectOptionData } from '@arco-design/web-vue/es/select/interface';
  import useRequest from '@/hooks/request';
  import { computed } from 'vue';
  import { queryAllTeam, TeamRecord } from '@/api/team';

  export interface Props {
    multiple?: boolean;
    prefix?: boolean;
    placeholder?: string;
  }
  const props = withDefaults(defineProps<Props>(), {
    multiple: true,
    prefix: false,
    placeholder: '请选择团队',
  });
  const { response: data } = useRequest<TeamRecord[]>(queryAllTeam);

  const teamOptions = computed<SelectOptionData[]>(() => {
    if (!data.value) {
      return [];
    }

    const opts = data.value.map((item: TeamRecord) => ({
      label: `${item.teamName}`,
      value: item.id,
    }));
    return opts;
  });
</script>
