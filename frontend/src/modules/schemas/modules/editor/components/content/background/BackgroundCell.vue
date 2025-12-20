<template>
  <rect
    :width="BEAD_SIZE"
    :height="BEAD_SIZE"
    :fill="fillColor"
    fill-opacity="0.2"
    :stroke="strokeColor"
    stroke-width="1"
    class="background__cell"
  />
</template>

<script setup lang="ts">
import { BEAD_SIZE } from '@editor/const';
import { computed } from 'vue';
import { useEditorStore } from '@editor/stores';

const editorStore = useEditorStore();

const emptyColor = '#CFCFCF';
const backgroundColor = computed(() => editorStore.schema.backgroundColor);

const fillColor = computed(() => {
  return `color-mix(in srgb, ${backgroundColor.value}, ${emptyColor})`;
});

const strokeColor = computed(() => {
  return `color-mix(in srgb, ${backgroundColor.value}, ${emptyColor} 50%)`;
});
</script>

<style scoped>
@layer page {
  .background__cell {
    transition: fill 0.15s ease-out;
  }
}
</style>
