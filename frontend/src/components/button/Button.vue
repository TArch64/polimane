<template>
  <ButtonRoot
    :to
    class="button tap-animation"
    :class="classes"
    :disabled="disabled || loading ? 'disabled' : undefined"
  >
    <Component
      :is="prependIcon"
      class="button__prepend-icon"
      v-if="prependIcon"
    />

    <Spinner class="button__loading-icon" v-if="loading" />

    <slot v-if="icon" />

    <span class="button__text" :class="textClasses" v-else>
      <slot />
    </span>
  </ButtonRoot>
</template>

<script setup lang="ts">
import { computed, type Slot } from 'vue';
import type { RouteLocationRaw } from 'vue-router';
import type { IconComponent } from '../icon';
import Spinner from '../Spinner.vue';
import ButtonRoot from './ButtonRoot.vue';
import type { ButtonSize } from './ButtonSize';
import type { ButtonVariant } from './ButtonVariant';

const props = withDefaults(defineProps<{
  to?: RouteLocationRaw;
  icon?: boolean;
  size?: ButtonSize;
  variant?: ButtonVariant;
  prependIcon?: IconComponent;
  mobileIconOnly?: boolean;
  danger?: boolean;
  disabled?: boolean;
  loading?: boolean;
  active?: boolean;
  truncate?: boolean;
}>(), {
  icon: false,
  danger: false,
  variant: 'secondary',
  size: 'md',
  mobileIconOnly: false,
});

defineSlots<{
  default: Slot;
}>();

const classes = computed(() => [
  props.size && props.variant !== 'inline' && `button--${props.size}`,
  props.variant && `button--${props.variant}`,
  {
    'button--icon': props.icon,
    'button--danger': props.danger,
    'button--loading': props.loading,
    'button--active': props.active,
    'button--icon-only': props.mobileIconOnly,
  },
]);

const textClasses = computed(() => ({
  'text-truncate': props.truncate,
}));
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
    position: relative;
    --tap-scale: 0.98;
  }

  .button__prepend-icon {
    margin-left: -4px;
    margin-right: 4px;
    flex-shrink: 0;
  }

  .button--icon {
    --tap-scale: 0.9;
  }

  .button--sm {
    min-height: 24px;

    &.button--icon {
      padding: 4px;
    }
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
    transition: background-color 0.15s ease-out, color 0.15s ease-out;
    will-change: background-color, color;

    &:where(:hover, .router-link-exact-active, .button--active):not([disabled]) {
      color: var(--button-hover-foreground, var(--button-foreground));
      background-color: var(--button-hover-background);
    }

    &[disabled]:not(.button--loading) {
      background-color: var(--button-disabled-background);
      color: var(--button-disabled-foreground);
      cursor: default;
    }

    &.button--loading {
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
    --button-hover-foreground: var(--button-base-color);
    --button-disabled-foreground: color-mix(in srgb, var(--button-base-color), transparent 70%);
  }

  .button--inline {
    color: inherit;
    font-size: inherit;
    display: inline-flex;
    text-decoration: underline;
    --button-disabled-foreground: color-mix(in srgb, currentColor, transparent 70%);

    &:hover:not([disabled]) {
      text-decoration: none;
    }
  }

  .button__loading-icon {
    position: absolute;
    top: 50%;
    left: 50%;
    translate: -50% -50%;
  }

  .button--loading > *:not(.button__loading-icon) {
    visibility: hidden;
  }

  @media (max-width: 768px) {
    .button--icon-only {
      padding: 6px;

      .button__text {
        display: none;
      }

      .button__prepend-icon {
        margin-left: 0;
        margin-right: 0;
      }
    }
  }
}
</style>
