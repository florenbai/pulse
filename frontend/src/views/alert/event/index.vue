<template>
  <div class="container">
    <Breadcrumb :items="[{ label: 'menu.alert.dashboard' }]" />
    <a-space direction="vertical" :size="12" fill>
      <a-space direction="vertical" :size="16" fill>
        <div class="space-unit">
          <StatisticComponents />
        </div>
        <div>
          <a-card class="general-card">
            <IncidentList></IncidentList>
          </a-card>
        </div>
      </a-space>
    </a-space>
  </div>
</template>

<script setup lang="ts">
  import StatisticComponents from './components/statistic.vue';
  import IncidentList from './components/incident-list.vue';
</script>

<style scoped lang="less">
  .container {
    padding: 0 20px 20px 20px;
  }
</style>
