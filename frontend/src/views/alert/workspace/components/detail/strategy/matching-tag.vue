<template>
  <a-card>
    <a-input-group>
      <a-input :style="{ width: '160px' }" placeholder="标签" />
      <a-select :style="{ width: '320px' }" placeholder="Please select ...">
        <a-option>包含</a-option>
        <a-option>不包含</a-option>
      </a-select>
      <a-input :style="{ width: '160px' }" placeholder="值" />
      <icon-plus-circle size="20" style="cursor: pointer; margin-left: 20px" />
    </a-input-group>
  </a-card>
</template>
