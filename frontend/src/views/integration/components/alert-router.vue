<template>
  <a-drawer
    :width="600"
    :visible="visible"
    unmount-on-close
    @ok="handleOk"
    @cancel="handleCancel"
  >
    <template #title> 配置告警路由 </template>

    <a-form
      ref="formRef"
      layout="vertical"
      :model="formData"
      style="margin-top: 20px"
    >
      <div>
        <a-card title="限定条件" :bordered="false">
          <template #extra>
            <a-switch
              v-model="tagMapping"
              :before-change="handleChangeTagMapping"
            >
            </a-switch>
          </template>

          <div v-if="tagMapping && formData.filters.length > 0">
            <div
              v-for="(filter, filterIndex) of formData.filters"
              :key="filterIndex"
            >
              <a-card :style="{ 'margin-top': filterIndex > 0 ? '10px' : '0' }">
                <div style="display: flex">
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
                  >+ 添加并列条件</a-button
                >
              </a-card>
            </div>
            <a-button
              style="flex: auto; margin-top: 8px"
              type="dashed"
              @click="createFilter"
              >+ 或者</a-button
            >
          </div>
        </a-card>
      </div>
      <a-form-item
        field="workspaces"
        label="转发工作区"
        validate-trigger="input"
        :rules="[
          {
            required: true,
            message: '请选则转发工作区',
          },
        ]"
      >
        <WorkspaceSelect
          v-model="formData.workspaces"
          :multiple="true"
        ></WorkspaceSelect>
      </a-form-item>
      <a-form-item field="next" label="匹配命中后">
        <a-radio-group v-model="formData.next" type="button">
          <a-radio :value="1">继续匹配</a-radio>
          <a-radio :value="0">停止匹配</a-radio>
        </a-radio-group>
      </a-form-item>
    </a-form>
  </a-drawer>
</template>

<script lang="ts" setup>
  import {
    IntegrationRouter,
    submitIntegrationRouterForm,
    updateIntegrationRouterForm,
  } from '@/api/integration';
  import WorkspaceSelect from '@/components/select/source/workspace-select.vue';
  import { FormInstance, Message } from '@arco-design/web-vue';
  import { ref } from 'vue';

  const emits = defineEmits(['refresh']);
  const sourceId = ref<number>();
  const visible = ref(false);
  const tagMapping = ref(false);
  const formRef = ref<FormInstance>();
  const formData = ref<IntegrationRouter>({
    id: 0,
    workspaces: [],
    workspaceNames: [],
    filters: [],
    next: 1,
    sort: 0,
  });

  const handleChangeTagMapping = () => {
    if (tagMapping.value) {
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

  const createCondition = (filterIndex: number) => {
    formData.value.filters[filterIndex].values.push({
      tag: '',
      operation: 'IN',
      value: [],
    });
  };

  const deleteFilterItem = (filterIndex: number, index: number) => {
    if (index === 0) {
      formData.value.filters[filterIndex].values.shift();
    } else {
      formData.value.filters[filterIndex].values.splice(index, 1);
    }

    if (
      formData.value.filters[filterIndex].values.length === 0 &&
      formData.value.filters.length !== 0
    ) {
      if (filterIndex === 0) {
        formData.value.filters.shift();
      } else {
        formData.value.filters.splice(filterIndex);
      }
    }
    if (formData.value.filters.length === 0) {
      handleChangeTagMapping();
    }
  };

  const handleCancel = () => {
    formData.value = {
      id: 0,
      workspaces: [],
      workspaceNames: [],
      filters: [],
      next: 1,
      sort: 0,
    };
    visible.value = false;
    tagMapping.value = false;
    emits('refresh', sourceId.value);
  };

  const createFilter = () => {
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

  const show = (id: number, item?: IntegrationRouter) => {
    sourceId.value = id;
    visible.value = true;
    if (item !== undefined) {
      formData.value = { ...item };
      if (formData.value.filters !== null) {
        tagMapping.value = true;
      } else {
        formData.value.filters = [];
      }
    }
  };

  const handleOk = async () => {
    const res = await formRef.value?.validate();
    if (!res) {
      try {
        if (formData.value.id === 0) {
          await submitIntegrationRouterForm(
            sourceId.value as number,
            formData.value
          );
          Message.success('路由添加成功');
        } else {
          await updateIntegrationRouterForm(formData.value.id, formData.value);
          Message.success('路由编辑成功');
        }

        emits('refresh', sourceId.value);
        handleCancel();
      } catch (err) {
        //
      }
    }
  };

  defineExpose({
    show,
  });
</script>

<style scoped lang="less">
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
