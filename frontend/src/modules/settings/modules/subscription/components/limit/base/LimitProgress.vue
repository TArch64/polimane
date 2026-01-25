<template>
  <div ref="wrapperRef">
    <svg :viewBox class="subscription-limit__progress">
      <rect
        width="100%"
        height="100%"
        :rx="corderRadius"
        :ry="corderRadius"
        fill="var(--color-background-2)"
      />

      <rect
        :width="`${value}%`"
        height="100%"
        :rx="corderRadius"
        :ry="corderRadius"
        :fill="color"
        class="subscription-limit__progress-value"
      />
    </svg>
  </div>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue';
import { useElementSize } from '@vueuse/core';

const props = defineProps<{
  value: number;
}>();

const wrapperRef = ref<HTMLElement | null>(null);

const { width } = useElementSize(wrapperRef, undefined, { box: 'content-box' });

const height = 7;
const corderRadius = height / 2;

const viewBox = computed(() => `0 0 ${width.value} ${height}`);

const color = computed(() => {
  if (props.value < 50) {
    return 'var(--color-primary)';
  }
  if (props.value < 100) {
    return 'var(--color-warning)';
  }
  return 'var(--color-danger)';
});
</script>

<style scoped>
@layer page {
  .subscription-limit__progress {
    display: block;
    height: v-bind("height + 'px'");
    width: 100%;
  }

  .subscription-limit__progress-value {
    transition: width 0.3s ease-out, fill 0.15s ease-out;
    will-change: width, fill;
  }
}
</style>
