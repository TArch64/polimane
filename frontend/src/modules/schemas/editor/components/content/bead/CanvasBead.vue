<template>
  <KonvaRect
    :config
    ref="rootRef"
    @click="beadsStore.paint(bead)"
    @mousemove="onMouseMove"
  />
</template>

<script setup lang="ts">
import Konva from 'konva';
import { computed } from 'vue';
import type { ISchemaBead, ISchemaRow } from '@/models';
import { useNodeConfigs, useNodeCursor, useNodeRef } from '@/modules/schemas/editor/composables';
import { useBeadsStore, usePaletteStore } from '@/modules/schemas/editor/stores';

const props = defineProps<{
  row: ISchemaRow;
  bead: ISchemaBead;
}>();

const paletteStore = usePaletteStore();
const beadsStore = useBeadsStore(() => props.row);

const config = useNodeConfigs<Konva.RectConfig>([
  {
    width: 14,
    height: 14,
    cornerRadius: 4,
  },
  computed(() => ({
    fill: props.bead.color || 'rgba(0, 0, 0, 0.05)',
  })),
]);

const rootRef = useNodeRef<Konva.Rect>();
useNodeCursor(rootRef, 'crosshair');

function onMouseMove() {
  if (paletteStore.isPainting) {
    beadsStore.paint(props.bead);
  }
}
</script>
