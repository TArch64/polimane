<template>
  <section class="banner" :class="classes">
    <Component
      :is="prependIcon"
      :size="28"
      class="banner__prepend-icon"
      v-if="prependIcon"
    />

    <div class="banner__content">
      <slot />
    </div>
  </section>
</template>

<script setup lang="ts">
import { computed, type Slot } from 'vue';
import type { ComponentVariant } from '@/types';
import type { IconComponent } from '@/components/icon';

const props = withDefaults(defineProps<{
  type?: 'info' | 'warning' | 'danger';
  variant?: ComponentVariant;
  size?: 'sm' | 'md';
  prependIcon?: IconComponent;
}>(), {
  variant: 'main',
  size: 'md',
});

defineSlots<{
  default: Slot;
}>();

const classes = computed(() => [
  `banner--${props.variant}`,
  `banner--${props.size}`,
  props.type && `banner--${props.type}`,
]);
</script>

<style scoped>
@layer components {
  .banner {
    display: flex;

    :deep(:where(p, h1, h2, h3, h4):not(:first-child)) {
      margin-top: var(--banner-text-spacing);
    }
  }

  .banner--main {
    background-color: var(--color-background-1);
  }

  .banner--control {
    background-color: var(--color-background-2);
  }

  .banner--sm {
    --banner-text-spacing: 4px;
    padding: 8px 12px;
    border-radius: var(--rounded-md);
    font-size: var(--font-sm);
  }

  .banner--md {
    --banner-text-spacing: 6px;
    padding: 16px 20px;
    border-radius: var(--rounded-lg);
    font-size: var(--font-md);
  }

  .banner--danger {
    --button-icon-color: var(--color-danger);
  }

  .banner--warning {
    --button-icon-color: var(--color-warning);
  }

  .banner--info {
    --button-icon-color: var(--color-info);
  }

  .banner__prepend-icon {
    flex-shrink: 0;
    margin-right: 16px;
    color: var(--button-icon-color);
  }

  .banner__content {
    flex-grow: 1;
    flex-basis: 0;
    min-width: 0;
    align-self: center;
  }
}
</style>
