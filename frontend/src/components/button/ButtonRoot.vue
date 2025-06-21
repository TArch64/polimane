<template>
  <component :is="binding.is" v-bind="binding.props">
    <slot />
  </component>
</template>

<script setup lang="ts">
import { computed, type Slot } from 'vue';
import type { RouteLocationRaw } from 'vue-router';
import { makeBinding } from '../binding';
import RouterLink from './RouterLink.vue';

const props = defineProps<{
  to?: RouteLocationRaw;
}>();

defineSlots<{
  default: Slot;
}>();

const linkBinding = makeBinding(RouterLink, () => ({
  to: props.to!,
}));

const buttonBinding = makeBinding('button', () => ({
  type: 'button',
}));

const binding = computed(() => {
  return props.to ? linkBinding.value : buttonBinding.value;
});
</script>
