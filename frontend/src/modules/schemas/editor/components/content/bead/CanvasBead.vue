<template>
  <GroupRenderer ref="rootRef">
    <KonvaRect :config="backgroundConfig" />

    <KonvaRect
      ref="beadRef"
      :config="beadConfig"
      @click="beadsStore.paint(bead)"
      @mousemove="onMouseMove"
    />
  </GroupRenderer>
</template>

<script setup lang="ts">
import Konva from 'konva';
import { computed } from 'vue';
import type { ISchemaBead, ISchemaRow } from '@/models';
import {
  useNodeConfigs,
  useNodeCursor,
  useNodeFiller,
  useNodeRef,
} from '@/modules/schemas/editor/composables';
import { useBeadsStore, usePaletteStore } from '@/modules/schemas/editor/stores';
import { GroupRenderer } from '../base';

const props = defineProps<{
  row: ISchemaRow;
  bead: ISchemaBead;
}>();

const paletteStore = usePaletteStore();
const beadsStore = useBeadsStore(() => props.row);

const rootRef = useNodeRef<Konva.Rect>();
const beadRef = useNodeRef<Konva.Rect>();

useNodeCursor(rootRef, 'crosshair');

const backgroundConfig = useNodeFiller(beadRef, {
  padding: 0.5,
});

const beadConfig = useNodeConfigs<Konva.RectConfig>([
  {
    width: 14,
    height: 14,
    cornerRadius: 999,
  },
  computed(() => ({
    fill: props.bead.color || 'rgba(0, 0, 0, 0.05)',
  })),
]);

function onMouseMove() {
  if (paletteStore.isPainting) {
    beadsStore.paint(props.bead);
  }
}
</script>
