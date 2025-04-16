<template>
  <div class="container">
    <Breadcrumb
      :items="[{ label: 'menu.schedule' }, { label: 'menu.schedule.detail' }]"
    />
    <a-card class="general-card" title="值班详情">
      <a-divider style="margin-top: 0" />
      <a-row style="margin-bottom: 16px">
        <a-col :span="12">
          <a-month-picker
            v-model="yearMonth"
            style="width: 300px; margin-bottom: 24px; margin-right: 24px"
          />
        </a-col>
      </a-row>
      <a-table
        row-key="id"
        :loading="loading"
        :data="renderData"
        :bordered="{ wrapper: true, cell: true }"
        :selected-keys="selected"
        :pagination="false"
        @cell-click="cellClick"
      >
        <template #columns>
          <a-table-column title="姓名" data-index="nickname" align="center">
            <template #cell="{ record }">
              <div style="text-align: center">{{ record.nickname }}</div>
            </template>
          </a-table-column>
          <a-table-column
            v-for="(item, index) of columns"
            :key="index"
            :title="item.title as string"
            :cell-style="item.cellStyle"
          >
            <a-table-column
              v-for="(v, k) of item.children"
              :key="k"
              :title="v.title as string"
              :cell-style="v.cellStyle"
              align="center"
            >
              <template #cell="{ record }">
                <a-popover title="设置排班" trigger="click">
                  <div class="custom-table-cell">
                    <div
                      v-if="
                        record.plan.length > 0 && record.plan[index].isSchedule
                      "
                      class="schedule-text"
                      style="
                        --lightest: #fcdf7e;
                        --light: #fad355;
                        --medium: #ad7a03;
                        --darkest: #382201;
                      "
                      >值班</div
                    >
                  </div>
                  <template #content>
                    <a-button-group style="margin-top: 10px">
                      <a-button type="dashed" @click="clear(record, index)">
                        <icon-stop />
                        清除
                      </a-button>
                      <a-button type="dashed" @click="schedule(record, index)">
                        排班
                        <icon-check />
                      </a-button>
                    </a-button-group>
                  </template>
                </a-popover>
              </template>
            </a-table-column>
          </a-table-column>
        </template>
      </a-table>
    </a-card>
  </div>
</template>

<script lang="ts" setup>
  import { useRoute } from 'vue-router';
  import { computed, onMounted, ref, watch } from 'vue';
  import type { TableColumnData, TableData } from '@arco-design/web-vue';
  import useLoading from '@/hooks/loading';
  import {
    querySchedulePlanList,
    SchedulePlanParams,
    SchedulePlanRecord,
    setSchedulePlan,
    clearSchedulePlan,
  } from '@/api/schedule';

  const nowDate = new Date();
  const year = ref<number>(nowDate.getFullYear());
  const month = ref<number>(nowDate.getMonth() + 1);
  const rm = month.value < 10 ? `0${month.value}` : month.value;
  const yearMonth = ref(`${year.value}-${rm}`);
  const id = ref();
  const { loading, setLoading } = useLoading(true);
  const renderData = ref<SchedulePlanRecord[]>([]);
  const selected = ref<number[]>([]);

  function getWeekday(date: Date) {
    const weekdays = ['周日', '周一', '周二', '周三', '周四', '周五', '周六'];
    return weekdays[date.getDay()];
  }

  const route = useRoute();
  watch(
    route,
    () => {
      if (route.params.id) {
        id.value = Number(route.params?.id);
      }
    },
    {
      deep: true,
      immediate: true,
    }
  );

  const fetchData = async (
    params: SchedulePlanParams = { month: yearMonth.value }
  ) => {
    setLoading(true);
    try {
      const { data } = await querySchedulePlanList(id.value, params);
      renderData.value = data;
    } catch (err) {
      // you can report use errorHandler or other
    } finally {
      setLoading(false);
    }
  };

  const cellClick = (record: TableData, column: TableColumnData) => {
    if (column.dataIndex === 'nickname') {
      selected.value[0] = record.id;
    } else {
      selected.value = [];
    }
  };

  watch(yearMonth, () => {
    const updateDay = new Date(yearMonth.value);
    year.value = updateDay.getFullYear();
    month.value = updateDay.getMonth() + 1;
    renderData.value = [];
    fetchData();
  });

  const columns = computed(() => {
    const lastDay = new Date(year.value as number, month.value as number, 0);

    const datesWithWeekdays: TableColumnData[] = [];
    for (let day = 1; day <= lastDay.getDate(); day += 1) {
      const date = new Date(
        year.value as number,
        (month.value as number) - 1,
        day
      );
      const weekday = getWeekday(date);

      if (
        nowDate.getDate() === day &&
        nowDate.getFullYear() === year.value &&
        nowDate.getMonth() + 1 === month.value
      ) {
        datesWithWeekdays.push({
          title: weekday,
          cellStyle: {
            color: '#336df4 !important',
            fontSize: '13px',
          },
          children: [
            {
              title: String(day),
              dataIndex: String(day),
              align: 'center',
              cellStyle: {
                fontWeight: 'bold',
                color: '#336df4 !important',
                fontSize: '13px',
              },
            },
          ],
        });
      } else {
        datesWithWeekdays.push({
          title: weekday,
          children: [
            {
              title: String(day),
              dataIndex: String(day),
              align: 'center',
              cellStyle: {
                fontWeight: 'bold',
                color: 'black !important',
                fontSize: '13px',
              },
            },
          ],
        });
      }
    }
    return datesWithWeekdays;
  });

  const clear = async (record: TableData, columnIndex: number) => {
    setLoading(true);
    try {
      await clearSchedulePlan(record.id, { index: [columnIndex] });
      fetchData();
    } catch (err) {
      // you can report use errorHandler or other
    } finally {
      setLoading(false);
    }
  };

  const schedule = async (record: TableData, columnIndex: number) => {
    setLoading(true);
    try {
      await setSchedulePlan(record.id, { index: [columnIndex] });
      fetchData();
    } catch (err) {
      // you can report use errorHandler or other
    } finally {
      setLoading(false);
    }
  };

  onMounted(() => {
    fetchData();
  });
</script>

<style scoped lang="less">
  .container {
    padding: 0 20px 20px 20px;
  }

  .custom-table-cell:empty {
    min-height: 30px;
    min-width: 30px;
  }

  .custom-table-cell {
    min-width: 30px;
  }

  .schedule-text {
    align-items: center;
    background-color: var(--lightest);
    border-radius: 4px;
    color: var(--darkest);
    display: flex;
    font-size: 14px;
    height: 100%;
    justify-content: center;
    overflow: hidden;
    position: relative;
    text-align: center;
    -webkit-user-select: none;
    -ms-user-select: none;
    user-select: none;
    word-break: break-all;
    z-index: 6;
    font-weight: normal;
  }
  .schedule-text:after {
    border-left: 8px solid transparent;
    border-top: 8px solid #e22e28;
    content: '';
    height: 0;
    position: absolute;
    right: 0;
    top: 0;
    width: 0;
    z-index: 6;
  }
  .arco-table .arco-table-cell {
    padding: 1px 3px !important;
  }
  .arco-table-td {
    line-height: 2.5 !important;
  }
</style>
