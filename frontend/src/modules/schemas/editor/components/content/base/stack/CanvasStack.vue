<template>
  <GroupRenderer :config ref="groupRef" @layout="onLayout">
    <slot />
  </GroupRenderer>
</template>

<script setup lang="ts">
import { nextTick, type Slot } from 'vue';
import Konva from 'konva';
import { createWaiter } from '@/helpers';
import { useNodeRef } from '@/modules/schemas/editor/composables';
import { getClientRect } from '@/modules/schemas/editor/helpers';
import { GroupRenderer, type IGroupLayoutEvent } from '../GroupRenderer';
import type { StackUpdateFn } from './StackUpdateFn';

const props = defineProps<{
  update: StackUpdateFn<keyof Konva.NodeConfig>;
  gap: number;
  config?: Partial<Konva.GroupConfig>;
}>();

const emit = defineEmits<{
  layout: [event: IGroupLayoutEvent];
}>();

defineSlots<{
  default: Slot;
}>();

const groupRef = useNodeRef<Konva.Group | null>();
const rendered: Record<number, boolean> = {};

async function onLayout(event: IGroupLayoutEvent): Promise<void> {
  const waiter = createWaiter();
  const nodeRects = event.nodes.map(getClientRect);
  const totalGap = (event.nodes.length - 1) * props.gap;
  let next = 0;

  const parentRect = getClientRect(groupRef.value!).with({
    width: Math.max(...nodeRects.map((rect) => rect.width)) + totalGap,
    height: Math.max(...nodeRects.map((rect) => rect.height)) + totalGap,
  });

  for (const [index, child] of event.nodes.entries()) {
    const update = props.update({
      next,
      parentRect,
      childRect: nodeRects[index]!,
    });

    if (update.extra) {
      child.setAttrs(update.extra);
    }

    if (!rendered[child._id]) {
      rendered[child._id] = true;
      child.setAttr(update.property, next);
    } else {
      if (groupRef.value!.listening()) {
        groupRef.value!.listening(false);
      }

      const release = waiter.add();

      child.to({
        duration: 0.15,
        easing: Konva.Easings.EaseOut,
        [update.property]: next,
        onFinish: () => release(),
      });
    }

    next += update.next + props.gap;
  }

  if (!groupRef.value!.listening()) {
    await waiter.wait();
    await nextTick();
    groupRef.value!.listening(true);
  }

  emit('layout', event);
}

defineExpose({
  getNode: () => groupRef.value!,
});
</script>
