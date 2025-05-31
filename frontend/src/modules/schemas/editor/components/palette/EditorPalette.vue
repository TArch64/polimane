<template>
  <Card as="section" class="palette">
    <ColorList v-model="store.palette" />
  </Card>
</template>

<script setup lang="ts">
import { usePaletteStore } from '@/modules/schemas/editor/stores';
import { Card } from '@/components/card';
import ColorList from './ColorList.vue';

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
    gap: 16px;
  }
}
</style>
