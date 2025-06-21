<template>
  <CanvasStack
    :gap
    :config
    :update
    ref="stackRef"
    @layout="$emit('layout', $event)"
    v-slot="ctx"
  >
    <slot v-bind="ctx" />
  </CanvasStack>
</template>

<script setup lang="ts">
import type { Slot } from 'vue';
import Konva from 'konva';
import { useNodeRef } from '@/modules/schemas/editor/composables';
import { NodeRect } from '@/models';
import type { IGroupLayoutEvent } from '../GroupRenderer';
import CanvasStack from './CanvasStack.vue';
import type { StackAlignment } from './StackAlignment';
import type { StackUpdateFn } from './StackUpdateFn';

const props = withDefaults(defineProps<{
  gap?: number;
  config?: Partial<Konva.GroupConfig>;
  align?: StackAlignment;
}>(), {
  gap: 0,
  align: 'start',
});

defineEmits<{
  layout: [event: IGroupLayoutEvent];
}>();

defineSlots<{
  default: Slot;
}>();

const stackRef = useNodeRef<Konva.Group>();

function getAlignValue(parentRect: NodeRect, childRect: NodeRect): number {
  if (props.align === 'start') {
    return 0;
  }

  const freeSpace = parentRect.height - childRect.height;

  if (props.align === 'end') {
    return freeSpace;
  }

  return freeSpace / 2;
}

const update: StackUpdateFn<'x'> = (payload) => ({
  next: payload.childRect.width + props.gap,
  property: 'x',

  extra: {
    y: getAlignValue(payload.parentRect, payload.childRect),
  },
});

defineExpose({ getNode: () => stackRef.value! });
</script>
