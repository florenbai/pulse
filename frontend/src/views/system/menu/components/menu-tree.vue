<template>
  <a-tree-select
    :placeholder="props.placeholder"
    :data="menuList"
    :field-names="{
      key: 'key',
      title: 'name',
      children: 'children',
    }"
  >
  </a-tree-select>
</template>

<script lang="ts" setup>
  import useRequest from '@/hooks/request';
  import { computed } from 'vue';
  import { getMenuNameList } from '@/api/user';
  import { TreeNodeData } from '@arco-design/web-vue';

  export interface Props {
    placeholder?: string;
  }
  const props = withDefaults(defineProps<Props>(), {
    placeholder: '选择上级菜单',
  });
  const { response: data } = useRequest<TreeNodeData[]>(getMenuNameList);

  const menuList = computed<TreeNodeData[]>(() => {
    const menuData = [{ key: 0, name: '无' }] as any[];
    if (data.value === undefined) {
      return menuData;
    }
    menuData.push(...data.value);
    return menuData;
  });
</script>
