<template>
  <Card
    :as="ToolbarGrid"
    ref="rootRef"
    class="color-palette"
  >
    <ColorPaletteSpot
      v-for="(color, index) of store.palette"
      :key="index"
      :active="color === store.activeColor"
      @update:active="store.activateColor(color)"
      @choose="$emit('close')"
      v-model="store.palette[index]!"
    />
  </Card>
</template>

<script setup lang="ts">
import { useToolsStore } from '@editor/stores';
import { onClickOutside } from '@vueuse/core';
import { Card } from '@/components/card';
import { useDomRef } from '@/composables';
import ToolbarGrid from '../ToolbarGrid.vue';
import ColorPaletteSpot from './ColorPaletteSpot.vue';

const emit = defineEmits<{
  close: [];
}>();

const rootRef = useDomRef<HTMLElement>();
const store = useToolsStore();

onClickOutside(rootRef, () => emit('close'));
</script>

<style scoped>
@layer page {
  .color-palette {
    padding: 6px 8px;
  }
}
</style>
