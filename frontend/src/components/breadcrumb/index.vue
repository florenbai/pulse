<template>
  <a-breadcrumb class="container-breadcrumb">
    <a-breadcrumb-item>
      <icon-apps />
    </a-breadcrumb-item>
    <a-breadcrumb-item v-for="item in items" :key="item.label">
      <a-link :href="item.path">
        <span v-if="item.locale !== undefined && item.locale === false">
          {{ item.label }}
        </span>
        <span v-else>
          {{ $t(item.label) }}
        </span>
      </a-link>
    </a-breadcrumb-item>
  </a-breadcrumb>
</template>

<script lang="ts" setup>
  import { BreadcrumbTab } from '@/types/breadcrumb';
  import { PropType } from 'vue';

  defineProps({
    items: {
      type: Array as PropType<BreadcrumbTab[]>,
      default() {
        return [];
      },
    },
  });
</script>

<style scoped lang="less">
  .container-breadcrumb {
    margin: 16px 0;
    :deep(.arco-breadcrumb-item) {
      color: rgb(var(--gray-6));
      &:last-child {
        color: rgb(var(--gray-8));
      }
    }
  }
</style>
