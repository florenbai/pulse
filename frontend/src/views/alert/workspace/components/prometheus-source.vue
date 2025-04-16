<template>
  <a-drawer
    :width="340"
    :visible="props.visible"
    unmount-on-close
    @ok="handleOk"
    @cancel="handleCancel"
  >
    <template #title> Prometheus告警源配置 </template>
    <div
      >You can customize modal body text by the current situation. This modal
      will be closed immediately once you press the OK button.
    </div>
  </a-drawer>
</template>

<script lang="ts" setup>
  export interface Props {
    visible: boolean;
  }
  const props = defineProps<Props>();
  const emits = defineEmits(['changeVisible']);
  const handleOk = () => {
    emits('changeVisible', 'Prometheus');
  };
  const handleCancel = () => {
    emits('changeVisible', 'Prometheus');
  };
</script>
