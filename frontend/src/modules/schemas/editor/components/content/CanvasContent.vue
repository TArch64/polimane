<template>
  <GroupRenderer :config ref="rootRef">
    <CanvasBeadGrid
      v-for="{sector, grid} of sectors"
      :key="sector"
      :grid="grid.value"
    />
  </GroupRenderer>
</template>

<script setup lang="ts">
import Konva from 'konva';
import { useNodeCentering, useNodeRef } from '@/modules/schemas/editor/composables';
import { GroupRenderer } from './base';
import { useBeadsGrid } from './useBeadsGrid';
import { CanvasBeadGrid } from './CanvasBeadGrid';

const sectors = useBeadsGrid();

const rootRef = useNodeRef<Konva.Group>();
const config = useNodeCentering(rootRef, { padding: 30 });
</script>
