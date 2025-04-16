<template>
  <a-modal
    v-model:visible="teamModalRef.visible.value"
    :title="title"
    @cancel="handleCancel"
    @before-ok="handleBeforeOk"
  >
    <a-form ref="formRef" :model="form">
      <a-form-item
        field="teamName"
        label="团队名称"
        :rules="[
          {
            required: true,
            message: '请输入团队名称',
          },
        ]"
      >
        <a-input v-model="form.teamName" placeholder="请输入团队名称" />
      </a-form-item>
      <a-form-item
        field="teamDesc"
        label="团队描述"
        :rules="[
          {
            required: true,
            message: '请输入团队描述',
          },
        ]"
      >
        <a-textarea v-model="form.teamDesc" placeholder="请输入团队描述" />
      </a-form-item>
      <a-form-item
        field="teamMembers"
        label="团队成员"
        :rules="[
          {
            required: true,
            message: '请选择团队成员',
          },
        ]"
      >
        <UserSelect
          v-model="form.teamMembers"
          multiple
          :prefix="false"
          placeholder="请选择团队成员"
        />
      </a-form-item>
      <a-form-item field="status" label="状态">
        <a-switch
          v-model="form.status"
          :checked-value="form.status"
          :unchecked-value="form.status"
        >
          <template #checked> 启用 </template>
          <template #unchecked> 禁用 </template>
        </a-switch>
      </a-form-item>
    </a-form>
  </a-modal>
</template>

<script lang="ts" setup>
  import { ref, toRefs, watch } from 'vue';
  import { FormInstance } from '@arco-design/web-vue/es/form';
  import { Message } from '@arco-design/web-vue';
  import {
    submitTeamForm,
    TeamRequest,
    queryTeamInfo,
    submitEditTeamForm,
  } from '@/api/team';
  import UserSelect from '@/components/select/user-multiple-select.vue';

  export interface Props {
    id: number;
    visible: boolean;
  }
  const props = defineProps<Props>();
  const emits = defineEmits(['visibleChange', 'reloadChange']);
  const title = ref('新增团队');
  const formRef = ref<FormInstance>();
  const teamModalRef = toRefs(props);
  const form = ref<TeamRequest>({
    teamName: '',
    teamDesc: '',
    teamMembers: [],
    status: true,
  });

  watch(teamModalRef.id, async () => {
    if (teamModalRef.id.value !== undefined || teamModalRef.id.value !== 0) {
      try {
        const res = await queryTeamInfo(teamModalRef.id.value);
        form.value = res.data;
        title.value = '编辑团队';
      } catch (err) {
        //
      }
    }
  });

  const handleBeforeOk = async (done: any) => {
    const res = await formRef.value?.validate();
    if (!res) {
      try {
        if (teamModalRef.id.value > 0) {
          await submitEditTeamForm(teamModalRef.id.value, form.value);
        } else {
          await submitTeamForm(form.value);
        }

        Message.success({
          content: '操作成功',
          duration: 5 * 1000,
        });
        form.value = {} as TeamRequest;
        done();
        emits('reloadChange');
      } catch (err) {
        done(false);
      }
    } else {
      done(false);
    }
  };
  const handleCancel = () => {
    emits('visibleChange');
  };
</script>
