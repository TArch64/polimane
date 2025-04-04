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
import { computed, type Slot } from 'vue';
import type { RouteLocationRaw } from 'vue-router';
import ButtonRoot from './ButtonRoot.vue';

type ButtonSize = 'md' | 'lg';
type ButtonVariant = 'primary' | 'secondary';

const props = withDefaults(defineProps<{
  to?: RouteLocationRaw;
  icon?: boolean;
  size?: ButtonSize;
  variant?: ButtonVariant;
}>(), {
  icon: false,
});

defineSlots<{
  default: Slot;
}>();

const classes = computed(() => [
  props.icon && 'button--icon',
  props.size && `button--${props.size}`,
  props.variant && `button--${props.variant}`,
]);
</script>

<style scoped>
@layer components {
  :global(:root) {
    --button-primary-background: var(--color-primary);
    --button-primary-hover-background: color-mix(in srgb, var(--color-primary), transparent 20%);
    --button-primary-foreground: var(--color-white);

    --button-secondary-background: var(--color-white);
    --button-secondary-hover-background: color-mix(in srgb, var(--color-primary), transparent 90%);
    --button-secondary-foreground: var(--color-primary);
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

    &.button--icon {
      padding: 4px;
    }
  }

  .button--lg {
    font-size: var(--font-md);
    padding: 8px 12px;

    &.button--icon {
      padding: 8px;
    }
  }

  .button--primary,
  .button--secondary {
    background-color: var(--button-background);
    color: var(--button-foreground);
    border-radius: var(--rounded-md);
    transition: background-color 0.15s ease-out;
    will-change: background-color;

    &:hover {
      background-color: var(--button-hover-background);
    }
  }

  .button--primary {
    --button-background: var(--button-primary-background);
    --button-hover-background: var(--button-primary-hover-background);
    --button-foreground: var(--button-primary-foreground);
  }

  .button--secondary {
    --button-background: var(--button-secondary-background);
    --button-hover-background: var(--button-secondary-hover-background);
    --button-foreground: var(--button-secondary-foreground);
  }
}
</style>
