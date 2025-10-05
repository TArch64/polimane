<template>
  <Card as="section" class="color-palette">
    <ColorPaletteSpot
      v-for="(_, index) of store.palette"
      :key="index"
      :active="activeColorIndex === index"
      @update:active="activeColorIndex = index"
      v-model="store.palette[index]!"
    />

    <ColorEraser />
  </Card>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { type ActiveToolId, usePaletteStore } from '@editor/stores';
import { Card } from '@/components/card';
import ColorEraser from './ColorEraser.vue';
import ColorPaletteSpot from './ColorPaletteSpot.vue';

const store = usePaletteStore();

const activeColorIndex = computed({
  get: (): number => store.activeToolId === 'eraser' ? -1 : store.activeToolId,
  set: (index: ActiveToolId) => store.activateTool(index),
});
</script>

<style scoped>
@layer page {
  .color-palette {
    display: grid;
    gap: 6px;
    padding: 6px 8px;
    grid-template-columns: repeat(2, var(--color-button-size));
    --color-button-size: 20px;
  }
}
</style>
