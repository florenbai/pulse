<template>
  <div>
    <a-grid
      :cols="{ xs: 1, sm: 2, md: 3, lg: 3, xl: 4, xxl: 4 }"
      :col-gap="12"
      :row-gap="16"
      class="grid"
    >
      <strategyTemplate :id="props.id"></strategyTemplate>
    </a-grid>
  </div>
</template>

<script lang="ts" setup>
  import strategyTemplate from './strategy-template.vue';

  export interface Props {
    id: number | undefined;
  }
  const props = defineProps<Props>();
</script>
