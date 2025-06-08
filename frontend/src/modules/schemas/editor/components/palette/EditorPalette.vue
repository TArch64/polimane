<template>
  <Card as="section" class="palette">
    <ColorList :size="store.palette.length + 1">
      <ColorItem
        v-for="(_, index) of store.palette"
        :key="index"
        :active="activeColorIndex === index"
        @update:active="activeColorIndex = index"
        v-model="store.palette[index]"
      />

      <ColorEraser />
    </ColorList>
  </Card>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { useEventListener } from '@vueuse/core';
import { type ActiveToolId, usePaletteStore } from '@/modules/schemas/editor/stores';
import { Card } from '@/components/card';
import ColorList from './ColorList.vue';
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
  .palette {
    position: absolute;
    top: 8px;
    right: 8px;
    padding: 8px;
    display: flex;
    align-items: flex-start;
    gap: 8px;
    --color-button-size: 20px;
  }
}
</style>
