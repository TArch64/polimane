<template>
  <CanvasStack
    :gap
    :config
    :update
    :initial="0"
    ref="stackRef"
    v-slot="ctx"
  >
    <slot v-bind="ctx" />
  </CanvasStack>
</template>

<script setup lang="ts">
import type { Slot } from 'vue';
import Konva from 'konva';
import { useNodeRef } from '@/modules/schemas/editor/composables';
import type { INodeRect } from '../INodeRect';
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

defineSlots<{
  default: Slot;
}>();

const stackRef = useNodeRef<Konva.Group>();

function getAlignValue(parent: Konva.Group, childRect: INodeRect): number {
  if (props.align === 'start') {
    return 0;
  }

  const freeSpace = parent.getClientRect().height - childRect.height;

  if (props.align === 'end') {
    return freeSpace;
  }

  return freeSpace / 2;
}

const update: StackUpdateFn = (payload) => {
  const childRect = payload.child.getClientRect();

  payload.child.y(getAlignValue(payload.parent, childRect));
  let tween: Konva.Tween | undefined;

  if (payload.isInitial) {
    payload.child.x(payload.next);
  } else {
    tween = new Konva.Tween({
      node: payload.child,
      duration: 0.15,
      x: payload.next,
    });
  }

  return { next: childRect.width + props.gap, tween };
};

defineExpose({ getNode: () => stackRef.value! });
</script>
