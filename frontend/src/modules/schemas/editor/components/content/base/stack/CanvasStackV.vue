<template>
  <CanvasStack
    :gap
    :config
    :update
    :initial="0"
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
import type { INodeRect } from '../NodeRect';
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

function getAlignValue(parent: Konva.Group, childRect: INodeRect): number {
  if (props.align === 'start') {
    return 0;
  }

  const freeSpace = parent.getClientRect().width - childRect.width;

  if (props.align === 'end') {
    return freeSpace;
  }

  return freeSpace / 2;
}

const update: StackUpdateFn<'y'> = (payload) => {
  const childRect = payload.child.getClientRect();

  return {
    next: childRect.height + props.gap,
    property: 'y',

    extra: {
      x: getAlignValue(payload.parent, childRect),
    },
  };
};

defineExpose({ getNode: () => stackRef.value! });
</script>
