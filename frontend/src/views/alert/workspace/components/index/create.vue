<template>
  <a-drawer v-model:visible="visible" width="25%" footer title="配置工作区">
    <a-card :bordered="false">
      <div class="main-content">
        <a-form ref="formRef" :model="formData" layout="vertical">
          <a-form-item
            field="workspaceName"
            label="工作区名称"
            validate-trigger="input"
            :rules="[
              {
                required: true,
                message: '请输入工作区名称',
              },
            ]"
          >
            <a-input
              v-model="formData.workspaceName"
              placeholder="请输入工作区名称，推荐以服务、功能、业务或团队的维度来创建工作区"
            />
          </a-form-item>
          <a-form-item
            field="workspaceDesc"
            label="工作区描述"
            validate-trigger="input"
          >
            <a-textarea
              v-model="formData.workspaceDesc"
              placeholder="请输入工作区描述"
            />
          </a-form-item>
          <a-form-item
            field="teams"
            label="所属团队"
            :rules="[
              {
                required: true,
                message: '请选择所属团队',
              },
            ]"
          >
            <TeamSelect
              v-model="formData.teams"
              :multiple="true"
              :prefix="false"
            />
            <template #extra>
              <div style="line-height: 20px"
                >如果不存在团队，需要您
                <a href="javascript:void(0)" @click="gotoAddTeam"
                  >创建团队</a
                ></div
              >
            </template>
          </a-form-item>
          <a-form-item
            field="strategy"
            label="全局告警策略"
            :rules="[
              {
                required: true,
                message: '请选择全局告警策略',
              },
            ]"
          >
            <StrategySelect
              v-model="formData.strategy"
              :multiple="false"
              :prefix="false"
            />
            <template #extra>
              <div style="line-height: 20px"
                >如果不存在全局告警策略，需要您
                <a href="javascript:void(0)" @click="gotoAddStrategy"
                  >创建全局告警策略</a
                ></div
              >
            </template>
          </a-form-item>
        </a-form>
      </div>
    </a-card>
    <template #footer>
      <a-button @click="closeModal">取消</a-button>
      <a-button type="primary" @click="handleOk">完成</a-button>
    </template>
  </a-drawer>
</template>

<script lang="ts" setup>
  import { ref } from 'vue';
  import TeamSelect from '@/components/select/team-multiple-select.vue';
  import StrategySelect from '@/components/select/workspace/system-strategy.vue';
  import { FormInstance } from '@arco-design/web-vue/es/form';
  import { BaseInfoRequest, submitWorkspaceForm } from '@/api/workspace';
  import router from '@/router';
  import { Message } from '@arco-design/web-vue';

  const emits = defineEmits(['refresh']);
  const visible = ref(false);
  const formRef = ref<FormInstance>();
  const formData = ref<BaseInfoRequest>({
    workspaceName: '',
    workspaceDesc: '',
    teams: [],
    strategy: undefined,
  });
  const show = () => {
    visible.value = true;
  };
  const handleOk = async () => {
    const res = await formRef.value?.validate();
    if (!res) {
      try {
        await submitWorkspaceForm(formData.value);
        Message.success({
          content: '新增完成',
          duration: 5 * 1000,
        });
        emits('refresh');
        visible.value = false;
      } catch (err) {
        //
      }
    }
  };

  const closeModal = () => {
    visible.value = false;
  };

  const gotoAddTeam = () => {
    router.push('/system/team');
  };

  const gotoAddStrategy = () => {
    router.push('/system/strategy');
  };

  defineExpose({
    show,
  });
</script>

<style scoped lang="less">
  .main-bottom {
    display: flex;
    justify-content: center;

    button {
      margin: 0 20px;
    }
  }
</style>
