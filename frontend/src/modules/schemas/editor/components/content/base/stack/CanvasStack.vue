<template>
  <GroupRenderer :config ref="groupRef" @layout="onLayout">
    <slot />
  </GroupRenderer>
</template>

<script setup lang="ts">
import { nextTick, type Slot } from 'vue';
import Konva from 'konva';
import { useNodeRef } from '@/modules/schemas/editor/composables';
import { createWaiter } from '@/helpers';
import { GroupRenderer, type IGroupLayoutEvent } from '../GroupRenderer';
import type { StackUpdateFn } from './StackUpdateFn';

const props = defineProps<{
  initial: number;
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
  let next = props.initial;

  for (const child of event.nodes) {
    const update = props.update({
      parent: groupRef.value!,
      next,
      child,
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

    next += update.next;
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
