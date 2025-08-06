<template>
  <GroupRenderer
    ref="rootRef"
    @click="beadsStore.paint(bead)"
    @mouseenter="tryPaint"
    @mouseleave="tryPaint"
    @mousemove="tryPaint"
  >
    <KonvaRect :config="backgroundConfig" />
    <KonvaRect ref="beadRef" :config="beadConfig" />
  </GroupRenderer>
</template>

<script setup lang="ts">
import Konva from 'konva';
import { computed } from 'vue';
import type { ISchemaBead, SchemaRow } from '@/models';
import {
  useNodeConfigs,
  useNodeCursor,
  useNodeFiller,
  useNodeRef,
} from '@/modules/schemas/editor/composables';
import { useBeadsStore, usePaletteStore } from '@/modules/schemas/editor/stores';
import { useThemeVar } from '@/composables';
import { GroupRenderer } from '../base';

const props = defineProps<{
  row: SchemaRow;
  bead: ISchemaBead;
}>();

const colorBackground2 = useThemeVar('--color-background-2');
const colorBackground3 = useThemeVar('--color-background-3');
const roundedFull = useThemeVar('--rounded-full');

const paletteStore = usePaletteStore();
const beadsStore = useBeadsStore(() => props.row);

const rootRef = useNodeRef<Konva.Rect>();
const beadRef = useNodeRef<Konva.Rect>();

useNodeCursor(rootRef, 'crosshair');

const backgroundConfig = useNodeConfigs<Konva.RectConfig>([
  () => ({
    fill: colorBackground2.value,
  }),

  useNodeFiller(beadRef, {
    padding: 0.5,
  }),
]);

const beadConfig = useNodeConfigs<Konva.RectConfig>([
  () => ({
    width: 14,
    height: 14,
    cornerRadius: roundedFull.value,
  }),

  computed(() => ({
    fill: props.bead.color || colorBackground3.value,
  })),
]);

function tryPaint() {
  if (paletteStore.isPainting) {
    beadsStore.paint(props.bead);
  }
}
</script>
