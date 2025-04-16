<template>
  <div style="margin-left: 40px">
    <a-grid
      :cols="{ xs: 1, sm: 2, md: 3, lg: 3, xl: 4, xxl: 4 }"
      :col-gap="12"
      :row-gap="16"
      class="grid"
    >
      <a-grid-item
        v-for="item in alertSource"
        :key="item.id"
        class="grid-item"
        @click="checkSource(item.id)"
      >
        <div class="logo"> <icon-font :type="item.icon" :size="32" /></div>
        <div class="text-center text-l1 text-main">
          {{ item.sourceName }}
        </div>
        <div v-if="alertSourceModel.alertSource.includes(item.id)">
          <div class="grid-item-checked-bg"></div>
          <div class="grid-item-checked"><icon-check /> </div
        ></div>
      </a-grid-item>
    </a-grid>
  </div>
</template>

<script lang="ts" setup>
  import { AlertSourceRecode, queryAllAlertSource } from '@/api/alert-source';
  import { onMounted, ref } from 'vue';
  import { Icon } from '@arco-design/web-vue';
  import { AlertSourceRequest } from '@/api/workspace';

  const IconFont = Icon.addFromIconFontCn({
    src: 'https://at.alicdn.com/t/c/font_4395376_zlyzz5aqyk.js?spm=a313x.manage_type_myprojects.i1.10.15d63a81Wp8j8D&file=font_4395376_zlyzz5aqyk.js',
  });

  const emits = defineEmits(['changeStep']);
  const alertSource = ref<AlertSourceRecode[]>([]);
  const alertSourceModel = ref<AlertSourceRequest>({ alertSource: [] });
  onMounted(async () => {
    try {
      const res = await queryAllAlertSource();
      alertSource.value = res.data;
    } catch (err) {
      //
    }
  });

  const checkSource = (id: number) => {
    if (alertSourceModel.value.alertSource.includes(id)) {
      alertSourceModel.value.alertSource =
        alertSourceModel.value.alertSource.filter((element) => element !== id);
    } else {
      alertSourceModel.value.alertSource.push(id);
    }
  };

  const submitAlertSource = async (): Promise<boolean> => {
    if (alertSourceModel.value.alertSource.length < 1) {
      return false;
    }
    emits('changeStep', { ...alertSourceModel.value });
    return true;
  };

  defineExpose({
    submitAlertSource,
  });
</script>

<style scoped>
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
</style>
