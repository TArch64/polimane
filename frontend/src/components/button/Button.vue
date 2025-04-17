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
type ButtonVariant = 'primary' | 'secondary' | 'danger';

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
    --button-primary-disabled-background: color-mix(in srgb, var(--color-primary), transparent 70%);
    --button-primary-foreground: var(--color-white);

    --button-secondary-background: tranparent;
    --button-secondary-hover-background: color-mix(in srgb, var(--color-primary), transparent 90%);
    --button-secondary-disabled-background: color-mix(in srgb, var(--color-primary), transparent 70%);
    --button-secondary-foreground: var(--color-primary);

    --button-danger-background: tranparent;
    --button-danger-hover-background: color-mix(in srgb, var(--color-danger), transparent 90%);
    --button-danger-disabled-background: color-mix(in srgb, var(--color-danger), transparent 70%);
    --button-danger-foreground: var(--color-danger);
  }

  .button {
    display: flex;
    align-items: center;
    justify-content: center;
    cursor: pointer;
  }

  .button--md {
    font-size: var(--font-sm);
    line-height: 18px;
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
  .button--secondary,
  .button--danger {
    background-color: var(--button-background);
    color: var(--button-foreground);
    border-radius: var(--rounded-md);
    transition: background-color 0.15s ease-out;
    will-change: background-color;

    &:hover:not([disabled]) {
      background-color: var(--button-hover-background);
    }

    &[disabled] {
      background-color: var(--button-disabled-background);
    }
  }

  .button--primary {
    --button-background: var(--button-primary-background);
    --button-hover-background: var(--button-primary-hover-background);
    --button-disabled-background: var(--button-primary-disabled-background);
    --button-foreground: var(--button-primary-foreground);
  }

  .button--secondary {
    --button-background: var(--button-secondary-background);
    --button-hover-background: var(--button-secondary-hover-background);
    --button-disabled-background: var(--button-secondary-disabled-background);
    --button-foreground: var(--button-secondary-foreground);
  }

  .button--danger {
    --button-background: var(--button-danger-background);
    --button-hover-background: var(--button-danger-hover-background);
    --button-disabled-background: var(--button-danger-disabled-background);
    --button-foreground: var(--button-danger-foreground);
  }
}
</style>
