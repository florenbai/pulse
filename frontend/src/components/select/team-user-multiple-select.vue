<template>
  <a-select placeholder="请选择通知的团队成员" :options="teamOptions" multiple>
  </a-select>
</template>

<script lang="ts" setup>
  import { SelectOptionData } from '@arco-design/web-vue/es/select/interface';
  import useRequest from '@/hooks/request';
  import { computed } from 'vue';
  import { queryAllTeamUser, TeamUserRecord } from '@/api/team';

  export interface Props {
    team: number[];
  }
  const props = defineProps<Props>();
  const { response: data } = useRequest<TeamUserRecord[]>(
    queryAllTeamUser.bind(null, props.team)
  );

  const teamOptions = computed<SelectOptionData[]>(() => {
    if (data.value === undefined) {
      return [];
    }

    const opts = data.value.map((item: TeamUserRecord) => ({
      label: `${item.nickname}`,
      value: item.userId,
    }));
    return opts;
  });
</script>
