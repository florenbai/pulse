<template>
  <div class="container">
    <Breadcrumb
      :items="[
        { label: 'menu.alert.template', path: '/owl/template/list' },
        { label: 'menu.alert.template.detail' },
      ]"
    />
    <a-card class="general-card" :title="title">
      <a-table
        row-key="index"
        :loading="loading"
        :pagination="false"
        :columns="columns"
        :data="renderData"
        :bordered="false"
      >
        <template #finished="{ record }">
          <a-space>
            <a-tag v-if="record.finished === false">未配置</a-tag>
            <a-tag v-else color="green">已配置</a-tag>
          </a-space>
        </template>
        <template #operations="{ record }">
          <a-space>
            <a-button type="text" size="small" @click="open(record)">
              配置
            </a-button>
          </a-space>
        </template>
      </a-table>
    </a-card>
    <channelDrawer ref="channelDrawerRef" @refresh="getData"></channelDrawer>
  </div>
</template>

<script setup lang="ts">
  import { queryAlertTemplateDetail } from '@/api/alert-template';
  import useLoading from '@/hooks/loading';
  import { onMounted, ref, watch } from 'vue';
  import { useRoute } from 'vue-router';
  import channelDrawer from './components/channel-drawer.vue';

  const { loading, setLoading } = useLoading(true);
  const renderData = ref();
  const id = ref<number>();
  const channelDrawerRef = ref();
  const templateData = ref();
  const title = ref();
  const route = useRoute();

  watch(
    route,
    () => {
      if (route.params.id) {
        id.value = Number(route.params?.id);
      }
    },
    {
      deep: true,
      immediate: true,
    }
  );
  const columns = [
    {
      title: '通知渠道',
      dataIndex: 'channelName',
    },
    {
      title: '状态',
      dataIndex: 'finished',
      slotName: 'finished',
    },
    {
      title: '操作',
      dataIndex: 'operations',
      slotName: 'operations',
    },
  ];

  const open = (t: any) => {
    channelDrawerRef.value.show(t);
  };

  const getData = async () => {
    try {
      const { data } = await queryAlertTemplateDetail(id.value as number);
      renderData.value = data.channels;
      templateData.value = data;
      title.value = data.templateName;
      setLoading(false);
    } catch (err) {
      // you can report use errorHandler or other
    }
  };

  onMounted(() => {
    getData();
  });
</script>

<style scoped lang="less">
  .container {
    padding: 0 20px 20px 20px;
    :deep(.arco-list-content) {
      overflow-x: hidden;
    }

    :deep(.arco-card-meta-title) {
      font-size: 14px;
    }
  }
  .tabs-custom {
    padding-top: 20px;
  }
  .tabs-custom ::before {
    height: 0px !important;
  }
</style>
