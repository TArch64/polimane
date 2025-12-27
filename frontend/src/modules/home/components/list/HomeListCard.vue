<template>
  <Card
    interactable
    ref="cardRef"
    class="home-list-card tap-animation"
    :active="selected"
    :binding="cardBinding"
  >
    <slot />
  </Card>
</template>

<script setup lang="ts">
import { type RouteLocationRaw, RouterLink } from 'vue-router';
import { computed, type Slot, toRef } from 'vue';
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

useContextMenu({
  el: cardRef,
  control: false,
  title: toRef(props, 'menuTitle'),
  actions: toRef(props, 'menuActions'),
});
</script>

<style scoped>
@layer page {
  .home-list-card {
    overflow: clip;
    box-shadow: var(--box-shadow);
    --tap-scale: 0.98;
  }
}
</style>
