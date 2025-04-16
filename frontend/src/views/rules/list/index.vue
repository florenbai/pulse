<template>
  <div class="container">
    <Breadcrumb :items="[{ label: 'menu.alert.rule' }]" />
    <div class="layout">
      <div class="layout-left-side">
        <RuleGroup @select-rule-group="selectRuleGroup" />
      </div>
      <div class="layout-content">
        <a-space :size="16" direction="vertical" fill>
          <GroupRuleTable :select-key="selectKey" :select-group="selectVal" />
        </a-space>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
  import { ref } from 'vue';
  import RuleGroup from './components/rule-group.vue';
  import GroupRuleTable from './components/group-table.vue';

  const selectKey = ref<number>();
  const selectVal = ref<string>();
  const selectRuleGroup = (k: number, v: string) => {
    selectKey.value = k;
    selectVal.value = v;
    window.scrollTo({
      top: 0,
      left: 0,
      behavior: 'smooth', // 'smooth' 表示平滑滚动，'auto' 表示瞬间滚动
    });
  };
</script>

<script lang="ts">
  export default {
    name: 'SearchTable',
  };
</script>

<style scoped lang="less">
  .container {
    padding: 0 20px 20px 20px;
  }

  .layout {
    display: flex;

    &-left-side {
      flex-basis: 300px;
    }

    &-content {
      flex: 1;
      padding: 0 16px;
    }

    &-right-side {
      flex-basis: 280px;
    }
  }
</style>

<style lang="less" scoped>
  // responsive
  @media (max-width: @screen-lg) {
    .layout {
      flex-wrap: wrap;
      &-left-side {
        flex: 1;
        flex-basis: 100%;
        margin-bottom: 16px;
      }

      &-content {
        flex: none;
        flex-basis: 100%;
        padding: 0;
        order: -1;
        margin-bottom: 16px;
      }

      &-right-side {
        flex-basis: 100%;
      }
    }
  }
</style>
