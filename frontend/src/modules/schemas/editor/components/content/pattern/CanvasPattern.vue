<template>
  <KonvaGroup
    :config
    @click="openAddRowModal"
    v-on="borderHover.listeners"
  >
    <KonvaRect ref="borderRef" :config="borderConfig" />
    <CanvasPatternLabel :pattern />

    <KonvaGroup ref="contentGroupRef">
      <CanvasPatternContent :pattern v-if="pattern.content.length" />
      <CanvasPatternEmpty v-else />
    </KonvaGroup>
  </KonvaGroup>
</template>

<script setup lang="ts">
import Konva from 'konva';
import { computed, type ComputedRef, onMounted, reactive, ref } from 'vue';
import type { KonvaGroup, KonvaRect } from 'vue-konva';
import type { ISchemaPattern } from '@/models';
import { useNodeHover, useNodeTween } from '@/modules/schemas/editor/composables';
import { useModal } from '@/components/modal';
import { getPatternAddRowModal } from '../../modals';
import CanvasPatternLabel from './CanvasPatternLabel.vue';
import CanvasPatternContent from './CanvasPatternContent.vue';
import CanvasPatternEmpty from './CanvasPatternEmpty.vue';

const props = defineProps<{
  pattern: ISchemaPattern;
}>();

const addModal = useModal(getPatternAddRowModal(props.pattern));
const openAddRowModal = () => addModal.open({ pattern: props.pattern });

const borderRef = ref<InstanceType<KonvaRect> | null>(null);
const borderNode: ComputedRef<Konva.Rect | null> = computed(() => borderRef.value?.getNode());

const contentGroupRef = ref<InstanceType<KonvaGroup>>(null!);
const contentGroupNode: ComputedRef<Konva.Group> = computed(() => contentGroupRef.value.getNode());

const config: Partial<Konva.GroupConfig> = {
  id: props.pattern.id,
};

const borderConfig: Partial<Konva.RectConfig> = reactive({
  x: 1,
  y: 1,
  stroke: 'rgba(0, 0, 0, 0.1)',
  strokeWidth: 1,
  cornerRadius: 8,
  dash: [10, 5],
});

const borderHover = useNodeHover();

useNodeTween(borderNode, () => borderHover.isHovered, (isHovered) => ({
  duration: 0.15,
  stroke: isHovered ? 'rgba(0, 0, 0, 0.5)' : borderConfig.stroke!,
  easing: Konva.Easings.EaseOut,
}));

onMounted(() => {
  const contentRect = contentGroupNode.value.getClientRect();
  borderConfig.width = Math.max((contentRect.width + 24), 1000);
  borderConfig.height = Math.max((contentRect.height + 16), 100);
});
</script>
