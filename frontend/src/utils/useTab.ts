import { useRoute } from 'vue-router';
import { ref, watchEffect } from 'vue';
import router from '@/router';
import { isEmptyObject } from './is';

export * from './is';

/**
 * 获取一个tab值和切换Tab方法
 * @param defaultTab 默认tab值
 */
function useTab(defaultTab?: string, tabName?: string) {
  if (isEmptyObject(tabName)) {
    tabName = 'tab';
  }
  const route = useRoute();
  const tab = ref<string>(String(route.query[tabName as string] ?? defaultTab));

  watchEffect(() => {
    const curTab = String(route.query[tabName as string] ?? defaultTab);
    if (curTab) tab.value = curTab;
  });
  const changTab = (value: string) => {
    const q = { ...route.query };
    q[`${tabName}`] = value;
    router.replace({ query: q });
  };
  return [tab, changTab] as const;
}
export default useTab;
