<template>
  <GroupRenderer
    :config
    @click="openAddRowModal"
    v-on="borderHover.listeners"
  >
    <KonvaRect ref="borderRef" :config="borderConfig" />
    <CanvasPatternLabel :pattern />

    <GroupRenderer ref="contentGroupRef" :config="contentGroupConfig" @layout="onContentLayout">
      <CanvasPatternContent :pattern v-if="pattern.content.length" />
      <CanvasPatternEmpty v-else />
    </GroupRenderer>
  </GroupRenderer>
</template>

<script setup lang="ts">
import Konva from 'konva';
import { reactive } from 'vue';
import type { ISchemaPattern } from '@/models';
import {
  useNodeCentering,
  useNodeHover,
  useNodeRef,
  useNodeTween,
} from '@/modules/schemas/editor/composables';
import { useModal } from '@/components/modal';
import { getPatternAddRowModal } from '../../modals';
import { GroupRenderer, type IGroupLayoutEvent } from '../base';
import CanvasPatternLabel from './CanvasPatternLabel.vue';
import CanvasPatternContent from './CanvasPatternContent.vue';
import CanvasPatternEmpty from './CanvasPatternEmpty.vue';

const props = defineProps<{
  pattern: ISchemaPattern;
}>();

const addModal = useModal(getPatternAddRowModal(props.pattern));
const openAddRowModal = () => addModal.open({ pattern: props.pattern });

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

const borderConfig: Partial<Konva.RectConfig> = reactive({
  x: 1,
  y: 1,
  stroke: 'rgba(0, 0, 0, 0.1)',
  strokeWidth: 1,
  cornerRadius: 8,
  dash: [10, 5],
});

const borderRef = useNodeRef<Konva.Rect | null>();
const borderHover = useNodeHover();

useNodeTween(borderRef, () => borderHover.isHovered, (isHovered) => ({
  duration: 0.15,
  stroke: isHovered ? 'rgba(0, 0, 0, 0.5)' : borderConfig.stroke!,
  easing: Konva.Easings.EaseOut,
}));

function onContentLayout(event: IGroupLayoutEvent) {
  borderConfig.width = Math.max(event.clientRect.width, 1000) + 24;
  borderConfig.height = Math.max(event.clientRect.height, 100) + 40;
}
</script>
