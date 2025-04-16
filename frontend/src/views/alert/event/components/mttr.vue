<template>
  <a-spin :loading="loading" style="width: 100%">
    <a-card :bordered="false" :style="cardStyle">
      <div class="content-wrap">
        <div class="content">
          <a-statistic
            title="告警平均恢复（分钟）"
            :value="renderData.today"
            :value-from="0"
            animation
            show-group-separator
          />
          <div class="desc">
            <a-typography-text type="secondary" class="label">
              较昨日
            </a-typography-text>
            <a-typography-text
              v-if="renderData.direction === 'up'"
              type="danger"
            >
              {{ renderData.diff }}
              <icon-arrow-rise />
            </a-typography-text>
            <a-typography-text v-else type="success">
              {{ renderData.diff }}
              <icon-arrow-fall />
            </a-typography-text>
          </div>
        </div>
      </div>
    </a-card>
  </a-spin>
</template>

<script lang="ts" setup>
  import { ref, PropType, CSSProperties } from 'vue';
  import useLoading from '@/hooks/loading';
  import { AlertMttaRecord, queryAllAlertMttr } from '@/api/alert-statistic';

  const props = defineProps({
    cardStyle: {
      type: Object as PropType<CSSProperties>,
      default: () => {
        return {};
      },
    },
  });

  const { loading, setLoading } = useLoading(true);
  const renderData = ref<AlertMttaRecord>({
    today: 0,
    yesterday: 0,
    diff: 0,
    direction: 'up',
  });
  const fetchData = async () => {
    try {
      const { data } = await queryAllAlertMttr();
      renderData.value = data;
    } catch (err) {
      // you can report use errorHandler or other
    } finally {
      setLoading(false);
    }
  };
  fetchData();
</script>

<style scoped lang="less">
  :deep(.arco-card) {
    border-radius: 4px;
  }
  :deep(.arco-card-body) {
    width: 100%;
    height: 134px;
    padding: 0;
  }
  .content-wrap {
    width: 100%;
    padding: 16px;
    white-space: nowrap;
  }
  :deep(.content) {
    float: left;
    width: 108px;
    height: 102px;
  }
  :deep(.arco-statistic) {
    .arco-statistic-title {
      font-size: 16px;
      font-weight: bold;
      white-space: nowrap;
    }
    .arco-statistic-content {
      margin-top: 10px;
    }
  }

  .chart {
    float: right;
    width: calc(100% - 108px);
    height: 90px;
    vertical-align: bottom;
  }

  .label {
    padding-right: 8px;
    font-size: 12px;
  }
  .arco-card-body {
    padding: 0 !important;
  }
</style>
