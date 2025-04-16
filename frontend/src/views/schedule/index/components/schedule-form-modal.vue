<template>
  <a-drawer
    :visible="visible"
    :title="title"
    :width="500"
    @cancel="handleCancel"
    @before-ok="handleBeforeOk"
  >
    <a-form ref="formRef" :model="form" layout="vertical">
      <a-form-item
        field="scheduleName"
        label="值班名称"
        :rules="[
          {
            required: true,
            message: '请输入值班名称',
          },
        ]"
      >
        <a-input v-model="form.scheduleName" placeholder="请输入值班名称" />
      </a-form-item>
      <a-form-item
        field="workspaceDesc"
        label="值班描述"
        validate-trigger="input"
      >
        <a-textarea v-model="form.scheduleDesc" placeholder="请输入值班描述" />
      </a-form-item>
      <a-form-item
        field="teamId"
        label="值班团队"
        :rules="[
          {
            required: true,
            message: '请选择值班团队',
          },
        ]"
      >
        <TeamSelect
          v-model="form.teamId"
          :multiple="false"
          @change="teamChange"
        />
        <template #extra>
          <div style="line-height: 20px"
            >如果不存在团队，需要您
            <a href="javascript:void(0)" @click="gotoAddTeam">创建团队</a></div
          >
        </template>
      </a-form-item>
      <a-form-item
        field="scheduleTimeRange"
        label="值班时间范围"
        :rules="[
          {
            required: true,
            message: '请选择值班时间范围',
          },
        ]"
      >
        <a-time-picker
          v-model="form.scheduleTimeRange"
          type="time-range"
          style="width: 252px"
          :default-value="['00:00:00', '23:59:59']"
        />
      </a-form-item>
      <a-form-item
        field="autoScheduling"
        label="自动排班"
        validate-trigger="input"
      >
        <a-switch
          v-model="switchStatus"
          :before-change="handleChangeIntercept"
        />
      </a-form-item>
      <a-form-item v-if="switchStatus" label="值班人设置">
        <a-table
          v-model:selectedKeys="selectedKeys"
          row-key="userId"
          :columns="columns"
          :data="teamData"
          :row-selection="rowSelection"
          :draggable="{ type: 'handle', width: 40 }"
          style="display: block; width: 300px"
          :pagination="false"
          @change="handleChange"
        />
      </a-form-item>
      <a-form-item v-if="switchStatus" field="schedulePeriod" label="轮换周期">
        <PeriodSelect
          v-model="form.schedulePeriod"
          @update="changePeriod"
        ></PeriodSelect>
      </a-form-item>
      <!--
      <a-form-item
        v-if="switchStatus"
        field="scheduleUserCount"
        label="值班人数"
      >
        <a-input-number
          v-model="form.scheduleUserCount"
          placeholder="请输入值班名称"
        />
      </a-form-item>
      -->
    </a-form>
  </a-drawer>
</template>

<script lang="ts" setup>
  import { reactive, ref } from 'vue';
  import { FormInstance } from '@arco-design/web-vue/es/form';
  import { ScheduleForm, submitScheduleForm } from '@/api/schedule';
  import TeamSelect from '@/components/select/team-multiple-select.vue';
  import PeriodSelect from '@/components/select/workspace/period-select.vue';
  import router from '@/router';
  import { Message } from '@arco-design/web-vue';
  import { queryAllTeamUser } from '@/api/team';

  const emits = defineEmits(['reloadChange']);
  const title = ref('新增值班');
  const visible = ref(false);
  const teamData = ref([]);
  const formRef = ref<FormInstance>();
  const selectedKeys = ref([]);
  const switchStatus = ref(false);
  const form = ref<ScheduleForm>({
    scheduleName: '',
    scheduleDesc: '',
    teamId: undefined,
    schedulePeriod: ['1', 'day'],
    scheduleTimeRange: [],
    users: [],
  });

  const rowSelection = reactive({
    showCheckedAll: true,
    onlyCurrent: false,
  });

  const columns = reactive([
    {
      title: '值班人',
      dataIndex: 'nickname',
      align: 'center' as const,
    },
  ]);

  const handleChange = (_data: any) => {
    teamData.value = _data;
  };

  const show = () => {
    visible.value = true;
  };
  const gotoAddTeam = () => {
    router.push('/system/team');
  };

  const handleCancel = () => {
    visible.value = false;
  };

  const teamChange = () => {
    switchStatus.value = false;
    teamData.value = [];
  };

  const handleChangeIntercept = async () => {
    if (!switchStatus.value) {
      if (form.value.teamId === undefined) {
        Message.error({
          content: '请选择值班团队',
          duration: 5 * 1000,
        });
        return false;
      }
      try {
        const teams = [];
        teams.push(form.value.teamId);
        const { data } = await queryAllTeamUser(teams);
        teamData.value = data;
      } catch (err) {
        //
      }
    }
    return true;
  };

  const handleBeforeOk = async () => {
    const res = await formRef.value?.validate();
    form.value.users = selectedKeys.value;
    if (!res) {
      try {
        await submitScheduleForm(form.value);

        Message.success({
          content: '操作成功',
          duration: 5 * 1000,
        });
        handleCancel();
        emits('reloadChange');
      } catch (err) {
        //
      }

      return true;
    }
    return false;
  };

  const changePeriod = (sdata: string[]) => {
    form.value.schedulePeriod = sdata;
  };

  defineExpose({
    show,
  });
</script>
