<template>
  <KonvaGroup
    :config
    ref="rootRef"
    @mousedown="paint"
  >
    <CanvasBeadGrid
      v-for="{sector, grid} of sectors"
      :key="sector"
      :grid="grid.value"
    />
  </KonvaGroup>
</template>

<script setup lang="ts">
import Konva from 'konva';
import { computed } from 'vue';
import { useNodeCursor, useNodeListener, useNodeRef } from '@/modules/schemas/editor/composables';
import { useEditorStore, usePaletteStore } from '@/modules/schemas/editor/stores';
import { BEAD_SIZE, useBeadsGrid } from './useBeadsGrid';
import { CanvasBeadGrid } from './CanvasBeadGrid';

const props = defineProps<{
  stageConfig: Required<Pick<Konva.StageConfig, 'width' | 'height'>>;
}>();

const editorStore = useEditorStore();
const paletteStore = usePaletteStore();

const sectors = useBeadsGrid();

const rootRef = useNodeRef<Konva.Group>();

function calcContentY(): number {
  const contentHeight = (editorStore.schema.size.top + editorStore.schema.size.bottom) * BEAD_SIZE;
  const stageHeight = props.stageConfig.height;
  return (stageHeight - contentHeight) / 2;
}

const config: Partial<Konva.GroupConfig> = {
  y: calcContentY(),
};

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
