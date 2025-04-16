<template>
  <a-grid-item
    :key="props.record.id"
    class="grid-item"
    @click="openEditAlertSource()"
  >
    <div style="display: inline-flex">
      <div class="logo" style="margin-right: 8px">
        <icon-font :type="props.record.icon" :size="36"
      /></div>
      <div class="text-center text-l1 text-main">
        {{ sourceName }}
        <div class="description">{{ props.record.description }}</div>
        <div v-if="props.record.eventCount === 0" class="text-tip">
          <span style="color: darkorange; display: inline-flex">
            尚未收到该告警源的告警
          </span>
        </div>
        <div v-else class="text-tip">
          <span style="color: #2b912b; display: inline-flex">
            累计收到 {{ props.record.eventCount }} 条告警
          </span>
        </div>
      </div>
      <div
        class="delete-icon"
        @click.stop="(event) => deleteSource(props.record.id, event)"
      >
        <icon-delete :size="20" />
      </div>
    </div>
  </a-grid-item>
  <a-drawer
    :width="600"
    :visible="visible"
    unmount-on-close
    ok-text="确认"
    @ok="handleOk"
    @cancel="handleCancel"
  >
    <template #title> 告警源接入说明 </template>
    <div>
      <a-card title="">
        <a-space direction="vertical" fill>
          <a-row class="grid-demo">
            <a-col :span="24">
              <div>
                <span class="title">告警源名称</span>
              </div>
            </a-col>
          </a-row>
          <a-row class="grid-demo">
            <a-col :span="23">
              <div>
                <a-input
                  v-model:model-value="sourceName"
                  class="source-input"
                />
              </div>
            </a-col>
          </a-row>
          <a-row class="grid-demo">
            <a-col :span="24">
              <div>
                <span class="title">告警源描述</span>
              </div>
            </a-col>
          </a-row>
          <a-row class="grid-demo">
            <a-col :span="23">
              <div>
                <a-textarea
                  v-model:model-value="description"
                  class="source-input"
                />
              </div>
            </a-col>
          </a-row>
          <a-row class="grid-demo">
            <a-col :span="24">
              <div>
                <span class="title">接入地址</span>
              </div>
            </a-col>
          </a-row>
          <a-row class="grid-demo">
            <a-col :span="23">
              <div>
                {{ axios.defaults.baseURL }}/api/v1/alert-event/push/{{
                  props.record.id
                }}
              </div>
            </a-col>
            <a-col :span="1">
              <a-link
                @click="
                  copyToClipboard(
                    axios.defaults.baseURL +
                      '/api/v1/alert-event/push/' +
                      props.record.id
                  )
                "
              >
                <template #icon>
                  <icon-copy size="20" />
                </template>
              </a-link>
            </a-col>
          </a-row>
          <a-row class="grid-demo">
            <a-col :span="24">
              <div>
                <span class="title">请求token</span>
              </div>
            </a-col>
          </a-row>
          <a-row class="grid-demo">
            <a-col :span="23">
              {{ token }}
            </a-col>
            <a-col :span="1">
              <a-popconfirm content="是否要刷新token" @ok="refreshToken">
                <a-link>
                  <template #icon>
                    <icon-refresh size="20" />
                  </template>
                </a-link>
              </a-popconfirm>
            </a-col>
          </a-row>
          <a-row v-if="props.record.icon !== 'icon-custom'" class="grid-demo">
            <a-col :span="24">
              <div>
                <span class="title">告警等级字段</span>
              </div>
            </a-col>
          </a-row>
          <a-row v-if="props.record.icon !== 'icon-custom'" class="grid-demo">
            <a-col :span="24">
              <a-input
                v-model:model-value="alertLevelField"
                class="source-input"
                @change="levelFieldChange"
              />
            </a-col>
          </a-row>
          <a-row v-if="props.record.icon !== 'icon-custom'" class="grid-demo">
            <a-col :span="24">
              <div>
                <span class="title">告警等级映射</span>
              </div>
            </a-col>
          </a-row>
          <a-row v-if="props.record.icon !== 'icon-custom'" class="grid-demo">
            <a-col :span="24">
              <a-table
                :columns="columns"
                :data="levelMapping"
                :pagination="false"
              >
                <template #levelName="{ rowIndex }">
                  <a-tag :color="levelMapping[rowIndex].color">{{
                    levelMapping[rowIndex].levelName
                  }}</a-tag>
                </template>
                <template #alertSourceLevel="{ rowIndex }">
                  <a-input
                    v-model="levelMapping[rowIndex].alertLevel"
                    @change="updateLevelMapping(levelMapping[rowIndex])"
                  />
                </template>
              </a-table>
            </a-col>
          </a-row>
        </a-space>
      </a-card>
      <div style="margin-top: 10px"></div>
      <a-card title="接入说明">
        <a-descriptions
          v-if="props.record.icon === 'icon-custom'"
          title="请求参数"
          bordered
        />
        <a-table
          v-if="props.record.icon === 'icon-custom'"
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
    </div>
  </a-drawer>
</template>

