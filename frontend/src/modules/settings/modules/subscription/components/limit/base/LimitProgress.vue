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
        fill="var(--color-primary)"
      />
    </svg>
  </div>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue';
import { useElementSize } from '@vueuse/core';

defineProps<{
  value: number;
}>();

const wrapperRef = ref<HTMLElement | null>(null);
const { width } = useElementSize(wrapperRef, undefined, { box: 'content-box' });
const height = 7;
const corderRadius = height / 2;
const viewBox = computed(() => `0 0 ${width.value} ${height}`);
</script>

<style scoped>
@layer page {
  .subscription-limit__progress {
    display: block;
    height: v-bind("height + 'px'");
    width: 100%;
  }
}
</style>
