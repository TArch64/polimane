<template>
  <Card as="section" class="palette">
    <ColorEraser />
    <div class="palette__divider" />
    <ColorList v-model="store.palette" />
  </Card>
</template>

<script setup lang="ts">
import { usePaletteStore } from '@/modules/schemas/editor/stores';
import { Card } from '@/components/card';
import ColorList from './ColorList.vue';
import ColorEraser from './ColorEraser.vue';

const store = usePaletteStore();
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

  .palette__divider {
    align-self: stretch;
    border-left: var(--divider);
  }
}
</style>
