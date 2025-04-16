<template>
  <div class="container">
    <Breadcrumb :items="[{ label: 'menu.channel' }]" />
    <a-row :gutter="20" align="stretch">
      <a-col :span="24">
        <a-card class="general-card" title="告警渠道列表">
          <a-row style="margin-bottom: 16px">
            <a-col :span="12">
              <a-space>
                <a-button type="primary" @click="handleClick"
                  ><template #icon> <icon-plus /> </template
                  >新增告警渠道</a-button
                >
              </a-space>
            </a-col>
          </a-row>
          <a-table :columns="columns" :data="renderData" :loading="loading">
            <template #channelType="{ record }">
              <a-tag color="blue"> {{ record.channelType }}</a-tag>
            </template>
            <template #operations="{ record }">
              <a-space>
                <a-button type="text" size="small" @click="editChannel(record)">
                  编辑
                </a-button>
                <a-popconfirm
                  content="是否确认要删除通知渠道?"
                  @ok="deleteChannelRecord(record.id)"
                >
                  <a-button type="text" size="small"> 删除 </a-button>
                </a-popconfirm>
              </a-space>
            </template>
          </a-table>
          <a-pagination
            v-if="
              pagination.total != undefined &&
              pagination.total - pagination.pageSize > 0
            "
            :total="(pagination.total as number)"
          />
        </a-card>
      </a-col>
    </a-row>
    <channelDrawer ref="channelDrawerRef" @refresh="fetchData"></channelDrawer>
  </div>
</template>

<script lang="ts" setup>
  import { reactive, ref } from 'vue';
  import { Pagination } from '@/types/global';
  import useLoading from '@/hooks/loading';
  import {
    AlertTemplateParams,
    queryAlertChannelList,
    AlertChannelRecode,
    deleteChannel,
  } from '@/api/alert-template';
  import { Message } from '@arco-design/web-vue';
  import channelDrawer from './components/channel-drawer.vue';

  const { loading, setLoading } = useLoading(true);
  const renderData = ref<AlertChannelRecode[]>([]);
  const channelDrawerRef = ref();
  const basePagination: Pagination = {
    page: 1,
    pageSize: 20,
    total: 1,
  };
  const columns = [
    {
      title: '告警渠道名称',
      dataIndex: 'channelName',
    },
    {
      title: '告警渠道类型',
      dataIndex: 'channelType',
      slotName: 'channelType',
    },
    {
      title: '告警渠道标识',
      dataIndex: 'channelSign',
    },
    {
      title: '告警渠道分组',
      dataIndex: 'channelGroup',
    },
    {
      title: '操作',
      dataIndex: 'operations',
      slotName: 'operations',
    },
  ];
  const pagination = reactive({
    ...basePagination,
  });
  const fetchData = async (
    params: AlertTemplateParams = { page: 1, pageSize: 20 }
  ) => {
    setLoading(true);
    try {
      const { data } = await queryAlertChannelList(params);
      renderData.value = data.list;
      pagination.page = params.page;
      pagination.total = data.total;
    } catch (err) {
      // you can report use errorHandler or other
    } finally {
      setLoading(false);
    }
  };
  fetchData();
  const handleClick = () => {
    channelDrawerRef.value.show();
  };

  const editChannel = (item: AlertChannelRecode) => {
    channelDrawerRef.value.show(item);
  };

  const deleteChannelRecord = async (id: number) => {
    setLoading(true);
    try {
      await deleteChannel(id);
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

<script lang="ts">
  export default {
    name: 'SearchTable',
  };
</script>

<style scoped lang="less">
  .container {
    padding: 0 20px 20px 20px;
  }
  :deep(.arco-table-th) {
    &:last-child {
      .arco-table-th-item-title {
        margin-left: 16px;
      }
    }
  }
  .action-icon {
    margin-left: 12px;
    cursor: pointer;
  }
  .active {
    color: #0960bd;
    background-color: #e3f4fc;
  }
  .setting {
    display: flex;
    align-items: center;
    width: 200px;
    .title {
      margin-left: 12px;
      cursor: pointer;
    }
  }
  .arco-btn.arco-btn-text {
    padding: 0;
    height: auto;
  }

  .list-wrap {
    margin-top: 30px;
  }
</style>
