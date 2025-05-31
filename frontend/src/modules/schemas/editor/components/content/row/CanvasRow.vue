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
  useNodeContextMenu,
  useNodeFiller,
  useNodeRef,
  useRowContextMenuActions,
  useRowTitle,
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

useNodeContextMenu({
  nodeRef: rootRef,
  title: useRowTitle(() => props.row),
  actions: useRowContextMenuActions(() => props.row),
});

const backgroundConfig = useNodeFiller(contentRef);

activeObject.focus.onExactActive((trigger) => {
  if (trigger !== ActiveObjectTrigger.CANVAS) {
    scrollNodeIntoView(rootRef.value, {
      scale: true,
      minWidth: 700,
    });
  }
});
</script>
