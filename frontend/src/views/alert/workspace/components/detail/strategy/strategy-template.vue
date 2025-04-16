<template>
  <draggable
    :list="alertStrategyList"
    handle=".handle"
    item-key="id"
    style="display: contents"
    :component-data="{ tag: 'ul', name: 'flip-list', type: 'transition' }"
    @change="dragMoved"
  >
    <template #item="{ index, element }">
      <a-grid-item class="grid-item" style="border-style: dashed">
        <div class="card-wrap" @click="editStrategy(element)">
          <a-space align="start">
            <a-card-meta>
              <template #title>
                <a-typography-text style="margin-right: 10px">
                  {{ element.strategyName }}
                </a-typography-text>

                <a-tag
                  v-if="element.status === 'enabled'"
                  size="small"
                  color="green"
                >
                  <template #icon>
                    <icon-check-circle-fill />
                  </template>
                  <span>启用</span>
                </a-tag>
                <a-tag
                  v-else-if="element.status === 'disabled'"
                  size="small"
                  color="red"
                >
                  <template #icon>
                    <icon-stop />
                  </template>
                  <span>禁用</span>
                </a-tag>
                <span class="right-link">
                  <a-link>策略顺序{{ index + 1 }}</a-link>
                </span>
                <span class="right-link-hover">
                  <a-button type="text" class="handle" @click.stop="">
                    <template #icon>
                      <icon-drag-dot-vertical size="15" />
                    </template>
                  </a-button>
                  <a-button
                    type="text"
                    @click.stop="scheduleSwitch(element, index)"
                  >
                    <template #icon>
                      <icon-stop
                        v-if="element.status === 'enabled'"
                        size="15"
                      />
                      <icon-check-circle
                        v-else-if="element.status === 'disabled'"
                        size="15"
                      />
                    </template>
                  </a-button>

                  <a-button
                    v-if="element.status === 'disabled'"
                    type="text"
                    @click.stop="(event) => deleteStrategy(element, event)"
                  >
                    <template #icon>
                      <icon-delete size="15" />
                    </template>
                  </a-button>
                </span>
              </template>

              <template #description>
                <div class="description">
                  {{ element.strategyDesc }}
                </div>
                <div class="info">
                  {{ element.nickname }} {{ element.updatedAt }}
                </div>
              </template>
            </a-card-meta>
          </a-space>
        </div>
      </a-grid-item>
    </template>
  </draggable>
  <!--
  <a-grid-item
    v-for="(item, index) of alertStrategyList"
    :key="index"
    class="grid-item"
  >
    <div class="card-wrap" @click="editStrategy(item)">
      <a-space align="start">
        <a-card-meta>
          <template #title>
            <a-typography-text style="margin-right: 10px">
              {{ item.strategyName }}
            </a-typography-text>

            <a-tag v-if="item.status === 'enabled'" size="small" color="green">
              <template #icon>
                <icon-check-circle-fill />
              </template>
              <span>启用</span>
            </a-tag>
            <a-tag
              v-else-if="item.status === 'disabled'"
              size="small"
              color="red"
            >
              <template #icon>
                <icon-stop />
              </template>
              <span>禁用</span>
            </a-tag>
            <span class="right-link">
              <a-link>策略顺序{{ index + 1 }}</a-link>
            </span>
            <span class="right-link-hover">
              <drag
                :alert-strategy-list="alertStrategyList"
                @fetch-list-data="fetchListData"
              ></drag>
              <a-button type="text" @click.stop="scheduleSwitch(item, index)">
                <template #icon>
                  <icon-stop v-if="item.status === 'enabled'" size="15" />
                  <icon-check-circle
                    v-else-if="item.status === 'disabled'"
                    size="15"
                  />
                </template>
              </a-button>

              <a-button
                v-if="item.status === 'disabled'"
                type="text"
                @click.stop="(event) => deleteStrategy(item, event)"
              >
                <template #icon>
                  <icon-delete size="15" />
                </template>
              </a-button>
            </span>
          </template>

          <template #description>
            <div class="description">
              {{ item.strategyDesc }}
            </div>
            <div class="info"> {{ item.nickname }} {{ item.updatedAt }} </div>
          </template>
        </a-card-meta>
      </a-space>
    </div>
  </a-grid-item>
  -->
  <a-grid-item class="grid-item" @click="handleSystemStrategyClick">
    <div class="card-wrap">
      <a-space align="start">
        <a-card-meta>
          <template #title>
            <a-typography-text style="margin-right: 10px">
              {{ systemStrategy?.strategyName }}
            </a-typography-text>

            <a-tag size="small" color="green">
              <template #icon>
                <icon-command />
              </template>
              <span>全局告警通知策略</span>
            </a-tag>
          </template>

          <template #description>
            <div class="description">
              {{ systemStrategy?.strategyDesc }}
            </div>
          </template>
        </a-card-meta>
      </a-space>
    </div>
  </a-grid-item>
  <a-grid-item
    class="grid-item"
    style="border-style: dashed"
    @click="handleClick"
  >
    <div style="text-align: center; margin: auto">
      <icon-plus style="margin-right: 1rem" />
      <span>新增告警通知策略</span></div
    >
  </a-grid-item>
  <a-drawer
    :width="600"
    :visible="visible"
    unmount-on-close
    @ok="handleOk"
    @cancel="handleCancel"
  >
    <template #title> 配置自定义告警通知策略 </template>
    <div>
      <a-alert
        >您可以配置多个自定义告警通知策略，当故障触发后，将按照配置顺序依次匹配。默认情况下，当告警通知策略匹配后，将会停止其他策略匹配。如果您需要继续匹配其他告警策略，请开启连续匹配功能。</a-alert
      >
    </div>

    <a-form
      ref="formRef"
      layout="vertical"
      :model="formData"
      style="margin-top: 20px"
    >
      <a-collapse
        :active-key="activeKeyRef"
        :bordered="false"
        @change="changeCollapse"
      >
        <a-collapse-item key="1" header="基础信息" style="padding-left: 0px">
          <a-form-item
            style="margin-top: 10px"
            field="strategyName"
            label="策略名称"
            validate-trigger="input"
            :rules="[
              {
                required: true,
                message: '请输入策略名称',
              },
            ]"
          >
            <a-input
              v-model="formData.strategyName"
              placeholder="请输入策略名称"
            />
          </a-form-item>
          <a-form-item
            style="margin-top: 10px"
            field="strategyDesc"
            label="策略描述"
            validate-trigger="input"
          >
            <a-input
              v-model="formData.strategyDesc"
              placeholder="请输入策略描述"
            />
          </a-form-item>
          <a-form-item
            field="templateId"
            label="告警模板"
            :rules="[
              {
                required: true,
                message: '请选择告警模板',
              },
            ]"
          >
            <alert-template-select
              v-model="formData.templateId"
              @change="fetchChannelGroup(formData.templateId)"
            />
          </a-form-item>
          <a-form-item
            style="margin-top: 10px"
            field="delay"
            label="通知延迟"
            validate-trigger="input"
            :rules="[
              {
                required: true,
                message: '请输入通知延迟',
              },
            ]"
          >
            <a-input-number
              v-model="formData.delay"
              placeholder="请输入通知延迟时间"
            >
              <template #append> 秒 </template>
            </a-input-number>
            <template #extra>
              <div style="line-height: 20px"
                >0表示实时通知。如果在等待期内故障自动恢复或被认领，则通知不会被发送。</div
              >
            </template>
          </a-form-item>
          <a-form-item
            style="margin-top: 10px"
            field="delay"
            label="接续匹配"
            validate-trigger="input"
          >
            <a-switch v-model="formData.continuous" />
            <template #extra>
              <div style="line-height: 20px">
                开启接续匹配后，无论是否匹配到自定义通知策略，都会继续其他自定义策略匹配。如果关闭，当匹配到通知策略后，将不在继续匹配其他自定义通知策略。
              </div>
            </template>
          </a-form-item>
          <a-form-item
            style="margin-top: 10px"
            field="systemStrategy"
            label="全局策略"
            validate-trigger="input"
          >
            <a-switch v-model="formData.systemStrategy" />
            <template #extra>
              <div style="line-height: 20px">
                开启全局策略后，无论是否匹配到通知策略，都会继续使用全局策略进行通知。如果关闭，当匹配到通知策略后，将不在继续执行全局策略。
              </div>
            </template>
          </a-form-item>
        </a-collapse-item>
        <a-collapse-item
          key="2"
          header="通知时间段策略"
          style="padding-left: 0px"
          :disabled="formData.timeSlot.enable === false"
        >
          <template #extra>
            <a-switch
              v-model="formData.timeSlot.enable"
              :before-change="handleChangeIntercept"
            >
            </a-switch>
          </template>
          <a-form-item field="week" label="通知时间周期">
            <a-select
              v-model="formData.timeSlot.weeks"
              :style="{ width: '380px' }"
              placeholder="请选择通知时间周期"
              multiple
              allow-clear
            >
              <a-option :value="1">星期一</a-option>
              <a-option :value="2">星期二</a-option>
              <a-option :value="3">星期三</a-option>
              <a-option :value="4">星期四</a-option>
              <a-option :value="5">星期五</a-option>
              <a-option :value="6">星期六</a-option>
              <a-option :value="0">星期日</a-option>
            </a-select>
          </a-form-item>
          <a-form-item field="week" label="通知时间段">
            <a-time-picker
              v-model="formData.timeSlot.times"
              type="time-range"
            />
          </a-form-item>
        </a-collapse-item>
        <a-collapse-item
          key="3"
          header="标签匹配策略"
          style="padding-left: 0px"
          :disabled="tagMapping === false"
        >
          <template #extra>
            <a-switch
              v-model="tagMapping"
              :before-change="handleChangeTagMapping"
            >
            </a-switch>
          </template>
          <a-form-item field="week" label="">
            <div>
              <div
                v-for="(filter, filterIndex) of formData.filters"
                :key="filterIndex"
              >
                <a-card
                  :style="{ 'margin-top': filterIndex > 0 ? '10px' : '0' }"
                >
                  <div style="display: flex">
                    <div
                      v-if="filter.values !== null && filter.values.length > 1"
                      class="rule-condition-logic-group__logic"
                    >
                      <div
                        class="rule-condition-logic-group__logic-item"
                        :style="{
                          height: `${tagHeight[filterIndex]}px`,
                        }"
                      >
                        <div
                          class="rule-condition-logic-group__logic-item__start"
                          style="width: 12px"
                        ></div>
                        <span
                          class="rule-condition-logic-group__logic-item__text"
                          >且</span
                        >
                        <div
                          class="rule-condition-logic-group__logic-item__end"
                          style="width: 12px"
                        ></div>
                      </div>
                    </div>
                    <div>
                      <a-input-group
                        v-for="(item, index) of filter.values"
                        :key="index"
                        style="margin-bottom: 8px"
                      >
                        <a-input
                          v-model="item.tag"
                          :style="{ width: '120px' }"
                          placeholder="标签"
                        />
                        <a-select
                          v-model="item.operation"
                          style="
                            width: 100px;
                            margin-left: 5px;
                            margin-right: 5px;
                          "
                        >
                          <a-option value="IN">包含</a-option>
                          <a-option value="NOT IN">不包含</a-option>
                        </a-select>
                        <a-input-tag
                          v-model="item.value"
                          allow-clear
                          :style="{ width: '180px' }"
                          placeholder="值"
                        />
                        <a-button
                          style="margin-left: 5px; background-color: #fff"
                          @click="deleteFilterItem(filterIndex, index)"
                        >
                          <template #icon> <icon-delete size="15" /> </template
                        ></a-button>
                      </a-input-group>
                    </div>
                  </div>

                  <a-button type="text" @click="createCondition(filterIndex)"
                    >+ 添加条件</a-button
                  >
                </a-card>
              </div>
            </div>

            <a-button
              style="flex: auto; margin-top: 8px"
              type="dashed"
              @click="createFilter"
              >+ 或者</a-button
            >
          </a-form-item>
        </a-collapse-item>
        <a-collapse-item key="4" header="分派策略" style="padding-left: 0px">
          <assignment-strategy
            v-for="(item, index) of formData.strategySet"
            :key="index"
            :index="index"
            :channel-group="channelGroupData"
            :strategy-set="item"
            @handle-delete="handleDelete"
            @strategy-set-change="strategySetChange"
            @handle-add="handleAdd"
          ></assignment-strategy>
        </a-collapse-item>
      </a-collapse>
    </a-form>
  </a-drawer>
  <strategyDrawer ref="systemStrategyRef"></strategyDrawer>
