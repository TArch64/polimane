<template>
  <component
    :is="binding.is"
    v-bind="binding.props"
  >
    <slot />
  </component>
</template>

<script setup lang="ts">
import { type Component, computed, type HTMLAttributes, type IntrinsicElementAttributes, type Slot } from 'vue';
import { type RouteLocationRaw, RouterLink } from 'vue-router';
import type { ComponentAs, ComponentProps, ComponentTag } from '@/types';

const props = defineProps<{
  to?: RouteLocationRaw;
}>();

defineSlots<{
  default: Slot;
}>();

interface IBinding<P extends object> {
  is: ComponentAs;
  props: P;
}

type ComponentBinding<C extends Component> = IBinding<Omit<HTMLAttributes, 'is'> & ComponentProps<C>>;
type TagBinding<T extends ComponentTag> = IBinding<IntrinsicElementAttributes[T]>;
type Binding<C extends ComponentTag | Component> = C extends ComponentTag ? TagBinding<C> : ComponentBinding<Extract<C, Component>>;

function makeBinding<C extends ComponentTag | Component>(is: C, props: () => Binding<C>['props']) {
  return computed(() => ({ is, props: props() }));
}

const linkBinding = makeBinding(RouterLink, () => ({
  to: props.to!,
  viewTransition: true,
}));

const buttonBinding = makeBinding('button', () => ({
  type: 'button',
}));

const binding = computed(() => {
  return props.to ? linkBinding.value : buttonBinding.value;
});
</script>