<script lang="ts" setup>
  import { h, onMounted, reactive, ref } from 'vue';
  import axios from 'axios';
  import {
    WorkspaceAlertSourceRecode,
    deleteAlertSource,
  } from '@/api/alert-source';
  import {
    editWorkspaceAlertSourceName,
    refreshWorkspaceAlertSourceToken,
    updateWorkspaceAlertSourceLevelMapping,
  } from '@/api/workspace';
  import {
    queryAllAlertMappingLevel,
    AlertLevelMappingRecode,
  } from '@/api/alert-level';
  import { Message, Notification } from '@arco-design/web-vue';

  export interface Props {
    record: WorkspaceAlertSourceRecode;
  }
  const props = defineProps<Props>();
  const emits = defineEmits(['handleClickWarning', 'refresh']);
  const visible = ref(false);
  const levelMapping = ref<AlertLevelMappingRecode[]>([]);
  const sourceName = ref(props.record.sourceName);
  const description = ref(props.record.description);
  const alertLevelField = ref(props.record.levelField);
  const token = ref(props.record.token);
  const columns = [
    {
      title: '系统告警等级',
      dataIndex: 'levelName',
      slotName: 'levelName',
    },
    {
      title: '告警源告警等级',
      dataIndex: 'alertSourceLevel',
      slotName: 'alertSourceLevel',
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
      name: 'recovered',
      type: '布尔型',
      essential: '是',
      desc: '告警恢复状态 true 恢复 false 告警 ',
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
  ]);
  const openEditAlertSource = () => {
    visible.value = true;
  };

  const sourceNameInputChange = async () => {
    try {
      await editWorkspaceAlertSourceName(props.record.id, {
        sourceName: sourceName.value,
        description: description.value,
      });

      Message.success({
        content: '成功更新了告警源信息',
        duration: 3 * 1000,
      });
      visible.value = false;
    } catch (err) {
      //
    }
  };

  const handleOk = () => {
    sourceNameInputChange();
  };

  const handleCancel = () => {
    visible.value = false;
  };
  const refreshToken = async () => {
    try {
      const { data } = await refreshWorkspaceAlertSourceToken(props.record.id);
      Message.success({
        content: '更新token成功',
        duration: 3 * 1000,
      });
      token.value = data;
    } catch (err) {
      //
    }
  };

  const updateLevelMapping = async (item: AlertLevelMappingRecode) => {
    try {
      const { data } = await updateWorkspaceAlertSourceLevelMapping({
        id: item.id,
        alertSourceLevel: item.alertLevel,
      });
      Message.success({
        content: '告警等级更新成功',
        duration: 3 * 1000,
      });
      token.value = data;
    } catch (err) {
      //
    }
  };

  const levelFieldChange = async () => {
    try {
      await editWorkspaceAlertSourceName(props.record.id, {
        levelField: alertLevelField.value,
      });
      Message.success({
        content: '成功更新了告警源信息',
        duration: 3 * 1000,
      });
    } catch (err) {
      //
    }
  };

  async function copyToClipboard(text: string) {
    try {
      await navigator.clipboard.writeText(text);
      Message.success({
        content: '已为您复制到粘贴板',
        duration: 5 * 1000,
      });
    } catch (err) {
      //
    }
  }

  const deleteSource = async (id: number, event: any) => {
    const x = `${event.clientX - 340}px`;
    const y = `${event.clientY - 100}px`;
    const nid = `${Date.now()}`;
    Notification.info({
      id: nid,
      title: '告警源删除提醒',
      position: 'topLeft',
      duration: 0,
      content: '请确认是否要删除该告警源',
      style: { left: x, top: y },
      footer: () =>
        h('Space', [
          h(
            'Button',
            {
              type: 'secondary',
              size: 'small',
              style: {
                color: '#var(--color-text-1)',
                backgroundColor: 'transparent',
                border: '1px solid var(--color-border-3)',
                height: '28px',
                padding: '0 15px',
                borderRadius: 'var(--border-radius-small)',
                fontSize: '14px',
                marginRight: '8px',
              },
              onClick: () => {
                Notification.remove(nid);
              },
            },
            '取消'
          ),
          h(
            'Button',
            {
              type: 'primary',
              size: 'small',
              style: {
                color: '#fff',
                backgroundColor: 'rgb(var(--primary-6))',
                border: '1px solid transparent',
                height: '28px',
                padding: '0 15px',
                borderRadius: 'var(--border-radius-small)',
                fontSize: '14px',
              },
              onClick: async () => {
                try {
                  await deleteAlertSource(id);
                  Message.success('告警源删除成功');
                  Notification.remove(nid);
                  emits('refresh');
                } catch (err) {
                  //
                }
              },
            },
            '确认'
          ),
        ]),
    });
  };

  onMounted(async () => {
    try {
      const { data } = await queryAllAlertMappingLevel(props.record.id);
      levelMapping.value = data;
    } catch (err) {
      //
    }
  });
</script>

<style scoped>
  .grid {
    gap: 0.75rem;
  }
  .grid-item {
    display: flex;
    height: 100%;
    cursor: pointer;
    position: relative;
    flex-direction: column;
    overflow: hidden;
    border-radius: 0.25rem;
    border-width: 1px;
    border-style: solid;
    padding-top: 1rem !important;
    padding-left: 1.5rem;
    padding-bottom: 1rem !important;
    transition-property: color, background-color, border-color,
      text-decoration-color, fill, stroke, opacity, box-shadow, transform,
      filter, backdrop-filter, -webkit-text-decoration-color,
      -webkit-backdrop-filter !important;
    transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1) !important;
    transition-duration: 0.15s !important;
    border-color: #e5e7eb;
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

  .text-tip {
    font-size: 12px;
    margin-top: 5px;
  }

  .delete-icon {
    position: absolute;
    right: 10px;
    top: 50%;
    transform: translate(-50%, -50%);
    display: none;
  }

  .grid-item:hover .delete-icon {
    display: flex;
  }

  .title {
    font-size: 14px;
    font-weight: 600;
  }

  .source-input {
    width: 500px;
  }

  .grid-demo {
    margin-bottom: 10px;
  }

  .description {
    font-size: 13px;
    margin-top: 5px;
    color: darkgray;
  }
</style>
