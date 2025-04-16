<template>
  <a-drawer
    :width="600"
    :visible="visible"
    unmount-on-close
    :ok-loading="loading"
    @ok="handleOk"
    @cancel="handleCancel"
  >
    <template #title> 配置告警规则 </template>
    <a-form
      ref="formRef"
      layout="vertical"
      :model="form"
      style="margin-top: 20px"
    >
      <div style="margin-bottom: 20px"
        >告警规则分组 : <a-tag color="arcoblue">{{ groupName }}</a-tag></div
      >
      <a-form-item
        field="name"
        label="告警规则标题"
        :rules="[
          {
            required: true,
            message: '请输入告警规则标题',
          },
        ]"
      >
        <a-input v-model="form.name" placeholder="请输入告警规则标题">
        </a-input>
      </a-form-item>
      <a-form-item
        field="query"
        label="查询语句"
        :rules="[
          {
            required: true,
            message: '请输入查询语句',
          },
        ]"
      >
        <a-input
          v-model="form.query"
          :data="alertSourceLabel"
          placeholder="请输入查询语句"
        >
          <template #option="item">
            <icon-common /> {{ item.data.value }}
          </template>
        </a-input>
      </a-form-item>
      <a-form-item
        field="duration"
        label="评估等待时间"
        :rules="[
          {
            required: true,
            message: '请输入评估等待时间',
          },
        ]"
      >
        <a-input-number
          v-model="form.duration"
          placeholder="请输入评估等待时间"
        >
          <template #suffix> 秒 </template>
        </a-input-number>
      </a-form-item>
      <a-form-item field="labels" label="标签">
        <a-card :style="{ 'margin-top': '10px', 'flex': 1 }">
          <a-input-group
            v-for="(value, item) in labels"
            :key="item"
            style="margin-bottom: 8px; width: 100%"
          >
            <a-input v-model="value.tag" placeholder="键" />
            <a-input v-model="value.value" allow-clear placeholder="值" />
            <a-button
              style="margin-left: 5px; background-color: #fff"
              @click="deleteLabel(item)"
            >
              <template #icon> <icon-delete size="15" /> </template
            ></a-button>
          </a-input-group>

          <a-button style="display: flex" type="text" @click="createLabel()"
            >+ 添加标签</a-button
          >
        </a-card>
      </a-form-item>
      <a-form-item field="annotations" label="注解">
        <a-card :style="{ 'margin-top': '10px', 'flex': 1 }">
          <a-input-group
            v-for="(value, item) in annotations"
            :key="item"
            style="margin-bottom: 8px; width: 100%"
          >
            <a-input v-model="value.tag" placeholder="键" />
            <a-input v-model="value.value" allow-clear placeholder="值" />
            <a-button
              style="margin-left: 5px; background-color: #fff"
              @click="deleteAnnotation(item)"
            >
              <template #icon> <icon-delete size="15" /> </template
            ></a-button>
          </a-input-group>

          <a-button
            style="display: flex"
            type="text"
            @click="createAnnotation()"
            >+ 添加注解</a-button
          >
        </a-card>
      </a-form-item>
    </a-form>
  </a-drawer>
</template>

<script lang="ts" setup>
  import { ref } from 'vue';
  import { FormInstance, Message } from '@arco-design/web-vue';
  import {
    AlertRuleLabels,
    AlertRuleRecode,
    querySourceLabels,
    submitAlertRule,
    updateAlertRule,
  } from '@/api/alert-rule';
  import useLoading from '@/hooks/loading';

  const emits = defineEmits(['refresh']);
  const { loading, setLoading } = useLoading(false);
  const labels = ref<AlertRuleLabels[]>([{ tag: '', value: '' }]);
  const annotations = ref<AlertRuleLabels[]>([{ tag: '', value: '' }]);
  const visible = ref(false);
  const editId = ref(0);
  const groupName = ref();
  const formRef = ref<FormInstance>();
  const alertSourceLabel = ref<string[]>([]);
  const form = ref({
    name: '',
    gid: 0,
    query: '',
    duration: 0,
    labels: [],
    annotation: [],
  });
  const handleCancel = () => {
    visible.value = false;
    form.value = {
      name: '',
      gid: 0,
      query: '',
      duration: 0,
      labels: [],
      annotation: [],
    };
    editId.value = 0;
    setLoading(false);
  };

  const handleOk = async () => {
    form.value.labels = labels as any;
    form.value.annotation = annotations as any;

    const res = await formRef.value?.validate();
    if (!res) {
      setLoading(true);
      try {
        if (editId.value > 0) {
          await updateAlertRule(editId.value, form.value);
        } else {
          await submitAlertRule(form.value);
        }

        Message.success({
          content: '操作成功',
          duration: 5 * 1000,
        });
        emits('refresh');
      } catch (err) {
        setLoading(false);
        return;
      }
    } else {
      setLoading(false);
      return;
    }
    handleCancel();
  };

  const getSourceLabels = async () => {
    try {
      const res = await querySourceLabels(form.value.gid);
      alertSourceLabel.value = res.data;
    } catch (err) {
      //
    }
  };

  const show = (id: number, name: string, record?: AlertRuleRecode) => {
    form.value.gid = id;
    groupName.value = name;
    visible.value = true;
    labels.value = [{ tag: '', value: '' }];
    annotations.value = [{ tag: '', value: '' }];
    getSourceLabels();
    if (record !== undefined) {
      editId.value = record.id;
      form.value.name = record.name;
      form.value.query = record.query;
      form.value.duration = record.duration;
      const annotationKey = Object.keys(record.annotations);
      const annotationValue = Object.values(record.annotations);
      labels.value = [];
      annotations.value = [];
      for (let i = 0; i < annotationKey.length; i += 1) {
        annotations.value.push({
          tag: annotationKey[i],
          value: annotationValue[i],
        });
      }
      const labelsKey = Object.keys(record.labels);
      const labelsValue = Object.values(record.labels);

      for (let i = 0; i < labelsKey.length; i += 1) {
        labels.value.push({
          tag: labelsKey[i],
          value: labelsValue[i],
        });
      }
    }
  };

  const createLabel = () => {
    labels.value.push({ tag: '', value: '' });
  };

  const deleteLabel = (key: number) => {
    labels.value.splice(key, 1);
  };

  const createAnnotation = () => {
    annotations.value.push({ tag: '', value: '' });
  };

  const deleteAnnotation = (key: number) => {
    annotations.value.splice(key);
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
