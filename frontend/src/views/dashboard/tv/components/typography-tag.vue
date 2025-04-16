<template>
  <div v-if="showExpand && expandable">
    <a-tag
      v-for="(name, index) of props.tags"
      :key="index"
      color="arcoblue"
      style="margin: 2px; white-space: normal; height: auto"
    >
      <span>{{ index }} : {{ name }}</span>
    </a-tag>
    <a-link style="color: darkgray" @click="expandableHandle">折叠</a-link>
  </div>
  <div v-else-if="showExpand && !expandable">
    <a-tag
      v-for="(value, index) of data"
      :key="index"
      color="arcoblue"
      style="margin: 2px; white-space: normal; height: auto"
    >
      {{ value[0] }} : {{ value[1] }}
    </a-tag>
    <a-link style="color: darkgray" @click="expandableHandle">展开</a-link>
  </div>
  <div v-else>
    <a-tag
      v-for="(name, index) of props.tags"
      :key="index"
      color="arcoblue"
      style="margin: 2px; white-space: normal; height: auto"
    >
      <span>{{ index }} : {{ name }}</span>
    </a-tag>
  </div>
</template>

<script lang="ts" setup>
  import { ref } from 'vue';

  export interface Props {
    tags: any;
  }
  const props = defineProps<Props>();
  const expandable = ref(false);
  const showExpand = ref(false);

  const data = new Map<string, any>();

  const calculate = (key: string, value: any, c: number): number => {
    c = c + key.length + value.length;
    if (c > 500 && data.size > 0) {
      return c;
    }
    data.set(key, value);
    return c;
  };
  const fetchData = () => {
    let i = 0;
    Object.entries(props.tags).forEach(([key, value]) => {
      i = calculate(key, value, i);
      if (i > 500) {
        showExpand.value = true;
      }
    });
  };
  fetchData();

  const expandableHandle = () => {
    if (expandable.value) {
      expandable.value = false;
    } else {
      expandable.value = true;
    }
  };
</script>
