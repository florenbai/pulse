<template>
  <a-drawer
    :width="600"
    :visible="visible"
    unmount-on-close
    @ok="handleOk"
    @cancel="handleCancel"
  >
    <template #title> 配置标签重写 </template>
    <a-alert
      >标签重写会先于路由匹配执行。当源标签名称与新标签名称相同时，系统将覆盖源标签。当源标签不存在时，系统不会执行标签重写。</a-alert
    >
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
          <div v-if="tagMapping" style="display: flex">
            <div
              v-if="
                formData.filters !== undefined && formData.filters.length > 1
              "
              class="rule-condition-logic-group__logic"
            >
              <div
                class="rule-condition-logic-group__logic-item"
                :style="{
                  height: `${tagHeight}px`,
                }"
              >
                <div
                  class="rule-condition-logic-group__logic-item__start"
                  style="width: 12px"
                ></div>
                <span class="rule-condition-logic-group__logic-item__text"
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
                v-for="(item, index) of formData.filters"
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
                  style="width: 100px; margin-left: 5px; margin-right: 5px"
                >
                  <a-option value="IN">包含</a-option>
                  <a-option value="NOT IN">不包含</a-option>
                </a-select>
                <a-input-tag
                  v-model="item.value"
                  allow-clear
                  :style="{ width: '240px' }"
                  placeholder="值"
                />
                <a-button
                  style="margin-left: 5px; background-color: #fff"
                  @click="deleteFilterItem(index)"
                >
                  <template #icon> <icon-delete size="15" /> </template
                ></a-button>
              </a-input-group>
            </div>
          </div>
          <a-button v-if="tagMapping" type="text" @click="createCondition"
            >+ 添加规则</a-button
          >
        </a-card>
      </div>
      <a-form-item
        field="rewriteType"
        label="重写类型"
        validate-trigger="input"
        :rules="[
          {
            required: true,
            message: '请选则转发工作区',
          },
        ]"
      >
        <a-select v-model="formData.rewriteType" placeholder="请选择重写类型">
          <a-option value="extract">提取</a-option>
          <a-option value="rewrite">重写</a-option>
          <a-option value="delete">删除</a-option>
        </a-select>
      </a-form-item>
      <div
        v-if="
          formData.rewriteType === 'extract' ||
          formData.rewriteType === 'rewrite'
        "
      >
        <a-form-item
          field="rewriteItem.oldTag"
          label="来源标签"
          validate-trigger="input"
          :rules="[
            {
              required: true,
              message: '请输入来源标签',
            },
          ]"
        >
          <a-input
            v-model="formData.rewriteItem.oldTag"
            placeholder="请输入来源标签"
          />
        </a-form-item>
        <a-form-item
          field="rewriteItem.expression"
          label="提取正则表达式"
          validate-trigger="input"
          tooltip="您可前往 https://www.jyshare.com/front-end/854/ 调试您的表达式。"
          :rules="[
            {
              required: true,
              message: '请输入提取值的正则表达式',
            },
          ]"
        >
          <a-input
            v-model="formData.rewriteItem.expression"
            placeholder="请输入提取值的正则表达式"
          />
        </a-form-item>
        <a-form-item
          v-if="formData.rewriteType === 'rewrite'"
          field="rewriteItem.value"
          label="替换值"
          validate-trigger="input"
          :rules="[
            {
              required: true,
              message: '请输入替换值',
            },
          ]"
        >
          <a-input
            v-model="formData.rewriteItem.value"
            placeholder="请输入替换值"
          />
        </a-form-item>
        <a-form-item
          field="rewriteItem.newTag"
          label="新标签"
          validate-trigger="input"
          :rules="[
            {
              required: true,
              message: '请输入新标签名称',
            },
          ]"
        >
          <a-input
            v-model="formData.rewriteItem.newTag"
            placeholder="请输入新标签名称"
          />
        </a-form-item>
        <a-card v-if="formData.rewriteType === 'extract'" title="样例">
          <p>假设告警标签 desc: 高防流量504异常1小时内告警8次。</p>
          <p
            >通过来源标签<a-tag>desc</a-tag>，使用正则表达式：/1小时内告警([0-9])次/提取数字8，生成新的标签<a-tag
              >value</a-tag
            ></p
          >
          <p
            >最终得到新告警标签 desc:
            高防流量504异常1小时内告警8次，value:8。</p
          >
        </a-card>
        <a-card v-if="formData.rewriteType === 'rewrite'" title="样例">
          <p>假设告警标签 desc: 高防流量504异常1小时内告警8次。</p>
          <p
            >通过来源标签<a-tag>desc</a-tag>，使用正则表达式：<a-tag>/高防流量([0-9]+)异常/</a-tag>匹配数字504。</p
          >
          <p>将匹配的数字替换成新的值，如httpCode:504</p>
          <p
            >最终得到新告警标签 desc:
            高防流量httpCode:504异常1小时内告警8次。</p
          >
        </a-card>
      </div>
      <div v-else-if="formData.rewriteType === 'delete'">
        <a-form-item
          field="rewriteItem.deleteTag"
          label="删除标签"
          validate-trigger="input"
          :rules="[
            {
              required: true,
              message: '请输入要删除的标签',
            },
          ]"
        >
          <a-input-tag
            v-model="formData.rewriteItem.deleteTag"
            allow-clear
            placeholder="请输入要删除的标签"
          />
        </a-form-item>
      </div>
    </a-form>
  </a-drawer>
