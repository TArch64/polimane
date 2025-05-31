<template>
  <GroupRenderer ref="rootRef">
    <CanvasStackH ref="contentRef" :gap="4">
      <CanvasBead
        v-for="bead of row.content"
        :key="bead.id"
        :bead
      />
    </CanvasStackH>

    <KonvaRect :config="backgroundConfig" />
  </GroupRenderer>
</template>

<script setup lang="ts">
import Konva from 'konva';
import type { ISchemaRow } from '@/models';
import {
  useActiveObject,
  useNodeConfigs,
  useNodeContextMenu,
  useNodeFiller,
  useNodeRef,
  useRowContextMenuActions,
} from '@/modules/schemas/editor/composables';
import { ActiveObjectTrigger } from '@/modules/schemas/editor/stores';
import { scrollNodeIntoView } from '@/modules/schemas/editor/helpers';
import { CanvasBead } from '../bead';
import { CanvasStackH, GroupRenderer } from '../base';

const props = defineProps<{
  row: ISchemaRow;
}>();

const activeObject = useActiveObject(() => props.row);

const rootRef = useNodeRef<Konva.Group>();
const contentRef = useNodeRef<Konva.Group>();

const actions = useRowContextMenuActions(() => props.row);
useNodeContextMenu(rootRef, actions);

const backgroundConfig = useNodeConfigs<Konva.RectConfig>([
  {
    offsetY: 4,
    offsetX: 4,
  },
  useNodeFiller(contentRef, {
    padding: {
      vertical: 8,
      horizontal: 8,
    },
  }),
]);

activeObject.focus.onExactActive((trigger) => {
  if (trigger !== ActiveObjectTrigger.CANVAS) {
    scrollNodeIntoView(rootRef.value, {
      scale: true,
      minWidth: 700,
    });
  }
});
</script>
