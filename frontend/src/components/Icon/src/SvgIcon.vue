<template>
  <svg
    class="svg-icon"
    :class="[$attrs.class, spin && 'svg-icon-spin']"
    :style="getStyle"
    aria-hidden="true"
  >
    <use :xlink:href="`#${icon}`"></use>
  </svg>
</template>

<script lang="ts">
  import type { CSSProperties } from 'vue';
  import { defineComponent, computed } from 'vue';

  export default defineComponent({
    name: 'SvgIcon',
    props: {
      icon: {
        type: String,
        required: true,
      },
      color: {
        type: String,
      },
      size: {
        type: [Number, String],
        default: 16,
      },
      spin: {
        type: Boolean,
        default: false,
      },
    },
    setup(props) {
      const getStyle = computed((): CSSProperties => {
        const { size, color } = props;
        let s = `${size}`;
        s = `${s.replace('px', '')}px`;
        return {
          color,
          width: s,
          height: s,
        };
      });
      return { getStyle };
    },
  });
</script>

<style lang="less" scoped>
  .svg-icon {
    display: inline-block;
    overflow: hidden;
    vertical-align: -0.15em;
    fill: currentColor;
  }
  .svg-icon-spin {
    animation: loadingCircle 1s infinite linear;
  }
</style>
