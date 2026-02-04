<template>
  <svg
    :viewBox
    class="popover__tip"
    xmlns="http://www.w3.org/2000/svg"
    preserveAspectRatio="none"
  >
    <path
      d="M0 1.5 L12 11 L24 1.5 Z"
      fill="var(--color-background-1)"
      stroke="var(--color-divider)"
      stroke-width="1"
      stroke-linejoin="round"
      stroke-linecap="round"
    />

    <line
      x1="1.5"
      y1="1.5"
      x2="22.5"
      y2="1.5"
      stroke="var(--color-background-1)"
      stroke-linecap="round"
      stroke-linejoin="round"
    />

    <line
      x1="1"
      y1="1"
      x2="23"
      y2="1"
      stroke="var(--color-background-1)"
      stroke-linecap="round"
      stroke-linejoin="round"
    />
  </svg>
</template>

<script setup lang="ts">
import { toReactive, useElementBounding } from '@vueuse/core';
import { computed } from 'vue';
import { POPOVER_TIP_HEIGHT, POPOVER_TIP_WIDTH } from './config';

const props = defineProps<{
  referenceEl: HTMLElement;
  floatingEl: HTMLElement;
}>();

const viewBox = `0 0 ${POPOVER_TIP_WIDTH} ${POPOVER_TIP_HEIGHT}`;

const referenceRect = toReactive(useElementBounding(() => props.referenceEl));
const floatingRect = toReactive(useElementBounding(() => props.floatingEl));

const left = computed(() => {
  return (referenceRect.left + referenceRect.width / 2) - floatingRect.left;
});
</script>

<style scoped>
@layer components {
  .popover__tip {
    position: absolute;
    left: v-bind("left + 'px'");
    bottom: v-bind("1 - POPOVER_TIP_HEIGHT + 'px'");
    width: v-bind("POPOVER_TIP_WIDTH + 'px'");
    height: v-bind("POPOVER_TIP_HEIGHT + 'px'");
  }
}
</style>
