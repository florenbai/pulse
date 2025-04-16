<template>
  <a-select
    placeholder="请选择通知渠道"
    :options="wechatOptions"
    :multiple="props.multiple"
    allow-clear
  >
  </a-select>
</template>

<script lang="ts" setup>
  import { AlertChannelRecode, queryChannelGroup } from '@/api/alert-template';
  import useRequest from '@/hooks/request';
  import { SelectOptionGroup } from '@arco-design/web-vue';
  import { computed } from 'vue';

  export interface Props {
    multiple: boolean;
  }
  const props = defineProps<Props>();
  const { response: data } =
    useRequest<AlertChannelRecode[]>(queryChannelGroup);
  const amap = new Map<string, AlertChannelRecode[]>();
  const wechatOptions = computed<SelectOptionGroup[]>(() => {
    const initData = [] as SelectOptionGroup[];

    data.value.forEach((value) => {
      const d = amap.get(value.channelGroup);
      if (d === undefined) {
        const v = [] as AlertChannelRecode[];
        v.push(value);
        amap.set(value.channelGroup, v);
      } else {
        d.push(value);
        amap.set(value.channelGroup, d);
      }
    });

    amap.forEach((value, key) => {
      const s = {} as SelectOptionGroup;
      s.label = key;
      s.isGroup = true;
      s.options = [];
      value.forEach((item) => {
        s.options.push({
          label: item.channelName,
          value: item.id,
        });
      });
      initData.push(s);
    });

    if (data.value === undefined) {
      return initData;
    }

    return initData;
  });
</script>
