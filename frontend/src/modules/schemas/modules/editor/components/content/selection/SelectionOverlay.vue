<template>
  <div class="canvas-selection__overlay" />
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { useSelectionStore } from '@editor/stores';
import { isVerticalDirection } from '@/enums';

const selectionStore = useSelectionStore();
const direction = computed(() => selectionStore.resize.direction);

const overlayCursor = computed(() => {
  if (!direction.value) return null;
  return isVerticalDirection(direction.value) ? 'ns-resize' : 'ew-resize';
});
</script>

<style scoped>
@layer page {
  .canvas-selection__overlay {
    position: fixed;
    z-index: 999;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    cursor: v-bind("overlayCursor")
  }
}
</style>
