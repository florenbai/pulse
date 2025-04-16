<template>
  <a-spin :loading="loading" style="width: 100%">
    <a-card
      :bordered="false"
      style="
        background: linear-gradient(
          rgb(247, 247, 255) 0%,
          rgb(236, 236, 255) 100%
        );
      "
    >
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
  import { ref } from 'vue';
  import useLoading from '@/hooks/loading';
  import { queryAllAlertLevel } from '@/api/alert-statistic';
  import useChartOption from '@/hooks/chart-option';

  const { loading, setLoading } = useLoading(true);

  const barChartOptionsFactory = () => {
    const data = ref<any>([]);
    const { chartOption } = useChartOption(() => {
      return {
        xAxis: {
          type: 'category',
          data: ['A0', 'A1', 'A2', 'A3', 'A4', 'A5'],
          axisLabel: {
            inside: true,
            color: '#000',
          },
          axisTick: {
            show: false,
          },
          axisLine: {
            show: false,
          },
          z: 10,
        },
        yAxis: {
          show: false,
        },
        tooltip: {
          show: true,
          trigger: 'axis',
        },
        series: {
          name: 'total',
          data: data.value,
          type: 'bar',
          itemStyle: {
            borderRadius: 2,
          },
        },
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
      const { data } = await queryAllAlertLevel();

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
    white-space: nowrap;
    padding: 16px;
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
    top: 60px;
  }

  .label {
    padding-right: 8px;
    font-size: 12px;
  }

  .arco-card-body {
    padding: 0 !important;
  }
</style>
