<template>
  <a-spin :loading="loading" style="width: 100%">
    <a-card :bordered="false" :style="cardStyle">
      <div class="content-wrap">
        <div class="content">
          <a-statistic title="告警等级分布" />
        </div>
        <div class="chart">
          <Chart v-if="!loading" :option="chartOption" />
        </div>
      </div>
    </a-card>
  </a-spin>
</template>

<script lang="ts" setup>
  import { ref, PropType, CSSProperties } from 'vue';
  import useLoading from '@/hooks/loading';
  import { queryAlertLevel } from '@/api/alert-statistic';
  import useChartOption from '@/hooks/chart-option';

  const { loading, setLoading } = useLoading(true);
  const props = defineProps({
    id: {
      type: Number,
    },
    cardStyle: {
      type: Object as PropType<CSSProperties>,
      default: () => {
        return {};
      },
    },
  });
  const barChartOptionsFactory = () => {
    const data = ref<any>([]);
    const { chartOption } = useChartOption(() => {
      return {
        xAxis: {
          type: 'category',
          data: ['A0', 'A1', 'A2', 'A3', 'A4', 'A5'],
          axisTick: {
            show: false,
          },
          axisLine: {
            show: false,
          },
        },

        yAxis: {
          show: false,
        },
        tooltip: {
          show: true,
          trigger: 'axis',
        },
        series: [
          {
            name: '告警数量',
            type: 'bar',
            stack: 'Total',
            label: {
              show: true,
              position: 'inside',
            },
            data: data.value,
          },
        ],
      };
    });
    return {
      data,
      chartOption,
    };
  };
  const { chartOption: barChartOption, data: BarData } =
    barChartOptionsFactory();
  const chartOption = ref({});
  const fetchData = async () => {
    try {
      const { data } = await queryAlertLevel(props.id as number);

      data.forEach((el) => {
        BarData.value.push({
          value: el.count,
          itemStyle: {
            color: el.color,
          },
        });
      });
      chartOption.value = barChartOption.value;
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
    width: 100%;
    height: 90px;
    vertical-align: bottom;
    position: absolute;
    top: 30px;
  }

  .label {
    padding-right: 8px;
    font-size: 12px;
  }
</style>
