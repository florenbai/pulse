<template>
  <a-config-provider :locale="locale">
    <router-view v-if="routerAlive" />
    <global-setting />
  </a-config-provider>
</template>

<script lang="ts" setup>
  import { computed, nextTick, provide, ref } from 'vue';
  import enUS from '@arco-design/web-vue/es/locale/lang/en-us';
  import zhCN from '@arco-design/web-vue/es/locale/lang/zh-cn';
  import GlobalSetting from '@/components/global-setting/index.vue';
  import useLocale from '@/hooks/locale';

  const { currentLocale } = useLocale();
  const routerAlive = ref(true);
  const locale = computed(() => {
    switch (currentLocale.value) {
      case 'zh-CN':
        return zhCN;
      case 'en-US':
        return enUS;
      default:
        return enUS;
    }
  });
  const reload = () => {
    routerAlive.value = false;
    nextTick(() => {
      routerAlive.value = true;
    });
  };
  provide('reload', reload);
</script>
