<template>
  <a-drawer
    v-model:visible="visible"
    width="30%"
    unmount-on-close
    :on-before-ok="handleOk"
  >
    <template #title> 编辑工作区 </template>
    <a-card :bordered="false" style="margin-top: 30px">
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
        </a-form>
      </div>
    </a-card>
  </a-drawer>
</template>

<script lang="ts" setup>
  import { ref } from 'vue';
  import { FormInstance } from '@arco-design/web-vue/es/form';
  import {
    BaseInfoRequest,
    queryWorkspaceBase,
    editWorkspaceForm,
  } from '@/api/workspace';
  import TeamSelect from '@/components/select/team-multiple-select.vue';
  import router from '@/router';
  import { Message } from '@arco-design/web-vue';

  const visible = ref(false);
  const emits = defineEmits(['refresh']);
  const wid = ref();
  const formRef = ref<FormInstance>();
  const formData = ref<BaseInfoRequest>({
    workspaceName: '',
    workspaceDesc: '',
    teams: [],
    strategy: 0,
  });

  const gotoAddTeam = () => {
    router.push('/system/team');
  };

  const handleOk = async (): Promise<boolean> => {
    const res = await formRef.value?.validate();
    if (!res) {
      try {
        await editWorkspaceForm(wid.value, formData.value);
        Message.success('编辑成功');
        emits('refresh');
        return true;
      } catch {
        //
      }
    }
    return false;
  };

  const getWorkspace = async (id: number) => {
    try {
      const res = await queryWorkspaceBase(id);
      formData.value = { ...res.data };
    } catch {
      //
    }
  };

  const show = (id: number) => {
    wid.value = id;
    getWorkspace(id);
    visible.value = true;
  };

  defineExpose({
    show,
  });
</script>
