<template>
  <KonvaGroup :config ref="groupRef">
    <ChildrenRenderer>
      <slot />
    </ChildrenRenderer>
  </KonvaGroup>
</template>

<script setup lang="ts">
import { computed, type ComputedRef, type Slot } from 'vue';
import Konva from 'konva';
import { useNodeRef } from '@/modules/schemas/editor/composables';
import type { StackUpdateFn } from './StackUpdateFn';
import { createCanvasStackRenderer } from './CanvasStackRenderer';

const props = defineProps<{
  initial: number;
  update: StackUpdateFn;
  gap: number;
  config?: Partial<Konva.GroupConfig>;
}>();

defineSlots<{
  default: Slot;
}>();

const groupRef = useNodeRef<Konva.Group | null>();
const children: ComputedRef<Konva.Node[]> = computed(() => groupRef.value?.children ?? []);

let updateTweens: Konva.Tween[] = [];

const ChildrenRenderer = createCanvasStackRenderer((isInitial, keys: unknown[]): void => {
  for (const tween of updateTweens) tween.destroy();
  updateTweens = [];

  const ordered = keys.map((key) => children.value.find((child) => child.id() === key)!).filter(Boolean);
  const list = ordered.length === children.value.length ? ordered : children.value;

  let next = props.initial;

  for (const child of list) {
    const update = props.update({
      parent: groupRef.value!,
      next,
      child,
      isInitial,
    });

    if (update.tween) {
      update.tween.play();
      updateTweens.push(update.tween);
    }

    next += update.next;
  }
});

defineExpose({
  getNode: () => groupRef.value!,
});
</script>
