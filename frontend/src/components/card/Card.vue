<template>
  <Component
    :is="binding.is"
    v-bind="binding.props"
    class="card"
    :class="classes"
  >
    <slot />
  </Component>
</template>

<script setup lang="ts">
import { computed, type Slot } from 'vue';
import type { ComponentAs } from '@/types';
import type { AnyBinding } from '../binding';

const props = withDefaults(defineProps<{
  as?: ComponentAs;
  binding?: AnyBinding;
  interactable?: boolean;
  variant?: 'main' | 'control';
}>(), {
  as: 'div',
  interactable: false,
  variant: 'main',

  binding: (props): AnyBinding => ({
    is: props.as ?? 'div',
    props: {},
  }),
});

defineSlots<{
  default: Slot;
}>();

const classes = computed(() => [
  `card--variant-${props.variant}`,
  {
    'card--interactable': props.interactable,
  },
]);
</script>

<style scoped>
@layer components {
  .card {
    padding: 8px 12px;
    background-color: var(--card-background);
    border: 1px solid var(--color-divider);
    border-radius: var(--rounded-md);
  }

  .card--variant-main {
    --card-background: var(--color-white);
  }

  .card--variant-control {
    --card-background: var(--color-background-2);
  }

  .card--interactable {
    transition: border-color 0.15s ease-out;

    &:hover,
    &:focus,
    &:focus-within {
      border-color: color-mix(in srgb, var(--color-primary), transparent 80%);
    }
  }
}
</style>
