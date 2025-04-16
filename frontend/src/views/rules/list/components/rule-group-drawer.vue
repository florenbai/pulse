<template>
  <a-drawer
    :width="600"
    :visible="visible"
    :loading="loading"
    unmount-on-close
    @ok="handleOk"
    @cancel="handleCancel"
  >
    <template #title> 配置规则分组 </template>
    <a-form
      ref="formRef"
      layout="vertical"
      :model="form"
      style="margin-top: 20px"
    >
      <a-form-item
        field="groupName"
        label="规则分组名称"
        :rules="[
          {
            required: true,
            message: '请输入规则分组名称',
          },
        ]"
      >
        <a-input v-model="form.groupName" placeholder="请输入规则分组名称">
        </a-input>
      </a-form-item>
      <a-form-item
        field="ruleSource"
        label="数据源"
        :rules="[
          {
            required: true,
            message: '请输入数据源',
          },
        ]"
      >
        <RuleSourceSelect v-model="form.ruleSource"></RuleSourceSelect>
      </a-form-item>
      <a-form-item
        field="file"
        label="文件地址"
        :rules="[
          {
            required: true,
            message: '请输入文件地址',
          },
        ]"
      >
        <a-input v-model="form.file" placeholder="请输入文件地址"> </a-input>
      </a-form-item>
    </a-form>
  </a-drawer>
</template>

<script lang="ts" setup>
  import { ref } from 'vue';
  import { FormInstance, Message } from '@arco-design/web-vue';
  import { submitAlertRuleGroup } from '@/api/alert-rule';
  import RuleSourceSelect from '@/components/select/rules/rule-source-select.vue';
  import useLoading from '@/hooks/loading';

  const { loading, setLoading } = useLoading(true);
  const emits = defineEmits(['refresh']);

  const visible = ref(false);
  const editId = ref(0);
  const formRef = ref<FormInstance>();
  const form = ref({
    groupName: '',
    ruleSource: undefined,
    file: '',
  });
  const handleCancel = () => {
    visible.value = false;
    form.value = {
      groupName: '',
      ruleSource: undefined,
      file: '',
    };
    editId.value = 0;
  };
  const handleOk = async () => {
    const res = await formRef.value?.validate();
    if (!res) {
      try {
        setLoading(true);
        await submitAlertRuleGroup(form.value);
        Message.success({
          content: '操作成功',
          duration: 5 * 1000,
        });
        emits('refresh');
      } catch (err) {
        return;
      } finally {
        setLoading(false);
      }
    } else {
      return;
    }
    handleCancel();
  };
  const show = () => {
    visible.value = true;
  };

  defineExpose({
    show,
  });
</script>

<style scoped>
  .grid {
    gap: 0.75rem !important;
  }
  .grid-item {
    position: relative !important;
    display: flex;
    height: 100% !important;
    cursor: pointer !important;
    flex-direction: column !important;
    align-items: center !important;
    overflow: hidden !important;
    border-radius: 0.25rem !important;
    border-width: 1px !important;
    border-style: solid !important;
    padding-top: 1rem !important;
    padding-bottom: 1rem !important;
    transition-property: color, background-color, border-color,
      text-decoration-color, fill, stroke, opacity, box-shadow, transform,
      filter, backdrop-filter, -webkit-text-decoration-color,
      -webkit-backdrop-filter !important;
    transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1) !important;
    transition-duration: 0.15s !important;
    border-color: #e5e7eb !important;
    background-color: rgb(var(--fc-fill-3) / var(--tw-bg-opacity)) !important;
  }

  .grid-item:hover {
    background-color: var(--color-secondary) !important;
  }

  .grid-item-checked {
    right: 0 !important;
    top: -5px !important;
    position: absolute !important;
    color: rgb(var(--primary-6));
  }

  .grid-item-checked-bg {
    rotate: 45deg !important;
    opacity: 1 !important;
    background-color: var(--color-primary-light-1);
    width: 26px !important;
    height: 26px !important;
    top: -11px !important;
    right: -11px !important;
    position: absolute !important;
  }
</style>
