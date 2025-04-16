<template>
  <div class="container">
    <Breadcrumb
      :items="[{ label: 'menu.dashboard' }, { label: 'menu.dataAnalysis' }]"
    />
    <a-space direction="vertical" :size="12" fill>
      <a-space direction="vertical" :size="16" fill>
        <div>
          <a-grid :cols="32" :col-gap="12" :row-gap="12">
            <a-grid-item
              :span="{ xs: 24, sm: 24, md: 24, lg: 24, xl: 16, xxl: 16 }"
            >
              <TopData></TopData>
            </a-grid-item>
            <a-grid-item
              :span="{ xs: 24, sm: 24, md: 24, lg: 24, xl: 16, xxl: 16 }"
            >
              <AlertLevelStatistic></AlertLevelStatistic>
            </a-grid-item>
          </a-grid>
        </div>
      </a-space>
      <a-space direction="vertical" :size="16" fill>
        <div>
          <AlertCountStatistic></AlertCountStatistic>
        </div>
      </a-space>
    </a-space>
  </div>
</template>

<script lang="ts" setup>
  import TopData from './components/top-alert.vue';
  import AlertLevelStatistic from './components/alert-level-statistic.vue';
  import AlertCountStatistic from './components/alert-count-statistic.vue';
</script>

<script lang="ts">
  export default {
    name: 'DataAnalysis',
  };
</script>

<style scoped lang="less">
  .container {
    padding: 0 20px 20px 20px;
    margin-bottom: 20px;
  }

  .space-unit {
    background-color: var(--color-bg-2);
    border-radius: 4px;
  }

  .title-fix {
    margin: 0 0 12px 0;
    font-size: 14;
  }
  :deep(.section-title) {
    margin: 0 0 12px 0;
    font-size: 14px;
  }
</style>
