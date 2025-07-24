<template>
  <ButtonRoot
    :to
    class="button"
    :class="classes"
    :disabled="disabled ? 'disabled' : undefined"
  >
    <Component
      class="button__prepend-icon"
      :is="prependIcon"
      v-if="prependIcon"
    />

    <slot />
  </ButtonRoot>
</template>

<script setup lang="ts">
import { computed, type Slot } from 'vue';
import type { RouteLocationRaw } from 'vue-router';
import type { IconComponent } from '../icon';
import ButtonRoot from './ButtonRoot.vue';
import type { ButtonSize } from './ButtonSize';
import type { ButtonVariant } from './ButtonVariant';

const props = withDefaults(defineProps<{
  to?: RouteLocationRaw;
  icon?: boolean;
  size?: ButtonSize;
  variant?: ButtonVariant;
  prependIcon?: IconComponent;
  danger?: boolean;
  disabled?: boolean;
}>(), {
  icon: false,
  danger: false,
  variant: 'secondary',
  size: 'md',
});

defineSlots<{
  default: Slot;
}>();

const classes = computed(() => [
  props.icon && 'button--icon',
  props.size && `button--${props.size}`,
  props.variant && `button--${props.variant}`,
  props.danger && 'button--danger',
]);
</script>

<style scoped>
@layer components {
  :global(:root) {
    --button-base-color: var(--color-primary);
  }

  .button {
    display: flex;
    align-items: center;
    justify-content: center;
    cursor: pointer;
  }

  .button__prepend-icon {
    margin-right: 4px;
  }

  .button--md {
    font-size: var(--font-sm);
    line-height: 18px;
    padding: 6px 12px;
    min-height: 30px;

    &.button--icon {
      padding: 6px;
    }
  }

  .button--lg {
    font-size: var(--font-md);
    padding: 8px 12px;

    &.button--icon {
      padding: 8px;
    }
  }

  .button--danger {
    --button-base-color: var(--color-danger);
  }

  .button--primary,
  .button--secondary {
    background-color: var(--button-background);
    color: var(--button-foreground);
    border-radius: var(--rounded-md);
    transition: background-color 0.15s ease-out;
    will-change: background-color;

    &:hover:not([disabled]),
    &.router-link-exact-active:not([disabled]) {
      background-color: var(--button-hover-background);
    }

    &[disabled] {
      background-color: var(--button-disabled-background);
      color: var(--button-disabled-foreground);
      cursor: default;
    }
  }

  .button--primary {
    --button-background: var(--button-base-color);
    --button-hover-background: color-mix(in srgb, var(--button-base-color), transparent 20%);
    --button-disabled-background: color-mix(in srgb, var(--button-base-color), transparent 70%);
    --button-foreground: var(--color-white);
    --button-disabled-foreground: var(--color-white);
  }

  .button--secondary {
    --button-background: transparent;
    --button-hover-background: color-mix(in srgb, var(--button-base-color), transparent 90%);
    --button-disabled-background: transparent;

    --button-foreground: var(--button-base-color);
    --button-disabled-foreground: color-mix(in srgb, var(--button-base-color), transparent 70%);
  }
}
</style>
