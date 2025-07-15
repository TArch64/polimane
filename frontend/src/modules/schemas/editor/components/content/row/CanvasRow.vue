<template>
  <GroupRenderer
    ref="rootRef"
    :config="rootConfig"
    @mousedown="onMouseDown"
  >
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
import { computed, ref } from 'vue';
import type { ISchemaRow } from '@/models';
import {
  type MaybeNodeConfig,
  useNodeConfigs,
  useNodeContextMenu,
  useNodeFiller,
  useNodeRef,
  useRowContextMenuActions,
  useRowTitle,
} from '@/modules/schemas/editor/composables';
import { useCursorStore, useDraggingStore } from '@/modules/schemas/editor/stores';
import { CanvasBead } from '../bead';
import { CanvasStackH, GroupRenderer } from '../base';

const props = defineProps<{
  row: ISchemaRow;
}>();

const draggingStore = useDraggingStore();
const cursorStore = useCursorStore();

const rootRef = useNodeRef<Konva.Group>();
const dragTranslation = ref(0);
const isDragging = computed(() => dragTranslation.value !== 0);

const rootConfig = computed((): Partial<Konva.GroupConfig> => ({
  offsetY: -dragTranslation.value,
}));

const contentRef = useNodeRef<Konva.Group>();

useNodeContextMenu({
  nodeRef: rootRef,
  title: useRowTitle(() => props.row),
  actions: useRowContextMenuActions(() => props.row),
});

const backgroundConfig = useNodeConfigs<Konva.RectConfig>([
  useNodeFiller(contentRef, {
    padding: computed(() => isDragging.value ? { horizontal: 4, vertical: 4 } : {}),
  }),
  {
    fill: '#F8F8F8',
    stroke: '#1A1A1A',
    strokeWidth: 0,
    cornerRadius: 2,
  },
  computed((): MaybeNodeConfig<Konva.RectConfig> => isDragging.value && {
    strokeWidth: 2,
    offsetY: 2,
    offsetX: 2,
  }),
]);

function onDragEnd(): void {
  rootRef.value.zIndex(undefined);
  rootRef.value.parent?.listening(true);
}

function onMouseDown() {
  if (cursorStore.isDragging) {
    const parent = rootRef.value.parent!;
    parent.listening(false);
    rootRef.value.zIndex(parent.children.length - 1);

    draggingStore.capture({
      object: props.row,
      dragTranslation,
      onCleanUp: onDragEnd,
    });
  }
}
</script>
