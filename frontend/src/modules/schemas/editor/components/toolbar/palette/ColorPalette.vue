<template>
  <Card
    :binding
    ref="rootRef"
    class="color-palette"
  >
    <ColorPaletteSpot
      v-for="(color, index) of store.palette"
      :key="index"
      :active="color === store.activeColor"
      :model-value="color"
      @update:active="store.activateColor(color)"
      @update:model-value="updateColor(index, $event)"
    />
  </Card>
</template>

<script setup lang="ts">
import { useToolsStore } from '@editor/stores';
import { Card } from '@/components/card';
import { onBackdropClick, useDomRef } from '@/composables';
import { makeBinding } from '@/components/binding';
import ToolbarGrid from '../ToolbarGrid.vue';
import ColorPaletteSpot from './ColorPaletteSpot.vue';

const emit = defineEmits<{
  close: [];
}>();

const binding = makeBinding(ToolbarGrid, () => ({
  as: 'dialog',
}));

const rootRef = useDomRef<HTMLDialogElement>();
const store = useToolsStore();

function updateColor(index: number, color: string): void {
  const isActive = store.activeColor === store.palette[index];
  store.palette[index] = color;
  if (isActive) store.activateColor(color);
}

onBackdropClick(rootRef, () => emit('close'));
</script>

<style scoped>
@layer page {
  .color-palette {
    padding: 8px 6px;

    &::backdrop {
      background: transparent;
    }
  }
}
</style>
