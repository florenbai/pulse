<template>
  <a-drawer
    :width="600"
    :visible="visible"
    unmount-on-close
    ok-text="确认"
    @ok="handleOk"
    @cancel="handleCancel"
  >
    <template #title> 配置告警集成 </template>
    <a-form
      ref="formRef"
      layout="vertical"
      :model="formData"
      style="margin-top: 20px"
    >
      <a-form-item
        field="name"
        label="集成名称"
        validate-trigger="input"
        :rules="[
          {
            required: true,
            message: '请输入集成名称',
          },
        ]"
      >
        <a-input v-model="formData.name" placeholder="请输入集成名称" />
      </a-form-item>
      <a-form-item field="description" label="集成描述">
        <a-textarea
          v-model="formData.description"
          placeholder="请输入集成描述"
          allow-clear
        />
      </a-form-item>
      <a-form-item
        field="token"
        label="集成Token"
        validate-trigger="input"
        :rules="[
          {
            required: true,
            message: '请输入集成Token',
          },
        ]"
      >
        <div style="display: flex; width: 100%">
          <a-input v-model="formData.token" placeholder="请输入集成Token">
            ></a-input
          >
          <a-button type="primary">
            <template #icon>
              <icon-refresh @click="refreshToken()" />
            </template>
          </a-button>
        </div>
      </a-form-item>
      <a-form-item
        v-if="formData.icon !== 'icon-custom'"
        field="levelField"
        label="告警等级字段"
        validate-trigger="input"
        :rules="[
          {
            required: true,
            message: '请输入告警等级字段',
          },
        ]"
      >
        <a-input
          v-model="formData.levelField"
          placeholder="请输入告警等级字段"
        />
      </a-form-item>
      <a-row v-if="formData.icon !== 'icon-custom'" class="grid-demo">
        <a-col :span="24">
          <div>
            <span class="title">告警等级映射</span>
          </div>
        </a-col>
      </a-row>
      <a-row v-if="formData.icon !== 'icon-custom'" class="grid-demo">
        <a-col :span="24">
          <a-table :columns="columns" :data="levelMapping" :pagination="false">
            <template #levelName="{ rowIndex }">
              <a-tag :color="levelMapping[rowIndex].color">{{
                levelMapping[rowIndex].levelName
              }}</a-tag>
            </template>
            <template #alertLevel="{ rowIndex }">
              <a-input v-model="levelMapping[rowIndex].alertLevel" />
            </template>
          </a-table>
        </a-col>
      </a-row>
    </a-form>

    <div style="margin-top: 10px"></div>
    <a-card title="接入说明">
      <a-descriptions
        v-if="formData.icon === 'icon-custom'"
        title="请求参数"
        bordered
      />
      <a-table
        v-if="formData.icon === 'icon-custom'"
        :columns="customColumns"
        :data="customData"
        :pagination="false"
      />
      <a-row class="grid-demo">
        <a-col :span="24">
          <p>通过webhook的方式接入系统时，请配置头信息。 </p>
          <p>
            头信息为基本身份验证
            <span style="font-weight: bold"
              >Authorization: Basic `credentials`</span
            >
          </p>
          <p> 然后我们构造credentials如下： </p>
          <p> 将token使用base64进行编码 </p>
        </a-col>
      </a-row>
      <a-row class="grid-demo">
        <a-col :span="24">
          <span> </span>
        </a-col>
      </a-row>
    </a-card>
  </a-drawer>
</template>

<script lang="ts" setup>
  import { FormInstance, Message } from '@arco-design/web-vue';
  import { reactive, ref } from 'vue';
  import { v4 as uuidv4 } from 'uuid';
  import {
    SystemIntegrationRecode,
    updateSystemIntegration,
  } from '@/api/integration';
  import {
    updateLevelMapping,
    AlertLevelMappingRecode,
    queryAllAlertMappingLevel,
  } from '@/api/alert-level';

  const visible = ref(false);
  const formRef = ref<FormInstance>();
  const formData = ref<SystemIntegrationRecode>({
    id: 0,
    name: '',
    description: '',
    token: '',
    levelField: '',
    icon: '',
    eventCount: 0,
    disabled: false,
  });
  const levelMapping = ref<AlertLevelMappingRecode[]>([]);
  const emits = defineEmits(['refresh']);
  const columns = [
    {
      title: '系统告警等级',
      dataIndex: 'levelName',
      slotName: 'levelName',
    },
    {
      title: '告警源告警等级',
      dataIndex: 'alertLevel',
      slotName: 'alertLevel',
    },
  ];

  const customColumns = [
    {
      title: '参数名称',
      dataIndex: 'name',
      ellipsis: true,
      tooltip: true,
      width: 180,
    },
    {
      title: '参数类型',
      dataIndex: 'type',
      width: 120,
    },
    {
      title: '是否必填',
      dataIndex: 'essential',
      ellipsis: true,
      width: 100,
    },
    {
      title: '说明',
      dataIndex: 'desc',
      ellipsis: true,
      tooltip: { position: 'left' },
      width: 400,
    },
  ];

  const customData = reactive([
    {
      key: '1',
      name: 'alertTitle',
      type: '字符串',
      essential: '是',
      desc: '告警标题',
    },
    {
      key: '2',
      name: 'description',
      type: '字符串',
      essential: '是',
      desc: '告警描述',
    },
    {
      key: '3',
      name: 'level',
      type: '字符串',
      essential: '是',
      desc: '告警等级',
    },
    {
      key: '4',
      name: 'status',
      type: '字符串',
      essential: '是',
      desc: '告警状态 resolved 恢复 firing 告警 ',
    },
    {
      key: '5',
      name: 'alertTime',
      type: 'DateTime',
      essential: '否',
      desc: '告警时间，如果不传，则使用当前时间',
    },
    {
      key: '6',
      name: 'labels',
      type: 'Map对象',
      essential: '否',
      desc: '告警标签 key代表标签，value代表值',
    },
    {
      key: '7',
      name: 'annotations',
      type: 'Map对象',
      essential: '否',
      desc: '告警注解 key代表标签，value代表值',
    },
    {
      key: '8',
      name: 'fingerprint',
      type: '字符串',
      essential: '是',
      desc: '告警指纹',
    },
  ]);

  const handleCancel = () => {
    visible.value = false;
  };

  const handleOk = async () => {
    const res = await formRef.value?.validate();
    if (!res) {
      try {
        await updateLevelMapping(levelMapping.value);
        await updateSystemIntegration(formData.value.id, formData.value);
        Message.success('告警集成编辑成功');
        handleCancel();
        emits('refresh', formData.value.id);
      } catch (err) {
        //
      }
    }
  };

  const refreshToken = () => {
    formData.value.token = uuidv4();
  };

  const show = async (item: SystemIntegrationRecode) => {
    visible.value = true;
    formData.value = { ...item };
    try {
      const { data } = await queryAllAlertMappingLevel(formData.value.id);
      levelMapping.value = data;
    } catch (err) {
      //
    }
  };

  defineExpose({
    show,
  });
</script>
