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
  import { AlertChannelRecode } from '@/api/alert-template';
  import { SelectOptionGroup } from '@arco-design/web-vue';
  import { computed } from 'vue';

  export interface Props {
    multiple: boolean;
    channelGroup: AlertChannelRecode[];
  }
  const props = defineProps<Props>();

  const amap = new Map<string, AlertChannelRecode[]>();
  const wechatOptions = computed<SelectOptionGroup[]>(() => {
    const initData = [] as SelectOptionGroup[];
    if (props.channelGroup.length === 0) {
      return initData;
    }
    props.channelGroup.forEach((value) => {
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

    return initData;
  });
</script>
