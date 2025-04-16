<template>
  <a-card
    class="general-card chat-panel"
    title="规则分组"
    :bordered="false"
    :header-style="{ paddingBottom: '0' }"
    :body-style="{
      height: '100%',
      paddingTop: '16px',
      display: 'flex',
      flexFlow: 'column',
    }"
  >
    <a-space :size="8">
      <a-input-search
        v-model="searchKey"
        placeholder="请输入分组名称"
        @change="searchGroup"
        @search="searchGroup"
      />
      <a-button type="text">
        <icon-plus @click="openGroupDrawer" />
      </a-button>
    </a-space>
    <div class="chat-panel-content">
      <a-spin :loading="loading" style="width: 100%">
        <a-tree
          ref="treeRef"
          :data="treeData"
          :default-selected-keys="defaultSelect"
          :field-names="{
            key: 'id',
            title: 'groupName',
          }"
          @select="selectTree"
        >
          <template #extra="nodeData">
            <a-popconfirm
              content="删除规则组后，该组的所有规则也将删除，您是否要删除?"
              @ok="onIconClick(nodeData)"
            >
              <IconDelete
                style="
                  position: absolute;
                  right: 8px;
                  font-size: 12px;
                  top: 10px;
                  color: #3370ff;
                "
              />
            </a-popconfirm>
          </template>
        </a-tree>
      </a-spin>
    </div>
  </a-card>
  <RuleGroupDrawer
    ref="ruleGroupDrawer"
    @refresh="fetchDataAndSelectLast"
  ></RuleGroupDrawer>
</template>

<script lang="ts" setup>
  import { ref } from 'vue';
  import useLoading from '@/hooks/loading';
  import {
    queryAlertRuleGroup,
    AlertRuleGroupRecode,
    deleteAlertRuleGroup,
  } from '@/api/alert-rule';
  import { Message } from '@arco-design/web-vue';
  import { TreeNodeKey } from '@arco-design/web-vue/es/tree/interface';
  import RuleGroupDrawer from './rule-group-drawer.vue';

  const { loading, setLoading } = useLoading(true);
  const treeData = ref<AlertRuleGroupRecode[]>([]);
  const baseTreeData = ref<AlertRuleGroupRecode[]>([]);
  const defaultSelect = ref<number[]>([]);
  const emits = defineEmits(['selectRuleGroup']);
  const searchKey = ref('');
  const ruleGroupDrawer = ref();
  const treeRef = ref();

  const fetchData = async () => {
    try {
      const { data } = await queryAlertRuleGroup();
      treeData.value = data;
      baseTreeData.value = data;
      if (treeData.value.length > 0) {
        defaultSelect.value.push(treeData.value[0].id);
        treeRef.value.selectNode(treeData.value[0].id as TreeNodeKey, true);
        emits(
          'selectRuleGroup',
          treeData.value[0].id,
          treeData.value[0].groupName
        );
      }
    } catch (err) {
      // you can report use errorHandler or other
    } finally {
      setLoading(false);
    }
  };

  const fetchDataAndSelectLast = async () => {
    try {
      const { data } = await queryAlertRuleGroup();
      treeData.value = data;
      baseTreeData.value = data;
      if (treeData.value.length > 0) {
        emits(
          'selectRuleGroup',
          treeData.value[treeData.value.length - 1].id,
          treeData.value[treeData.value.length - 1].groupName
        );
        setTimeout(() => {
          treeRef.value.selectNode(
            treeData.value[treeData.value.length - 1].id as TreeNodeKey,
            true
          );
        }, 500);
      }
    } catch (err) {
      // you can report use errorHandler or other
    } finally {
      setLoading(false);
    }
  };
  fetchData();

  const searchGroup = (val: string) => {
    if (val === '') {
      treeData.value = baseTreeData.value;
    } else {
      const result = [] as AlertRuleGroupRecode[];
      baseTreeData.value.forEach((item) => {
        if (item.groupName.toLowerCase().indexOf(val) > -1) {
          result.push({ ...item });
        }
      });
      treeData.value = result;
    }
  };

  const openGroupDrawer = () => {
    ruleGroupDrawer.value.show();
  };

  const selectTree = (selectedKeys: Array<string | number>) => {
    treeData.value.forEach((val) => {
      if (val.id === selectedKeys[0]) {
        emits('selectRuleGroup', val.id, val.groupName);
      }
    });
  };

  const onIconClick = async (nodeData: any) => {
    try {
      await deleteAlertRuleGroup(nodeData.id);
      Message.success({
        content: '操作成功',
        duration: 5 * 1000,
      });
      fetchData();
    } catch (err) {
      // you can report use errorHandler or other
    } finally {
      setLoading(false);
    }
  };
</script>

<style scoped lang="less">
  .chat-panel {
    display: flex;
    flex-direction: column;
    height: 100%;
    // padding: 20px;
    background-color: var(--color-bg-2);

    &-content {
      flex: 1;
      margin: 20px 0;
    }
  }
</style>
