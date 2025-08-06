<template>
  <GroupRenderer :config ref="rootRef" @click="onClick">
    <KonvaRect :config="borderConfig" />
    <CanvasPatternLabel :pattern />

    <GroupRenderer ref="contentGroupRef" :config="contentGroupConfig">
      <CanvasPatternContent :pattern v-if="pattern.content.length" />
      <CanvasPatternEmpty v-else />
    </GroupRenderer>
  </GroupRenderer>
</template>

<script setup lang="ts">
import Konva from 'konva';
import type { ISchemaPattern } from '@/models';
import {
  SCREENSHOT_IGNORE,
  useNodeCentering,
  useNodeConfigs,
  useNodeContextMenu,
  useNodeFiller,
  useNodeRef,
  usePatternContextMenuActions,
} from '@/modules/schemas/editor/composables';
import { useModal } from '@/components/modal';
import { useThemeVar } from '@/composables';
import { RowAddModal } from '../../modals';
import { GroupRenderer } from '../base';
import CanvasPatternLabel from './CanvasPatternLabel.vue';
import CanvasPatternContent from './CanvasPatternContent.vue';
import CanvasPatternEmpty from './CanvasPatternEmpty.vue';

const props = defineProps<{
  pattern: ISchemaPattern;
}>();

const colorDivider2 = useThemeVar('--color-divider-2');
const roundedMd = useThemeVar('--rounded-md');

const addModal = useModal(RowAddModal);

function onClick() {
  if (!props.pattern.content.length) {
    addModal.open({ pattern: props.pattern });
  }
}

const rootRef = useNodeRef<Konva.Group>();

useNodeContextMenu({
  nodeRef: rootRef,
  title: () => props.pattern.name,
  actions: usePatternContextMenuActions(() => props.pattern),
});

const config: Partial<Konva.GroupConfig> = {
  id: props.pattern.id,
};

const contentGroupRef = useNodeRef<Konva.Group>();

const contentGroupConfig = useNodeCentering(contentGroupRef, {
  padding: {
    vertical: 8,
    horizontal: 12,
  },
});

const borderConfig = useNodeConfigs<Konva.RectConfig>([
  () => ({
    name: SCREENSHOT_IGNORE,
    x: 1,
    y: 1,
    stroke: colorDivider2.value,
    strokeWidth: 1,
    cornerRadius: roundedMd.value,
    dash: [10, 5],
  }),

  useNodeFiller(contentGroupRef, {
    minSize: {
      width: 1000,
      height: 100,
    },

    padding: {
      horizontal: 12,
      vertical: 10,
    },
  }),
]);
</script>
