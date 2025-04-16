<template>
  <a-select
    placeholder="请选择工作区"
    :options="workspaceOptions"
    :multiple="props.multiple"
    allow-clear
  >
  </a-select>
</template>

<script lang="ts" setup>
  import { queryWorkspaceAll, WorkspaceRecode } from '@/api/workspace';
  import useRequest from '@/hooks/request';
  import { SelectOptionData } from '@arco-design/web-vue';
  import { computed } from 'vue';

  export interface Props {
    multiple: boolean;
  }
  const props = defineProps<Props>();
  const { response: data } = useRequest<WorkspaceRecode[]>(queryWorkspaceAll);

  const workspaceOptions = computed<SelectOptionData[]>(() => {
    if (!data.value) {
      return [];
    }

    const opts = data.value.map((item: WorkspaceRecode) => ({
      label: `${item.workspaceName}`,
      value: item.id,
    }));
    return opts;
  });
</script>
