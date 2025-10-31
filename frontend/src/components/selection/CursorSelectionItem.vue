<template>
  <slot :itemRef />
</template>

<script setup lang="ts" generic="I extends SelectionItem">
import type { Slot, VNodeRef } from 'vue';
import { unrefElement } from '@vueuse/core';
import type { ISelectionItem, SelectionItem, SelectionListRegistry } from './ListRegistry';

const props = defineProps<{
  item: I;
  registry: SelectionListRegistry<I>;
}>();

defineSlots<{
  default: Slot<{ itemRef: VNodeRef }>;
}>();

const createState = (el: HTMLElement): ISelectionItem<I> => ({
  el,
  data: props.item,
});

const itemRef: VNodeRef = (ref) => {
  const itemEl = ref ? unrefElement(ref as HTMLElement) : null;

  itemEl
    ? props.registry.set(props.item.id, createState(itemEl))
    : props.registry.delete(props.item.id);
};
</script>
