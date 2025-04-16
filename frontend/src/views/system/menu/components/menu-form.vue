<template>
  <a-drawer
    :visible="visible"
    title="配置菜单"
    :width="600"
    @ok="handleOk"
    @cancel="handleCancel"
  >
    <a-form ref="formRef" :model="formData" layout="vertical">
      <a-form-item
        field="name"
        label="菜单名称"
        :rules="[
          {
            required: true,
            message: '请输入菜单名称',
          },
        ]"
      >
        <a-input v-model="formData.name" placeholder="请输入菜单名称" />
      </a-form-item>
      <a-form-item field="type" label="菜单类型">
        <a-radio-group v-model="formData.type" type="button">
          <a-radio :value="0">目录</a-radio>
          <a-radio :value="1">菜单</a-radio>
          <a-radio :value="2">按钮</a-radio>
        </a-radio-group>
      </a-form-item>
      <a-form-item field="pid" label="上级菜单">
        <MenuTree v-model="formData.pid"></MenuTree>
      </a-form-item>
      <a-form-item
        v-if="formData.type === 2"
        field="permission"
        label="权限标识"
        tooltip='请使用 v-permission="[admin]" 等定义权限'
        :rules="[
          {
            required: true,
            message: '请输入权限标识',
          },
        ]"
      >
        <a-input v-model="formData.permission" placeholder="请输入权限标识" />
      </a-form-item>
      <a-form-item v-if="formData.type !== 2" label="图标">
        <a-input-search
          v-model="formData.icon"
          placeholder="选择图标/填写"
          search-button
          allow-clear
        >
          <template v-if="formData.icon" #prefix>
            <Icon :icon="formData.icon" :size="23"></Icon>
          </template>
          <template #button-icon>
            <a-popover position="bl" trigger="click">
              <icon-apps />
              <template #content>
                <IconPicker
                  @change="
                    (icon) => {
                      formData.icon = icon;
                    }
                  "
                ></IconPicker>
              </template>
            </a-popover>
          </template>
        </a-input-search>
      </a-form-item>
      <a-form-item
        v-if="formData.type !== 2"
        field="routeName"
        label="路由名称"
        :rules="[
          {
            required: true,
            message: '请输入路由名称',
          },
        ]"
      >
        <a-input v-model="formData.routeName" placeholder="请输入路由名称" />
      </a-form-item>
      <a-form-item
        v-if="formData.type !== 2"
        field="routePath"
        label="路由地址"
        :rules="[
          {
            required: true,
            message: '请输入路由地址',
          },
        ]"
      >
        <a-input v-model="formData.routePath" placeholder="请输入路由地址" />
      </a-form-item>
      <a-form-item
        v-if="formData.type !== 2"
        field="locale"
        label="多语言标识"
        :rules="[
          {
            required: true,
            message: '请输入多语言标识',
          },
        ]"
      >
        <a-input v-model="formData.locale" placeholder="请输入多语言标识" />
      </a-form-item>
      <a-form-item
        v-if="formData.type === 1"
        field="isLink"
        label="是否外链"
        tooltip="在浏览器新标签打开"
      >
        <a-switch v-model="formData.isLink">
          <template #checked> 是 </template>
          <template #unchecked> 否 </template>
        </a-switch>
      </a-form-item>
      <a-form-item
        v-if="formData.type !== 2"
        field="hideMenu"
        label="隐藏菜单"
        tooltip="是否在左侧菜单中隐藏"
      >
        <a-switch v-model="formData.hideMenu">
          <template #checked> 是 </template>
          <template #unchecked> 否 </template>
        </a-switch>
      </a-form-item>
      <a-form-item
        v-if="formData.type !== 2"
        field="hideChildMenu"
        label="显示单项"
        tooltip="是否强制在左侧菜单中显示单项"
      >
        <a-switch v-model="formData.hideChildMenu">
          <template #checked> 是 </template>
          <template #unchecked> 否 </template>
        </a-switch>
      </a-form-item>
      <a-form-item
        v-if="formData.type !== 2"
        field="requiresAuth"
        label="登录认证"
        tooltip="是否需要登录认证"
      >
        <a-switch v-model="formData.requiresAuth">
          <template #checked> 是 </template>
          <template #unchecked> 否 </template>
        </a-switch>
      </a-form-item>
      <a-form-item
        v-if="formData.type !== 2"
        field="activeMenu"
        label="高亮显示"
        tooltip="高亮设置的菜单项"
      >
        <a-switch v-model="formData.activeMenu">
          <template #checked> 是 </template>
          <template #unchecked> 否 </template>
        </a-switch>
      </a-form-item>
      <a-form-item
        v-if="formData.type !== 2"
        field="cache"
        label="排除缓存"
        tooltip="ignoreCache设置为true页面将不会被缓存"
      >
        <a-switch v-model="formData.cache">
          <template #checked> 是 </template>
          <template #unchecked> 否 </template>
        </a-switch>
      </a-form-item>
    </a-form>
  </a-drawer>
</template>

<script lang="ts" setup>
  import { ref } from 'vue';
  import { FormInstance } from '@arco-design/web-vue/es/form';
  import { Message } from '@arco-design/web-vue';
  import {
    editMenuForm,
    MenuForm,
    MenuRecode,
    submitMenuForm,
  } from '@/api/menu';
  import { IconPicker, Icon } from '@/components/Icon';
  import MenuTree from './menu-tree.vue';

  const visible = ref(false);
  const itemId = ref();
  const emits = defineEmits(['refresh']);
  const formRef = ref<FormInstance>();

  const formData = ref<MenuForm>({
    name: '',
    pid: 0,
    type: 1,
    icon: '',
    component: '',
    routePath: '',
    routeName: '',
    redirect: '',
    sort: 0,
    requiresAuth: false,
    permission: '',
    locale: '',
    hideMenu: false,
    hideChildMenu: false,
    isLink: false,
    cache: false,
    affix: false,
    enabled: false,
    activeMenu: false,
  });
  const show = (item: MenuRecode | undefined) => {
    if (item !== undefined) {
      formData.value = { ...item };
      itemId.value = item.id;
    }
    visible.value = true;
  };

  const handleCancel = () => {
    visible.value = false;
    itemId.value = undefined;
    formData.value = { type: 1 } as MenuForm;
  };

  const handleOk = async () => {
    const res = await formRef.value?.validate();
    if (!res) {
      try {
        if (itemId.value !== undefined) {
          await editMenuForm(itemId.value, formData.value);
          Message.success('菜单编辑成功');
        } else {
          await submitMenuForm(formData.value);
          Message.success('菜单添加成功');
        }

        handleCancel();
        emits('refresh');
      } catch (err) {
        //
      }
    }
  };

  defineExpose({
    show,
  });
</script>
