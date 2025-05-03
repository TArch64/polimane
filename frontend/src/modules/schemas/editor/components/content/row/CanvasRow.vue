<template>
  <CanvasStackH ref="rootRef">
    <CanvasBead
      v-for="bead of row.content"
      :key="bead.id"
      :bead
    />
  </CanvasStackH>
</template>

<script setup lang="ts">
import Konva from 'konva';
import type { ISchemaRow } from '@/models';
import { useActiveObject, useNodeRef } from '@/modules/schemas/editor/composables';
import { ActiveObjectTrigger } from '@/modules/schemas/editor/stores';
import { scrollNodeIntoView } from '@/modules/schemas/editor/helpers';
import { CanvasBead } from '../bead';
import { CanvasStackH } from '../base';

const props = defineProps<{
  row: ISchemaRow;
}>();

const activeObject = useActiveObject(() => props.row);

const rootRef = useNodeRef<Konva.Group>();

activeObject.focus.onExactActive((trigger) => {
  if (trigger !== ActiveObjectTrigger.CANVAS) {
    scrollNodeIntoView(rootRef.value, { scale: true });
  }
});
</script>
