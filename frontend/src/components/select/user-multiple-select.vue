<template>
  <a-select
    :placeholder="props.placeholder"
    :options="userOptions"
    :multiple="props.multiple"
    allow-search
  >
    <template v-if="props.prefix" #prefix> <icon-user /> 个人 </template>
  </a-select>
</template>

<script lang="ts" setup>
  import { SelectOptionData } from '@arco-design/web-vue/es/select/interface';
  import useRequest from '@/hooks/request';
  import { computed } from 'vue';
  import { queryAllUser, UserRecord } from '@/api/user';

  export interface Props {
    multiple: boolean;
    prefix: boolean;
    placeholder: string;
  }
  const props = withDefaults(defineProps<Props>(), {
    multiple: true,
    prefix: false,
    placeholder: '请选择用户',
  });
  const { response: data } = useRequest<UserRecord[]>(queryAllUser);

  const userOptions = computed<SelectOptionData[]>(() => {
    if (data.value === undefined) {
      return [];
    }

    const opts = data.value.map((item: UserRecord) => ({
      label: `${item.nickname}`,
      value: item.id,
    }));
    return opts;
  });
</script>
