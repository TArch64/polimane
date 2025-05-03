<template>
  <KonvaRect :config ref="rootRef" />
</template>

<script setup lang="ts">
import Konva from 'konva';
import { computed } from 'vue';
import type { ISchemaRow } from '@/models';
import { useActiveObject, useNodeConfigs, useNodeRef } from '@/modules/schemas/editor/composables';
import { ActiveObjectTrigger } from '@/modules/schemas/editor/stores';
import { scrollNodeIntoView } from '@/modules/schemas/editor/helpers';

const props = defineProps<{
  row: ISchemaRow;
}>();

const activeObject = useActiveObject(() => props.row);

const rootRef = useNodeRef<Konva.Group>();

const config = useNodeConfigs<Konva.RectConfig>([
  {
    width: 800,
    height: 20,
  },
  computed(() => ({
    fill: activeObject.focus.isExactActive ? 'rgba(0, 0, 0, 0.4)' : 'rgba(0, 0, 0, 0.1)',
  })),
]);

activeObject.focus.onExactActive((trigger) => {
  if (trigger !== ActiveObjectTrigger.CANVAS) {
    scrollNodeIntoView(rootRef.value, { scale: true });
  }
});
</script>
