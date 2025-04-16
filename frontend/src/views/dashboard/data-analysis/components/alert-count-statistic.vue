<template>
  <a-spin :loading="loading" style="width: 100%">
    <a-card class="general-card" :header-style="{ paddingBottom: '16px' }">
      <template #title> 告警数量统计 </template>
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
    AlertCountRecord,
    TopAlertParams,
    queryAlertCountStatistic,
  } from '@/api/alert-statistic';

  const { loading, setLoading } = useLoading(true);
  const dateList = ref<string[]>([]);
  const valueList = ref<number[]>([]);
  const listParams = ref<TopAlertParams>({
    timeRange: '7d',
  });

  const { chartOption } = useChartOption(() => {
    return {
      visualMap: [
        {
          show: false,
          type: 'continuous',
          seriesIndex: 0,
          dimension: 0,
          min: 0,
          max: dateList.value.length - 1,
        },
      ],
      title: [
        {
          left: 'center',
          text: 'Gradient along the y axis',
        },
      ],
      tooltip: {
        trigger: 'axis',
      },
      xAxis: [
        {
          data: dateList.value,
        },
      ],
      yAxis: [{}],
      grid: [
        {
          bottom: '10%',
        },
        {
          top: '1%',
        },
      ],
      series: [
        {
          type: 'line',
          showSymbol: false,
          data: valueList.value,
        },
      ],
    };
  });

  const fetchData = async () => {
    setLoading(true);
    dateList.value = [];
    valueList.value = [];
    try {
      const { data } = await queryAlertCountStatistic(listParams.value);

      data.forEach((el: AlertCountRecord) => {
        dateList.value.push(el.date);
        valueList.value.push(el.num);
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
