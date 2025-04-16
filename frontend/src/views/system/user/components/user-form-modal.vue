<template>
  <a-drawer
    v-model:visible="visible"
    :title="title"
    :width="500"
    @cancel="handleCancel"
    @ok="handleBeforeOk"
  >
    <a-form ref="formRef" :model="form">
      <a-form-item
        field="username"
        label="用户名称"
        :rules="[
          {
            required: true,
            message: '请输入用户名称',
          },
        ]"
      >
        <a-input v-model="form.username" placeholder="请输入用户名称" />
      </a-form-item>
      <a-form-item
        field="nickname"
        label="显示名称"
        :rules="[
          {
            required: true,
            message: '请输入显示名称',
          },
        ]"
      >
        <a-input v-model="form.nickname" placeholder="请输入显示名称" />
      </a-form-item>
      <a-form-item
        field="userid"
        label="企业微信"
        :rules="[
          {
            required: true,
            message: '请输入企业微信编号',
          },
        ]"
      >
        <a-input
          v-model="form.userid"
          placeholder="请输入企业微信编号"
          allow-clear
        />
      </a-form-item>
      <a-form-item
        field="email"
        label="用户邮箱"
        :rules="[
          {
            required: true,
            message: '请输入用户邮箱',
          },
        ]"
      >
        <a-input
          v-model="form.email"
          placeholder="请输入用户邮箱"
          allow-clear
        />
      </a-form-item>
      <a-form-item
        field="phone"
        label="用户手机"
        :rules="[
          {
            required: true,
            message: '请输入用户手机',
          },
        ]"
      >
        <a-input
          v-model="form.phone"
          placeholder="请输入用户手机"
          allow-clear
        />
      </a-form-item>
    </a-form>
  </a-drawer>
</template>

<script lang="ts" setup>
  import { ref } from 'vue';
  import { FormInstance } from '@arco-design/web-vue/es/form';
  import {
    submitUserForm,
    UserFormRequest,
    submitEditUserForm,
    UserRecord,
  } from '@/api/user';
  import { Message } from '@arco-design/web-vue';

  const emits = defineEmits(['visibleChange', 'reloadChange']);
  const title = ref('新增用户');
  const visible = ref(false);
  const editId = ref(0);
  const formRef = ref<FormInstance>();
  const form = ref<UserFormRequest>({
    username: '',
    nickname: '',
    password: '',
    email: '',
    phone: '',
    userid: '',
    status: true,
  });

  const show = (record: UserRecord | undefined) => {
    if (record !== undefined) {
      title.value = '编辑用户';
      editId.value = record.id;
      form.value = { ...record };
    }

    visible.value = true;
  };

  const handleBeforeOk = async () => {
    const res = await formRef.value?.validate();
    if (!res) {
      try {
        if (editId.value > 0) {
          await submitEditUserForm(editId.value, form.value);
        } else {
          await submitUserForm(form.value);
        }

        Message.success({
          content: '操作成功',
          duration: 5 * 1000,
        });
        form.value = {} as any;
        emits('reloadChange');
      } catch (err) {
        //
      }
    } else {
      visible.value = true;
    }
  };
  const handleCancel = () => {
    visible.value = false;
  };

  defineExpose({
    show,
  });
</script>