</template>

<script lang="ts" setup>
  import {
    TagRewriteItem,
    submitTagRewriteForm,
    updateTagRewriteForm,
  } from '@/api/integration';
  import { FormInstance, Message } from '@arco-design/web-vue';
  import { ref } from 'vue';

  const emits = defineEmits(['refresh']);
  const sourceId = ref<number>();
  const visible = ref(false);
  const tagMapping = ref(false);
  const tagHeight = ref<number>(40);
  const formRef = ref<FormInstance>();
  const formData = ref<TagRewriteItem>({
    id: 0,
    rewriteType: 'extract',
    filters: [],
    rewriteItem: {},
  });

  const handleChangeTagMapping = () => {
    if (tagMapping.value) {
      tagMapping.value = false;
      formData.value.filters = [];
    } else {
      if (formData.value.filters.length === 0) {
        formData.value.filters.push({
          tag: '',
          operation: 'IN',
          value: [],
        });
      }
      tagMapping.value = true;
    }
  };

  const createCondition = () => {
    formData.value.filters.push({
      tag: '',
      operation: 'IN',
      value: [],
    });
    if (formData.value.filters.length > 2) {
      tagHeight.value = (formData.value.filters.length - 2) * 40 + 40;
    }
  };

  const deleteFilterItem = (index: number) => {
    if (index === 0) {
      formData.value.filters.shift();
    } else {
      formData.value.filters.splice(index, 1);
    }
    if (formData.value.filters.length >= 2) {
      tagHeight.value -= 40;
    }
  };

  const handleCancel = () => {
    formData.value = {
      id: 0,
      rewriteType: 'extract',
      filters: [],
      rewriteItem: {},
    };
    visible.value = false;
    tagMapping.value = false;
  };

  const show = (id: number, item?: TagRewriteItem) => {
    sourceId.value = id;
    visible.value = true;
    if (item !== undefined) {
      formData.value = { ...item };
      if (formData.value.filters !== null) {
        tagMapping.value = true;
        if (formData.value.filters.length > 2) {
          tagHeight.value = (formData.value.filters.length - 2) * 40 + 40;
        }
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
          await submitTagRewriteForm(sourceId.value as number, formData.value);
          Message.success('重写规则添加成功');
        } else {
          await updateTagRewriteForm(formData.value.id, formData.value);
          Message.success('重写规则编辑成功');
        }

        emits('refresh');
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
