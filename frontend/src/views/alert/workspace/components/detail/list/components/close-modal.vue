<template>
  <a-button :type="props.type" @click="handleClick">关闭</a-button>
  <a-modal
    v-model:visible="visible"
    title="关闭告警"
    @cancel="handleCancel"
    @ok="handleOk"
  >
    <a-form :model="form">
      <a-form-item field="is_recovered" label="是否恢复">
        <a-radio-group v-model="isRecovered" :options="options" />
      </a-form-item>
    </a-form>
  </a-modal>
</template>

<script setup lang="ts">
  import { Message } from '@arco-design/web-vue';
  import { reactive, ref } from 'vue';
  import { closeAlertEvent } from '@/api/alert-event';

  export interface Props {
    selected: (string | number)[] | null;
    type: 'dashed' | 'text' | 'outline' | 'primary' | 'secondary' | undefined;
  }

  const isRecovered = ref(1);
  const props = defineProps<Props>();
  const emits = defineEmits(['refresh']);
  const visible = ref(false);
  const options = [
    { label: '已恢复', value: 1 },
    { label: '未恢复', value: 0 },
  ];
  const form = reactive({
    isRecovered: false,
    events: [] as (number | string)[],
  });

  const handleOk = async () => {
    if (props.selected === null) {
      Message.error('请您至少选中一个要关闭的告警');
      return;
    }
    if (props.selected?.length === 0) {
      Message.error('请您至少选中一个要关闭的告警');
      return;
    }
    if (isRecovered.value === 1) {
      form.isRecovered = true;
    }
    form.events = props.selected;
    try {
      await closeAlertEvent(form);
      Message.success('关闭成功');
      form.events = [];
      emits('refresh');
      visible.value = false;
    } catch (err) {
      // you can report use errorHandler or other
    }
  };

  const handleClick = () => {
    visible.value = true;
  };

  const handleCancel = () => {
    visible.value = false;
  };
</script>
