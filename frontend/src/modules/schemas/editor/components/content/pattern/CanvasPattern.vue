<template>
  <GroupRenderer
    :config
    ref="rootRef"
    @click="onClick"
    @mouseover="activeObject.hover.activate(ActiveObjectTrigger.CANVAS)"
    @mouseout="activeObject.hover.deactivate"
  >
    <KonvaRect ref="borderRef" :config="borderConfig" />
    <CanvasPatternLabel :pattern />

    <GroupRenderer ref="contentGroupRef" :config="contentGroupConfig">
      <CanvasPatternContent :pattern v-if="pattern.content.length" />
      <CanvasPatternEmpty v-else />
    </GroupRenderer>
  </GroupRenderer>
</template>

<script setup lang="ts">
import Konva from 'konva';
import { computed } from 'vue';
import type { ISchemaPattern } from '@/models';
import {
  useActiveObject,
  useNodeCentering,
  useNodeConfigs,
  useNodeContextMenu,
  useNodeFiller,
  useNodeRef,
  useNodeTween,
  usePatternContextMenuActions,
} from '@/modules/schemas/editor/composables';
import { useModal } from '@/components/modal';
import { ActiveObjectTrigger } from '@/modules/schemas/editor/stores';
import { scrollNodeIntoView } from '@/modules/schemas/editor/helpers';
import { getPatternAddRowModal } from '../../modals';
import { GroupRenderer } from '../base';
import CanvasPatternLabel from './CanvasPatternLabel.vue';
import CanvasPatternContent from './CanvasPatternContent.vue';
import CanvasPatternEmpty from './CanvasPatternEmpty.vue';

const props = defineProps<{
  pattern: ISchemaPattern;
}>();

const activeObject = useActiveObject(() => props.pattern);
const addModal = useModal(getPatternAddRowModal(props.pattern));

function onClick() {
  props.pattern.content.length
    ? activeObject.focus.activate(ActiveObjectTrigger.CANVAS)
    : addModal.open({ pattern: props.pattern });
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
  {
    x: 1,
    y: 1,
    stroke: 'rgba(0, 0, 0, 0.1)',
    strokeWidth: 1,
    cornerRadius: 8,
    dash: [10, 5],
  },

  useNodeFiller(contentGroupRef, {
    minSize: {
      width: 1000,
      height: 100,
    },

    padding: {
      horizontal: 24,
      vertical: 40,
    },
  }),
]);

const borderRef = useNodeRef<Konva.Rect | null>();

const borderAnimatedConfig = computed((): Partial<Konva.RectConfig> => {
  if (activeObject.focus.isActive) return { stroke: 'rgba(0, 0, 0, 0.7)' };
  if (activeObject.hover.isActive) return { stroke: 'rgba(0, 0, 0, 0.5)' };
  return { stroke: borderConfig.value.stroke! };
});

useNodeTween(borderRef, borderAnimatedConfig, (config) => ({
  ...config,
  duration: 0.15,
  easing: Konva.Easings.EaseOut,
}));

activeObject.focus.onExactActive((trigger) => {
  if (trigger !== ActiveObjectTrigger.CANVAS) {
    scrollNodeIntoView(rootRef.value, { scale: true });
  }
});
</script>
