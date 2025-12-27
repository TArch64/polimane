<template>
  <Component
    :is="binding.is"
    v-bind="binding.props"
    class="card"
    :class="classes"
  >
    <header v-if="title">
      <h2 class="card__title">
        {{ title }}
      </h2>
    </header>

    <slot />

    <VerticalSlideTransition v-bind="footerTransition">
      <footer class="card__footer" v-if="slots.footer">
        <slot name="footer" />
      </footer>
    </VerticalSlideTransition>
  </Component>
</template>

<script setup lang="ts">
import { computed, type Slot } from 'vue';
import type { ComponentAs } from '@/types';
import type { AnyBinding } from '../binding';
import { VerticalSlideTransition } from '../transition';
import type { ICardFooterTransition } from './ICardFooterTransition';

const props = withDefaults(defineProps<{
  as?: ComponentAs;
  binding?: AnyBinding;
  interactable?: boolean;
  active?: boolean;
  variant?: 'main' | 'control';
  title?: string;
  footerTransition?: Partial<ICardFooterTransition>;
}>(), {
  as: 'div',
  interactable: false,
  active: false,
  variant: 'main',
  title: '',

  footerTransition: () => ({
    duration: 300,
    shift: 0,
  }),

  binding: (props): AnyBinding => ({
    is: props.as ?? 'div',
    props: {},
  }),
});

const slots = defineSlots<{
  default: Slot;
  footer?: Slot;
}>();

const classes = computed(() => [
  `card--variant-${props.variant}`,
  {
    'card--interactable tap-animation': props.interactable,
    'card--active': props.active,
  },
]);
</script>

<style scoped>
@layer components {
  .card {
    padding: var(--card-padding-top) var(--card-padding-right) var(--card-padding-bottom) var(--card-padding-left);
    background-color: var(--card-background);
    border: var(--divider);
    border-radius: var(--rounded-lg);

    --card-padding-top: 8px;
    --card-padding-bottom: 8px;
    --card-padding-left: 12px;
    --card-padding-right: 12px;
  }

  .card--variant-main {
    --card-background: var(--color-white);
  }

  .card--variant-control {
    --card-background: var(--color-background-2);
  }

  .card--interactable {
    transition: border-color 0.15s ease-out;
    will-change: border-color;

    &:hover:not(.card--active),
    &:focus:not(.card--active),
    &:focus-within:not(.card--active) {
      border-color: var(--color-hover-divider);
    }
  }

  .card--active {
    border-color: var(--color-primary);
  }

  .card__title {
    font-size: var(--font-md);
    font-weight: 500;
    padding: 4px 0;
    margin-bottom: 4px;
  }

  .card__footer {
    display: flex;
    gap: 8px;
    padding: 4px 0;
    margin-top: 4px;
    will-change: height, padding, margin, opacity;
  }
}
</style>
