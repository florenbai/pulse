<template>
  <a-card :bordered="false">
    <a-descriptions title="标签重写" layout="horizontal"></a-descriptions>
    <a-grid
      :cols="{ xs: 1, sm: 2, md: 3, lg: 3, xl: 4, xxl: 4 }"
      :col-gap="12"
      :row-gap="16"
      class="grid"
    >
      <a-grid-item
        v-for="(element, index) in tagData"
        :key="index"
        class="grid-item"
        style="border-style: dashed"
        @click="editTag(element)"
      >
        <span class="right-link-hover">
          <a-button
            type="text"
            @click.stop="
              (event) => {
                deleteTag(element, event);
              }
            "
          >
            <template #icon>
              <icon-delete size="15" />
            </template>
          </a-button>
        </span>
        <div class="card-wrap">
          <a-space align="start" direction="vertical" fill>
            <div
              v-if="element.filters != undefined && element.filters.length > 0"
            >
              <div style="margin-bottom: 10px">根据以下条件</div>
              <a-row
                v-for="(f, i) in element.filters"
                :key="i"
                justify="space-between"
              >
                <a-space wrap>
                  <a-tag bordered>{{ f.tag }}</a-tag>
                  <a-tag bordered color="orangered">{{
                    f.operation === 'IN' ? '包含' : '不包含'
                  }}</a-tag>
                  <a-tag v-for="(v, k) in f.value" :key="k" color="blue">{{
                    v
                  }}</a-tag>
                </a-space>
              </a-row>
            </div>
            <div v-else>无需限定条件</div>
            <div
              v-if="element.rewriteType === 'extract'"
              style="margin-top: 10px"
            >
              使用正则表达式
              <a-tag color="blue">{{ element.rewriteItem.expression }}</a-tag
              >，<a-tag color="#f53f3f"> 提取 </a-tag> 标签
              <a-tag>{{ element.rewriteItem.oldTag }}</a-tag>
              的值，并将提取的值写入标签<a-tag>{{
                element.rewriteItem.newTag
              }}</a-tag
              >。当没有提取到值时，将不会新增标签。
            </div>
            <div
              v-else-if="element.rewriteType === 'rewrite'"
              style="margin-top: 10px"
            >
              使用正则表达式
              <a-tag color="blue">{{ element.rewriteItem.expression }}</a-tag
              >，<a-tag color="#f53f3f"> 匹配 </a-tag> 标签
              <a-tag>{{ element.rewriteItem.oldTag }}</a-tag>
              的值，将匹配的值替换成
              <a-tag>{{ element.rewriteItem.value }}</a-tag>
              ,并写入标签<a-tag>{{ element.rewriteItem.newTag }}</a-tag
              >。当没有匹配到值时，将不会重写。
            </div>
            <div
              v-else-if="element.rewriteType === 'delete'"
              style="margin-top: 10px"
            >
              删除标签
              <a-tag
                v-for="(item, k) in element.rewriteItem.deleteTag"
                :key="k"
                color="blue"
                >{{ item }}</a-tag
              >
            </div>
          </a-space>
        </div>
      </a-grid-item>
      <a-grid-item
        class="grid-item"
        style="border-style: dashed"
        @click="handleClick"
      >
        <div style="text-align: center; margin: auto">
          <icon-plus style="margin-right: 1rem" /> <span>标签重写</span></div
        >
      </a-grid-item>
    </a-grid>
  </a-card>
  <TagDrawer ref="tagDrawerRef" @refresh="getAllTag"></TagDrawer>
</template>

<script lang="ts" setup>
  import { h, onMounted, ref } from 'vue';
  import {
    TagRewriteItem,
    queryAllTagRewrite,
    deleteTagRewriteRecord,
  } from '@/api/integration';
  import { Message, Notification } from '@arco-design/web-vue';
  import TagDrawer from './tag-drawer.vue';

  export interface Props {
    integration: number | undefined;
  }
  const props = defineProps<Props>();
  const tagData = ref<TagRewriteItem[]>();
  const tagDrawerRef = ref();
  const handleClick = () => {
    tagDrawerRef.value.show(props.integration);
  };

  const getAllTag = async () => {
    try {
      const res = await queryAllTagRewrite(props.integration as number);
      tagData.value = res.data as TagRewriteItem[];
    } catch {
      //
    }
  };

  const editTag = (item: TagRewriteItem) => {
    tagDrawerRef.value.show(props.integration, item);
  };

  const deleteTag = (item: TagRewriteItem, event: any) => {
    const x = `${event.clientX - 340}px`;
    const y = `${event.clientY - 100}px`;
    const now = `${Date.now()}`;
    Notification.info({
      id: now,
      title: '标签重写删除提醒',
      position: 'topLeft',
      duration: 0,
      content: '请确认是否要删除该重写规则',
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
                Notification.remove(now);
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
                  await deleteTagRewriteRecord(item.id);
                  Message.success('规则删除成功');
                  Notification.remove(now);
                  getAllTag();
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
    getAllTag();
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
    min-height: 150px;
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
</style>
