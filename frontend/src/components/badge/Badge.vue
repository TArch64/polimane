<template>
  <Component
    :is="binding.is"
    v-bind="binding.props"
    class="badge"
    :class="classes"
  >
    <slot />
  </Component>
</template>

<script setup lang="ts">
import { computed, type Slot } from 'vue';
import type { ComponentAs } from '@/types';
import type { AnyBinding } from '@/components/binding';

const props = withDefaults(defineProps<{
  as?: ComponentAs;
  binding?: AnyBinding;
  interactable?: boolean;
}>(), {
  as: 'div',
  interactable: false,

  binding: (props): AnyBinding => ({
    is: props.as ?? 'div',
    props: {},
  }),
});

defineSlots<{
  default: Slot;
}>();

const classes = computed(() => ({
  'badge--interactable': props.interactable,
}));
</script>

<style scoped>
@layer components {
  .badge {
    border-radius: var(--rounded-md);
    border: var(--divider);
    background-color: var(--color-background-3);
    padding: 2px 8px;
  }

  .badge--interactable {
    cursor: pointer;
    transition: background-color 0.2s ease-out;
    will-change: background-color;

    &:hover:not(:disabled) {
      background-color: var(--color-background-2);
    }

    &:disabled {
      color: inherit;
    }
  }
}
</style>
