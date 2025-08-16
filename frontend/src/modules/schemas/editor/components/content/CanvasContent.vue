<template>
  <GroupRenderer
    :config
    ref="rootRef"
    @mousedown="paint"
  >
    <CanvasBeadGrid
      v-for="{sector, grid} of sectors"
      :key="sector"
      :grid="grid.value"
    />
  </GroupRenderer>
</template>

<script setup lang="ts">
import Konva from 'konva';
import { computed } from 'vue';
import {
  useNodeCentering,
  useNodeCursor,
  useNodeListener,
  useNodeRef,
} from '@/modules/schemas/editor/composables';
import { useEditorStore, usePaletteStore } from '@/modules/schemas/editor/stores';
import { GroupRenderer } from './base';
import { useBeadsGrid } from './useBeadsGrid';
import { CanvasBeadGrid } from './CanvasBeadGrid';

const editorStore = useEditorStore();
const paletteStore = usePaletteStore();

const sectors = useBeadsGrid();

const rootRef = useNodeRef<Konva.Group>();
const config = useNodeCentering(rootRef, { padding: 30 });
const isActive = computed(() => paletteStore.isPainting);

function paint(event: Konva.KonvaEventObject<MouseEvent>) {
  const position = event.target.getAttr('$position');

  if (!position) {
    return;
  }

  if (paletteStore.activeColor) {
    editorStore.schema.beads[position] = paletteStore.activeColor;
    return;
  }

  if (editorStore.schema.beads[position]) {
    delete editorStore.schema.beads[position];
  }
}

useNodeListener(rootRef, 'mousemove', paint, { isActive });
useNodeCursor(rootRef, 'crosshair');
</script>
