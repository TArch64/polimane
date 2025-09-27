<template>
  <KonvaGroup
    :config
    ref="rootRef"
    @mousedown="paint"
  >
    <KonvaRect :config="backgroundConfig" />

    <CanvasSector
      v-for="{ sector, grid } of sectors"
      :key="sector"
      :grid="grid.value"
    />
  </KonvaGroup>
</template>

<script setup lang="ts">
import Konva from 'konva';
import { computed } from 'vue';
import { useThemeVar } from '@/composables';
import {
  BEAD_SIZE,
  useBeadsGrid,
  useNodeCursor,
  useNodeListener,
  useNodeRef,
} from '../../composables';
import { useBeadsStore, useEditorStore, usePaletteStore } from '../../stores';
import CanvasSector from './CanvasSector.vue';

const props = defineProps<{
  stageConfig: Required<Pick<Konva.StageConfig, 'width' | 'height'>>;
}>();

const editorStore = useEditorStore();
const paletteStore = usePaletteStore();
const beadsStore = useBeadsStore();

const colorBackground2 = useThemeVar('--color-background-2');

const { sectors, gridSize } = useBeadsGrid(() => editorStore.schema);

const rootRef = useNodeRef<Konva.Group>();

function calcContentY(): number {
  const contentHeight = (editorStore.schema.size.top + editorStore.schema.size.bottom) * BEAD_SIZE;
  const stageHeight = props.stageConfig.height;
  return (stageHeight - contentHeight) / 2;
}

function calcContentX(): number {
  const contentWidth = (editorStore.schema.size.left + editorStore.schema.size.right) * BEAD_SIZE;
  const stageWidth = props.stageConfig.width;
  return (stageWidth - contentWidth) / 2;
}

const config: Partial<Konva.GroupConfig> = {
  y: calcContentY(),
  x: calcContentX(),
};

const backgroundConfig: Partial<Konva.RectConfig> = computed(() => ({
  x: gridSize.minX,
  y: gridSize.minY,
  width: gridSize.width,
  height: gridSize.height,
  fill: colorBackground2.value,
}));

const isActive = computed(() => paletteStore.isPainting);

function paint(event: Konva.KonvaEventObject<MouseEvent>) {
  const position = event.target.getAttr('$position');
  if (position) beadsStore.paint(position);
}

useNodeListener(rootRef, 'mousemove', paint, { isActive });
useNodeCursor(rootRef, 'crosshair');
</script>
