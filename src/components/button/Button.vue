<template>
  <ButtonRoot
    :to
    class="button"
    :class="classes"
  >
    <slot />
  </ButtonRoot>
</template>

<script setup lang="ts">
import ButtonRoot from './ButtonRoot.vue';
import { computed, type Slot } from 'vue';
import type { RouteLocationRaw } from 'vue-router';

type ButtonSize = 'md' | 'lg';
type ButtonVariant = 'primary';

const props = defineProps<{
  to?: RouteLocationRaw;
  size?: ButtonSize;
  variant?: ButtonVariant;
}>();

defineSlots<{
  default: Slot;
}>();

const classes = computed(() => [
  props.size && `button--${props.size}`,
  props.variant && `button--${props.variant}`,
]);
</script>

<style scoped>
@layer components {
  :global(:root) {
    --button-primary-background: var(--color-primary);
    --button-primary-foreground: var(--color-white);
  }

  .button {
    display: flex;
    align-items: center;
    justify-content: center;
    cursor: pointer;

    &:where(a) {
      color: inherit;
      text-decoration: none;
    }

    &:where(button) {
      border: none;
      background: none;
      padding: 0;
    }
  }

  .button--md {
    font-size: var(--font-sm);
    padding: 6px 12px;
  }

  .button--lg {
    font-size: var(--font-md);
    padding: 8px 12px;
  }

  .button--primary {
    background-color: var(--button-primary-background);
    color: var(--button-primary-foreground);
    border-radius: var(--rounded-md);
  }
}
</style>
