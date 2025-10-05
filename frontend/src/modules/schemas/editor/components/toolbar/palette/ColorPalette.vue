<template>
  <Card :as="ToolbarGrid" class="color-palette">
    <ColorPaletteSpot
      v-for="(_, index) of store.palette"
      :key="index"
      :active="activeColorIndex === index"
      @update:active="activeColorIndex = index"
      v-model="store.palette[index]!"
    />
  </Card>
</template>

<script setup lang="ts">
import { useToolsStore } from '@editor/stores';
import { Card } from '@/components/card';
import ToolbarGrid from '../ToolbarGrid.vue';
import ColorPaletteSpot from './ColorPaletteSpot.vue';

const activeColorIndex = defineModel<number>({
  required: true,
});

const store = useToolsStore();
</script>

<style scoped>
@layer page {
  .color-palette {
    padding: 6px 8px;
  }
}
</style>
