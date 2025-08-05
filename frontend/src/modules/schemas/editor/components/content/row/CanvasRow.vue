<template>
  <CanvasStackH ref="rootRef" :config="rootConfig">
    <CanvasBead
      v-for="bead of row.content"
      :key="bead.id"
      :row
      :bead
    />
  </CanvasStackH>
</template>

<script setup lang="ts">
import Konva from 'konva';
import { computed } from 'vue';
import type { ISchemaRow } from '@/models';
import {
  useNodeContextMenu,
  useNodeRef,
  useRowContextMenuActions,
  useRowTitle,
} from '@/modules/schemas/editor/composables';
import { CanvasBead } from '../bead';
import { CanvasStackH } from '../base';

const props = defineProps<{
  row: ISchemaRow;
}>();

const rootRef = useNodeRef<Konva.Group>();
const rootConfig = computed((): Partial<Konva.GroupConfig> => ({ id: props.row.id }));

useNodeContextMenu({
  nodeRef: rootRef,
  title: useRowTitle(() => props.row),
  actions: useRowContextMenuActions(() => props.row),
});
</script>
