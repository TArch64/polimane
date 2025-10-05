<template>
  <section class="color-list">
    <ColorItem
      v-for="(_, index) of store.palette"
      :key="index"
      :active="activeColorIndex === index"
      @update:active="activeColorIndex = index"
      v-model="store.palette[index]!"
    />

    <ColorEraser />
  </section>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { useEventListener } from '@vueuse/core';
import { type ActiveToolId, usePaletteStore } from '@editor/stores';
import ColorEraser from './ColorEraser.vue';
import ColorItem from './ColorItem.vue';

const store = usePaletteStore();

const activeColorIndex = computed({
  get: (): number => store.activeToolId === 'eraser' ? -1 : store.activeToolId,
  set: (index: ActiveToolId) => store.activateTool(index),
});

useEventListener('keydown', (event) => {
  if (!event.metaKey) return;
  if (!event.code.startsWith('Digit')) return;

  event.preventDefault();
  const index = Number(event.code.replace('Digit', ''));
  activeColorIndex.value = index === 0 ? 'eraser' : index - 1;
});
</script>

<style scoped>
@layer page {
  .color-list {
    display: grid;
    gap: 6px;
    grid-template-columns: repeat(2, var(--color-button-size));
    --color-button-size: 20px;
  }
}
</style>
