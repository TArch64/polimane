<template>
  <Card
    ref="cardRef"
    class="home-list-card"
    :active="selected || isContextMenuActive"
    :binding="cardBinding"
    :interactable="!disabled"
  >
    <slot />
  </Card>
</template>

<script setup lang="ts">
import { type RouteLocationRaw, RouterLink } from 'vue-router';
import { computed, ref, type Slot, toRef } from 'vue';
import { Card } from '@/components/card';
import { makeBinding } from '@/components/binding';
import { type MaybeContextMenuAction, useContextMenu } from '@/components/contextMenu';
import { useDomRef } from '@/composables';

const props = withDefaults(defineProps<{
  to: RouteLocationRaw;
  selected?: boolean;
  disabled?: boolean;
  menuTitle: string;
  menuActions: MaybeContextMenuAction[];
}>(), {
  selected: false,
  disabled: false,
});

defineSlots<{
  default: Slot;
}>();

const cardRef = useDomRef<HTMLElement>();

const linkBinding = makeBinding(RouterLink, () => ({
  draggable: false,
  to: props.to,
}));

const disabledBinding = makeBinding('div', () => ({
  draggable: false,
}));

const cardBinding = computed(() => {
  return props.disabled ? disabledBinding.value : linkBinding.value;
});

const isContextMenuActive = ref(false);

useContextMenu({
  el: cardRef,
  variant: 'main',
  title: toRef(props, 'menuTitle'),
  actions: toRef(props, 'menuActions'),
  isActive: isContextMenuActive,
});
</script>

<style scoped>
@layer page {
  .home-list-card {
    overflow: clip;
    box-shadow: var(--box-shadow);
  }
}
</style>
