<template>
  <GroupRenderer
    :config
    @click="onClick"
    @mouseover="hoverObjectStore.activateObject(pattern)"
    @mouseout="hoverObjectStore.deactivatePath"
  >
    <KonvaRect ref="borderRef" :config="borderConfig" />
    <CanvasPatternLabel :pattern />

    <GroupRenderer
      ref="contentGroupRef"
      :config="contentGroupConfig"
      @layout="onContentLayout"
    >
      <CanvasPatternContent :pattern v-if="pattern.content.length" />
      <CanvasPatternEmpty v-else />
    </GroupRenderer>
  </GroupRenderer>
</template>

<script setup lang="ts">
import Konva from 'konva';
import { computed, ref } from 'vue';
import type { ISchemaPattern } from '@/models';
import {
  useNodeCentering,
  useNodeConfigs,
  useNodeRef,
  useNodeTween,
} from '@/modules/schemas/editor/composables';
import { useModal } from '@/components/modal';
import { useFocusObjectStore, useHoverObjectStore } from '@/modules/schemas/editor/stores';
import { getPatternAddRowModal } from '../../modals';
import { GroupRenderer, type IGroupLayoutEvent, NodeRect } from '../base';
import CanvasPatternLabel from './CanvasPatternLabel.vue';
import CanvasPatternContent from './CanvasPatternContent.vue';
import CanvasPatternEmpty from './CanvasPatternEmpty.vue';

const props = defineProps<{
  pattern: ISchemaPattern;
}>();

const focusObjectStore = useFocusObjectStore();
const isFocus = focusObjectStore.useActiveObject(() => props.pattern);

const hoverObjectStore = useHoverObjectStore();
const isHover = hoverObjectStore.useActiveObject(() => props.pattern);

const addModal = useModal(getPatternAddRowModal(props.pattern));

function onClick() {
  props.pattern.content.length
    ? focusObjectStore.activateObject(props.pattern)
    : addModal.open({ pattern: props.pattern });
}

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

const contentLayoutRect = ref<NodeRect>(NodeRect.BLANK);

function onContentLayout(event: IGroupLayoutEvent) {
  if (!contentLayoutRect.value.isEqual(event.clientRect)) {
    contentLayoutRect.value = event.clientRect;
  }
}

const borderConfig = useNodeConfigs<Konva.RectConfig>([
  {
    x: 1,
    y: 1,
    stroke: 'rgba(0, 0, 0, 0.1)',
    strokeWidth: 1,
    cornerRadius: 8,
    dash: [10, 5],
  },

  computed(() => ({
    width: Math.max(contentLayoutRect.value.width, 1000) + 24,
    height: Math.max(contentLayoutRect.value.height, 100) + 40,
  })),
]);

const borderRef = useNodeRef<Konva.Rect | null>();

const borderAnimatedConfig = computed((): Partial<Konva.RectConfig> => {
  if (isFocus.value) return { stroke: 'rgba(0, 0, 0, 0.7)' };
  if (isHover.value) return { stroke: 'rgba(0, 0, 0, 0.5)' };
  return { stroke: borderConfig.value.stroke! };
});

useNodeTween(borderRef, borderAnimatedConfig, (config) => ({
  ...config,
  duration: 0.15,
  easing: Konva.Easings.EaseOut,
}));
</script>
