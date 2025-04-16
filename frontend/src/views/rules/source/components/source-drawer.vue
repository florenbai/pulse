<template>
  <a-drawer
    :width="600"
    :visible="visible"
    unmount-on-close
    :ok-loading="loading"
    @ok="handleOk"
    @cancel="handleCancel"
  >
    <template #title> 配置数据源 </template>

    <div style="margin-top: 40px">
      <a-grid
        :cols="{ xs: 1, sm: 2, md: 3, lg: 3, xl: 4, xxl: 4 }"
        :col-gap="12"
        :row-gap="16"
        class="grid"
      >
        <a-grid-item
          v-for="item in alertSource"
          :key="item.id"
          class="grid-item"
          @click="checkSource(item)"
        >
          <div class="logo"> <icon-font :type="item.icon" :size="32" /></div>
          <div class="text-center text-l1 text-main">
            {{ item.sourceName }}
          </div>
          <div v-if="checkedSource === item.id">
            <div class="grid-item-checked-bg"></div>
            <div class="grid-item-checked"><icon-check /> </div
          ></div>
        </a-grid-item>
      </a-grid>
    </div>
    <a-form
      ref="formRef"
      layout="vertical"
      :model="form"
      style="margin-top: 20px"
    >
      <a-form-item
        field="sourceName"
        label="数据源名称"
        :rules="[
          {
            required: true,
            message: '请输入数据源名称',
          },
        ]"
      >
        <a-input v-model="form.sourceName" placeholder="请输入数据源名称">
        </a-input>
      </a-form-item>
      <a-form-item field="mark" label="备注">
        <a-input v-model="form.mark" placeholder="备注"> </a-input>
      </a-form-item>
      <a-form-item
        field="sourceName"
        label="服务地址"
        :rules="[
          {
            required: true,
            message: '请输入服务地址',
          },
        ]"
      >
        <a-input v-model="form.address" placeholder="请输入服务地址"> </a-input>
      </a-form-item>
      <a-form-item field="source_sign" label="数据源标识">
        <a-input v-model="form.sign" placeholder="备注" readonly></a-input>
      </a-form-item>
    </a-form>
  </a-drawer>
</template>

<script lang="ts" setup>
  import { ref, onMounted } from 'vue';
  import { FormInstance, Message } from '@arco-design/web-vue';
  import { AlertSourceRecode, queryAllAlertSource } from '@/api/alert-source';
  import {
    AlertRuleSourceRecode,
    submitAlertRuleSource,
    updateAlertRuleSource,
  } from '@/api/alert-rule';
  import { v4 as uuidv4 } from 'uuid';
  import useLoading from '@/hooks/loading';

  const emits = defineEmits(['refresh']);
  const { loading, setLoading } = useLoading(false);

  const visible = ref(false);
  const editId = ref(0);
  const checkedSource = ref(0);
  const formRef = ref<FormInstance>();
  const alertSource = ref<AlertSourceRecode[]>([]);
  const form = ref({
    sourceName: '',
    sourceType: '',
    mark: '',
    sign: '',
    address: '',
  });
  const handleCancel = () => {
    visible.value = false;
    form.value = {
      sourceName: '',
      sourceType: '',
      mark: '',
      sign: '',
      address: '',
    };
    editId.value = 0;
  };
  const handleOk = async () => {
    const res = await formRef.value?.validate();
    if (!res) {
      try {
        setLoading(true);
        if (editId.value > 0) {
          await updateAlertRuleSource(editId.value, form.value);
        } else {
          await submitAlertRuleSource(form.value);
        }

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
    visible.value = false;
  };
  const show = (item?: AlertRuleSourceRecode) => {
    if (item !== undefined) {
      form.value = { ...item };
      editId.value = item.id;
      alertSource.value.forEach((v: any) => {
        if (v.sourceName === item.sourceType) {
          checkedSource.value = v.id;
        }
      });
    } else {
      form.value.sign = uuidv4();
    }
    visible.value = true;
  };

  onMounted(async () => {
    try {
      const res = await queryAllAlertSource();
      alertSource.value = res.data;
    } catch (err) {
      //
    }
  });

  const checkSource = (record: AlertSourceRecode) => {
    checkedSource.value = record.id;
    form.value.sourceType = record.sourceName;
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
