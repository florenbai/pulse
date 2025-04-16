<template>
  <a-spin :loading="loading" style="width: 100%">
    <a-card class="general-card" :header-style="{ paddingBottom: '16px' }">
      <template #title> 告警TOP10 </template>
      <template #extra>
        <StatisticTimeSelect
          fault="最近7天"
          @change-value="changeTimeRange"
        ></StatisticTimeSelect>
      </template>
      <Chart style="width: 100%; height: 340px" :option="chartOption" />
    </a-card>
  </a-spin>
</template>

<script lang="ts" setup>
  import { ref } from 'vue';
  import useLoading from '@/hooks/loading';
  import StatisticTimeSelect from '@/components/select/statistic-time-range-select.vue';
  import useChartOption from '@/hooks/chart-option';
  import {
    TopAlertParams,
    TopAlertRecord,
    queryTopAlert,
  } from '@/api/alert-statistic';

  const { loading, setLoading } = useLoading(true);
  const chartTitleData = ref<string[]>([]);
  const chartValueData = ref<number[]>([]);
  const listParams = ref<TopAlertParams>({
    timeRange: '7d',
  });

  const { chartOption } = useChartOption(() => {
    return {
      tooltip: {
        trigger: 'axis',
        axisPointer: {
          type: 'shadow',
        },
      },
      grid: {
        left: '3%',
        right: '4%',
        bottom: '3%',
        containLabel: true,
      },
      xAxis: [
        {
          type: 'category',
          data: chartTitleData.value,
          axisTick: {
            alignWithLabel: true,
          },
        },
      ],
      yAxis: [
        {
          type: 'value',
        },
      ],
      series: [
        {
          name: '告警数量',
          type: 'bar',
          barWidth: '60%',
          data: chartValueData.value,
        },
      ],
    };
  });

  const fetchData = async () => {
    setLoading(true);
    chartTitleData.value = [];
    chartValueData.value = [];
    try {
      const { data } = await queryTopAlert(listParams.value);

      data.forEach((el: TopAlertRecord) => {
        chartTitleData.value.push(el.alertTitle);
        chartValueData.value.push(el.num);
      });
    } catch (err) {
      // you can report use errorHandler or other
    } finally {
      setLoading(false);
    }
  };
  fetchData();

  const changeTimeRange = (value: string) => {
    listParams.value.timeRange = value;
    fetchData();
  };
</script>
