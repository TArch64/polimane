<template>
  <GroupRenderer ref="rootRef">
    <KonvaRect :config="backgroundConfig" />

    <CanvasStackH ref="contentRef" :gap="1">
      <CanvasBead
        v-for="bead of row.content"
        :key="bead.id"
        :row
        :bead
      />
    </CanvasStackH>
  </GroupRenderer>
</template>

<script setup lang="ts">
import Konva from 'konva';
import type { ISchemaRow } from '@/models';
import {
  useNodeContextMenu,
  useNodeFiller,
  useNodeRef,
  useRowContextMenuActions,
  useRowTitle,
} from '@/modules/schemas/editor/composables';
import { CanvasBead } from '../bead';
import { CanvasStackH, GroupRenderer } from '../base';

const props = defineProps<{
  row: ISchemaRow;
}>();

const rootRef = useNodeRef<Konva.Group>();
const contentRef = useNodeRef<Konva.Group>();

useNodeContextMenu({
  nodeRef: rootRef,
  title: useRowTitle(() => props.row),
  actions: useRowContextMenuActions(() => props.row),
});

const backgroundConfig = useNodeFiller(contentRef);
</script>
