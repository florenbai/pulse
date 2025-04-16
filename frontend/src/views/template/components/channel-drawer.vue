<template>
  <a-drawer
    :width="600"
    :visible="visible"
    unmount-on-close
    @ok="handleOk"
    @cancel="handleCancel"
  >
    <template #title> 告警模板内容设置 </template>

    <a-form
      ref="formRef"
      layout="vertical"
      :model="formData"
      style="margin-top: 20px"
    >
      <a-form-item
        style="margin-top: 10px"
        field="content"
        label="告警模板内容"
        validate-trigger="input"
        :rules="[
          {
            required: true,
            message: '请输入告警模板内容',
          },
        ]"
      >
        <a-textarea
          v-model="formData.content"
          :auto-size="{
            minRows: 5,
          }"
          placeholder="请输入告警模板内容"
          allow-clear
        />
      </a-form-item>
    </a-form>
    <a-collapse :default-active-key="['1']" :bordered="false">
      <a-collapse-item key="1" header="模板说明">
        <div>您可以使用Python模板语法 django 解析数据</div>
        <div
          >具体语法您点击
          <a-link
            href="https://www.osgeo.cn/django/ref/templates/language.html"
            target="_blank"
            >这里</a-link
          >进行了解</div
        >

        <div> 完整故障列表： </div>
        <div>
          <a-table :columns="columns" :data="data" :pagination="false" />
        </div>
      </a-collapse-item>
    </a-collapse>
  </a-drawer>
</template>

<script setup lang="ts">
  import {
    AlertTemplateChannel,
    AlertTemplateChannelRequest,
    AlertTemplateRecode,
    editTemplateChannelForm,
  } from '@/api/alert-template';
  import { FormInstance, Message } from '@arco-design/web-vue';
  import { reactive, ref } from 'vue';

  const emits = defineEmits(['refresh']);
  const visible = ref(false);
  const formRef = ref<FormInstance>();
  const editId = ref<number>();
  const formData = ref<AlertTemplateChannelRequest>({
    content: '',
  });

  const columns = [
    {
      title: '字段',
      dataIndex: 'name',
    },
    {
      title: '类型',
      dataIndex: 'type',
    },
    {
      title: '必须存在',
      dataIndex: 'exist',
    },
    {
      title: '释义',
      dataIndex: 'remark',
    },
  ];

  const data = reactive([
    {
      key: '1',
      name: 'Id',
      type: '数字',
      exist: '是',
      remark: '告警编号',
    },
    {
      key: '2',
      name: 'AlertTitle',
      type: '字符串',
      exist: '是',
      remark: '告警标题',
    },
    {
      key: '3',
      name: 'Description',
      type: '字符串',
      exist: '是',
      remark: '告警描述',
    },
    {
      key: '4',
      name: 'LevelName',
      type: '字符串',
      exist: '是',
      remark: '告警等级 支持A0-A5',
    },
    {
      key: '5',
      name: 'Tags',
      type: 'Map',
      exist: '是',
      remark: '告警标签',
    },
    {
      key: '6',
      name: 'TriggerTime',
      type: 'DateTime',
      exist: '是',
      remark: '最新告警时间',
    },
    {
      key: '7',
      name: 'SourceName ',
      type: '字符串',
      exist: '是',
      remark: '告警源',
    },
    {
      key: '8',
      name: 'Annotations',
      type: 'Map',
      exist: '是',
      remark: '告警附加信息',
    },
    {
      key: '9',
      name: 'NotifyCurNumber',
      type: '数字',
      exist: '是',
      remark: '告警数量',
    },
    {
      key: '10',
      name: 'Progress',
      type: '数字',
      exist: '是',
      remark: '处理进度 1 未认领 2 已认领 3 已关闭',
    },
    {
      key: '11',
      name: 'FirstAckTime',
      type: 'DateTime',
      exist: '否',
      remark: '确认时间',
    },
    {
      key: '12',
      name: 'RecoverTime',
      type: 'DateTime',
      exist: '否',
      remark: '恢复时间',
    },
    {
      key: '13',
      name: 'Nickname',
      type: '字符串',
      exist: '否',
      remark: '处理人',
    },
  ]);

  const handleOk = async () => {
    const res = await formRef.value?.validate();
    if (!res) {
      try {
        await editTemplateChannelForm(editId.value as number, formData.value);
        Message.success('渠道模板编辑成功');

        emits('refresh');
        visible.value = false;
      } catch (err) {
        //
      }
    }
  };
  const handleCancel = () => {
    visible.value = false;
  };

  const show = (item: AlertTemplateChannel | undefined) => {
    if (item !== undefined) {
      formData.value.content = item.content;
      editId.value = item.id;
    }
    visible.value = true;
  };

  defineExpose({
    show,
  });
</script>