</template>

<script lang="ts" setup>
  import { ref, onMounted, h } from 'vue';
  import { FormInstance, Message, Notification } from '@arco-design/web-vue';
  import alertTemplateSelect from '@/components/select/alert-template-select.vue';
  import draggable from 'vuedraggable';
  import {
    StrategyItem,
    AlertStrategyRequest,
    AlertStrategyRecord,
    submitStrategyForm,
    queryAllAlertStrategy,
    changeStrategyStatus,
    editStrategyForm,
    deleteStrategyRecord,
    editStrategySort,
    AlertStrategySort,
  } from '@/api/alert-strategy';
  import {
    querySystemStrategy,
    SystemStrategyRecord,
  } from '@/api/system-strategy';
  import strategyDrawer from '@/views/system/strategy/components/strategy-drawer.vue';
  import {
    AlertChannelRecode,
    queryTemplateEnableChannel,
  } from '@/api/alert-template';
  import assignmentStrategy from './assignment-strategy.vue';

  export interface Props {
    id: number | undefined;
  }
  const props = defineProps<Props>();
  const visible = ref(false);
  const editId = ref(0);
  const tagHeight = [40];
  const sortData = ref<AlertStrategySort[]>([]);
  const activeKeyRef = ref<(string | number)[]>(['1', '4']);
  const tagMapping = ref(false);
  const alertStrategyList = ref<AlertStrategyRecord[]>([]);
  const systemStrategy = ref<SystemStrategyRecord>();
  const systemStrategyRef = ref();
  const channelGroupData = ref<AlertChannelRecode[]>([]);
  const formRef = ref<FormInstance>();
  const formData = ref<AlertStrategyRequest>({
    workspaceId: props.id as number,
    strategyName: '',
    delay: 0,
    systemStrategy: false,
    continuous: false,
    strategyDesc: '',
    templateId: 1,
    timeSlot: {
      enable: false,
    },
    filters: [],
    strategySet: [
      {
        teams: [],
        schedules: [],
        members: [],
        alertChannels: [],
        alertLoop: {
          enable: false,
          interval: 0,
          count: 1,
        },
        progression: {
          enable: false,
          condition: 0,
          duration: 1,
        },
      },
    ],
  });
  const handleSystemStrategyClick = () => {
    systemStrategyRef.value.show(systemStrategy.value, true);
  };

  const dragMoved = async () => {
    for (let i = 0; i < alertStrategyList.value.length; i += 1) {
      sortData.value.push({
        id: alertStrategyList.value[i].id,
        weight: i + 1,
      });
    }
    try {
      await editStrategySort(sortData.value);
      Message.success('策略排序成功');
    } catch {
      //
    }
  };

  const handleClick = () => {
    formData.value = {
      workspaceId: props.id as number,
      strategyName: '',
      delay: 0,
      systemStrategy: false,
      continuous: false,
      strategyDesc: '',
      templateId: 1,
      timeSlot: {
        enable: false,
      },
      filters: [],
      strategySet: [
        {
          teams: [],
          schedules: [],
          members: [],
          alertChannels: [],
          alertLoop: {
            enable: false,
            interval: 0,
            count: 1,
          },
          progression: {
            enable: false,
            condition: 0,
            duration: 1,
          },
        },
      ],
    };
    editId.value = 0;
    visible.value = true;
  };

  const fetchChannelGroup = async (id: number) => {
    try {
      const res = await queryTemplateEnableChannel(id);
      channelGroupData.value = res.data;
    } catch (err) {
      //
    }
  };

  const fetchListData = async () => {
    try {
      const res = await queryAllAlertStrategy(props.id as number);
      alertStrategyList.value = res.data;
    } catch (err) {
      //
    }
  };

  const fetchSystemStrategy = async () => {
    try {
      const res = await querySystemStrategy(props.id as number);
      systemStrategy.value = res.data;
      fetchChannelGroup(systemStrategy.value?.templateId as number);
    } catch (err) {
      //
    }
  };

  const handleOk = async () => {
    const res = await formRef.value?.validate();
    if (!res) {
      try {
        if (editId.value !== 0) {
          await editStrategyForm(editId.value, formData.value);
          Message.success('策略编辑成功');
        } else {
          await submitStrategyForm(formData.value);
          Message.success('策略添加成功');
        }
        visible.value = false;
        fetchListData();
      } catch (err) {
        //
      }
    }
  };
  const handleCancel = () => {
    visible.value = false;
  };

  const changeCollapse = (activeKey: (string | number)[]) => {
    if (tagMapping.value && !activeKey.includes('3')) {
      activeKeyRef.value.push('3');
    } else if (formData.value.timeSlot.enable && !activeKey.includes('2')) {
      activeKeyRef.value.push('2');
    } else {
      activeKeyRef.value = activeKey;
    }
  };

  const handleChangeIntercept = () => {
    if (formData.value.timeSlot.enable) {
      const n = activeKeyRef.value.filter((item) => item !== '2');
      activeKeyRef.value = n;
      formData.value.timeSlot.enable = false;
      formData.value.timeSlot.times = undefined;
      formData.value.timeSlot.type = undefined;
      formData.value.timeSlot.weeks = undefined;
    } else {
      formData.value.timeSlot.enable = true;
    }
    return true;
  };

  const handleAdd = () => {
    formData.value.strategySet.push({
      members: [],
      teams: [],
      schedules: [],
      alertChannels: [],
      alertLoop: {
        enable: false,
        interval: 0,
        count: 1,
      },
      progression: {
        enable: false,
        condition: 0,
        duration: 1,
      },
    });
  };

  const handleDelete = (index: number) => {
    formData.value.strategySet.splice(index, 1);
  };

  const strategySetChange = (model: StrategyItem, index: number) => {
    formData.value.strategySet[index] = model;
  };

  const handleChangeTagMapping = () => {
    if (tagMapping.value) {
      const n = activeKeyRef.value.filter((item) => item !== '3');
      activeKeyRef.value = n;
      tagMapping.value = false;
      formData.value.filters = [];
    } else {
      if (formData.value.filters.values.length === 0) {
        formData.value.filters.push({
          values: [
            {
              tag: '',
              operation: 'IN',
              value: [],
            },
          ],
        });
      }
      tagMapping.value = true;
    }
  };

  const deleteFilterItem = (filterIndex: number, index: number) => {
    if (index === 0) {
      formData.value.filters[filterIndex].values.shift();
    } else {
      formData.value.filters[filterIndex].values.splice(index, 1);
    }
    if (formData.value.filters[filterIndex].values.length >= 2) {
      tagHeight[filterIndex] -= 40;
    }
    if (
      formData.value.filters[filterIndex].values.length === 0 &&
      formData.value.filters.length === 0
    ) {
      handleChangeTagMapping();
    } else if (
      formData.value.filters[filterIndex].values.length === 0 &&
      formData.value.filters.length !== 0
    ) {
      if (filterIndex === 0) {
        formData.value.filters.shift();
      } else {
        formData.value.filters.splice(filterIndex);
      }
    }
  };

  const changeScheduleStatus = async (id: number) => {
    try {
      await changeStrategyStatus(id);
    } catch (err) {
      //
    }
  };

  const scheduleSwitch = (item: AlertStrategyRecord, index: number) => {
    if (item.status === 'enabled') {
      item.status = 'disabled';
      alertStrategyList.value[index] = item;
    } else {
      item.status = 'enabled';
      alertStrategyList.value[index] = item;
    }
    changeScheduleStatus(item.id);
  };

  const createCondition = (filterIndex: number) => {
    formData.value.filters[filterIndex].values.push({
      tag: '',
      operation: 'IN',
      value: [],
    });
    if (formData.value.filters[filterIndex].values.length > 2) {
      tagHeight[filterIndex] =
        (formData.value.filters[filterIndex].values.length - 2) * 40 + 40;
    }
  };

  const createFilter = () => {
    tagHeight[formData.value.filters.length] = 40;
    formData.value.filters.push({
      values: [
        {
          tag: '',
          operation: 'IN',
          value: [],
        },
      ],
    });
  };

  const editStrategy = (item: AlertStrategyRecord) => {
    if (item.filters.length > 0) {
      if (!activeKeyRef.value.includes('3')) {
        activeKeyRef.value.push('3');
      }

      tagMapping.value = true;
    } else {
      if (activeKeyRef.value.includes('3')) {
        const n = activeKeyRef.value.filter((val) => val !== '3');
        activeKeyRef.value = n;
      }
      tagMapping.value = false;
    }
    if (item.timeSlot.enable) {
      if (!activeKeyRef.value.includes('2')) {
        activeKeyRef.value.push('2');
      }
    } else if (activeKeyRef.value.includes('2')) {
      const n = activeKeyRef.value.filter((val) => val !== '2');
      activeKeyRef.value = n;
    }
    formData.value = item;
    visible.value = true;
    editId.value = item.id;
  };

  const deleteStrategy = (item: AlertStrategyRecord, event: any) => {
    const x = `${event.clientX - 340}px`;
    const y = `${event.clientY - 100}px`;
    const id = `${Date.now()}`;
    Notification.info({
      id,
      title: '策略删除提醒',
      position: 'topLeft',
      duration: 0,
      content: '请确认是否要删除该策略',
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
                Notification.remove(id);
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
                  await deleteStrategyRecord(item.id);
                  Message.success('策略删除成功');
                  Notification.remove(id);
                  fetchListData();
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

  onMounted(() => {
    fetchListData();
    fetchSystemStrategy();
  });
</script>

<style scoped lang="less">
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
    border-color: #4f7fe1 !important;
  }

  .grid-item:hover .right-link-hover {
    display: inline-flex;
  }

  .grid-item:hover .right-link {
    display: none;
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
  .description {
    min-height: 40px;
    -webkit-line-clamp: 3;
    display: -webkit-inline-box;
    -webkit-box-orient: vertical;
    overflow: hidden;
    text-overflow: ellipsis;
    width: 85%;
  }

  .card-wrap {
    width: -webkit-fill-available;
    padding-left: 20px;
  }

  .right-link {
    position: absolute;
    right: 0;
    margin-right: 20px;
  }

  .right-link-hover {
    position: absolute;
    right: 0;
    margin-right: 20px;
    display: none;
  }

  .rule-condition-logic-group__logic {
    width: 27px;
    flex-grow: 0;
    flex-shrink: 0;
    overflow: hidden;
    position: relative;
    top: 16px;
  }

  .rule-condition-logic-group__logic-item {
    align-items: flex-start;
    box-sizing: initial;
    display: flex;
    flex-direction: column;
    overflow: hidden;
    pointer-events: none;
    position: absolute;
  }

  .rule-condition-logic-group__logic-item__start {
    border-top: 1px solid #d0d3d6;
    border-top-left-radius: 3px;
    border-left: 1px solid #d0d3d6;
    flex-grow: 1;
    margin-left: 7px;
    width: 12px;
  }

  .rule-condition-logic-group__logic-item__text {
    color: #646a73;
    line-height: 22px;
  }

  .rule-condition-logic-group__logic-item__end {
    width: 12px;
    border-bottom: 1px solid #d0d3d6;
    border-bottom-left-radius: 3px;
    border-left: 1px solid #d0d3d6;
    flex-grow: 1;
    margin-left: 7px;
  }
</style>

<style>
  .arco-collapse-item-content {
    background-color: #fff !important;
  }

  .arco-collapse-item-header {
    border: none !important;
  }
  .logic-group {
    width: 27px;
    right: 0px;
    flex-grow: 0;
    flex-shrink: 0;
    overflow: hidden;
    position: relative;
  }

  .logic-group-item {
    width: 12px;
    height: 120px;
    top: 16px;
    position: relative;
  }

  .logic-group-item__start {
    border-left: 1px solid #d0d3d6;
    flex-grow: 1;
    margin-left: 7px;
    width: 12px;
    border-top: 1px solid #d0d3d6;
    border-top-left-radius: 3px;
  }
</style>
