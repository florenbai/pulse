<template>
  <a-drawer
    :visible="visible"
    title="配置权限"
    :width="600"
    @ok="handleOk"
    @cancel="handleCancel"
  >
    <div>
      <a-alert>权限分为菜单权限和工作区数据权限。</a-alert>
    </div>

    <a-form ref="formRef" :model="formData" layout="vertical">
      <a-collapse :default-active-key="['1', '2']" :bordered="false">
        <a-collapse-item key="1" header="菜单权限" style="padding-left: 0px">
          <template #extra>
            <a-checkbox
              v-model="menuAllChecked"
              style="margin-right: 15px"
              @click.stop="toggleChecked"
            >
              全选</a-checkbox
            >
          </template>
          <a-form-item field="menu" label="">
            <a-tree
              ref="ruleTreeRef"
              v-model:checked-keys="formData.menus"
              :checkable="true"
              :default-expand-all="true"
              block-node
              :field-names="{
                key: 'key',
                title: 'name',
                children: 'children',
              }"
              :data="menuList"
              @check="onCheck"
            />
          </a-form-item>
        </a-collapse-item>
        <a-collapse-item key="2" header="工作区权限" style="padding-left: 0px">
          <a-radio-group v-model="formData.workspace" :options="dataOptions" />
        </a-collapse-item>
      </a-collapse>
    </a-form>
  </a-drawer>
</template>

<script lang="ts" setup>
  import { computed, ref } from 'vue';
  import { FormInstance } from '@arco-design/web-vue/es/form';
  import { Message, TreeNodeData } from '@arco-design/web-vue';
  import { MenuPermission } from '@/api/menu';
  import { queryAuthorizationInfo, teamPermission } from '@/api/team';
  import useRequest from '@/hooks/request';
  import { getMenuNameList } from '@/api/user';

  const { response: data } = useRequest<TreeNodeData[]>(getMenuNameList);
  const menuList = computed<TreeNodeData[]>(() => {
    const menuData = [] as any[];
    if (data.value === undefined) {
      return menuData;
    }
    menuData.push(...data.value);
    return menuData;
  });

  const visible = ref(false);
  const itemId = ref();
  const emits = defineEmits(['refresh']);
  const formRef = ref<FormInstance>();
  const menuAllChecked = ref(false);
  const dataOptions = [
    { label: '本团队工作区', value: 1 },
    { label: '全部工作区', value: 2 },
  ];
  const ruleTreeRef = ref();
  const formData = ref<MenuPermission>({
    menus: [],
    workspace: 1,
  });
  const show = async (id: number) => {
    itemId.value = id;
    visible.value = true;
    const res = await queryAuthorizationInfo(itemId.value);
    formData.value.workspace = res.data.workspace;
    formData.value.menus = res.data.menus;
  };

  const handleCancel = () => {
    visible.value = false;
    itemId.value = undefined;
    formData.value = { menus: [], workspace: 1 };
  };

  const handleOk = async () => {
    const res = await formRef.value?.validate();
    if (!res) {
      try {
        await teamPermission(itemId.value, formData.value);
        Message.success('授权成功');

        handleCancel();
        emits('refresh');
      } catch (err) {
        //
      }
    }
  };

  // 全选
  const toggleChecked = () => {
    if (menuAllChecked.value) {
      menuAllChecked.value = true;
      ruleTreeRef.value.checkAll(false);
    } else {
      menuAllChecked.value = false;
      ruleTreeRef.value.checkAll(true);
    }
  };
  const onCheck = (newCheckedKeys: any) => {
    formData.value.menus = newCheckedKeys;
  };
  defineExpose({
    show,
  });
</script>

<style lang="less" scoped>
  .menu_permission {
    width: 100%;
    .haeder {
      border-bottom: var(--color-neutral-4) solid 1px;
    }
  }
</style>
