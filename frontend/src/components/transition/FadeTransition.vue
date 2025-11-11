<template>
  <Transition
    :appear
    :duration
    name="fade-transition-"
    @before-leave="onBeforeLeave"
    v-on="state?.listeners ?? {}"
  >
    <slot />
  </Transition>
</template>

<script setup lang="ts">
import { computed, type Slot } from 'vue';
import type { ITransitionState } from '@/composables';
import { normalizeDuration, type TransitionDuration } from './TransitionDuration';

const props = withDefaults(defineProps<{
  duration?: TransitionDuration;
  state?: ITransitionState;
  appear?: boolean;
  switch?: boolean;
}>(), {
  duration: () => ({ enter: 150, leave: 100 }),
  state: undefined,
  appear: false,
  switch: false,
});

defineSlots<{
  default: Slot;
}>();

const duration = computed(() => normalizeDuration(props.duration));

function onBeforeLeave(el_: Element) {
  if (!props.switch) return;
  const el = el_ as HTMLElement;
  const rect = el.getBoundingClientRect();
  el.style.position = 'fixed';
  el.style.width = `${rect.width}px`;
  el.style.height = `${rect.height}px`;
  el.style.top = `${rect.top}px`;
  el.style.left = `${rect.left}px`;
}
</script>

<style scoped>
@layer components {
  .fade-transition--enter-from,
  .fade-transition--leave-to {
    opacity: 0;
  }

  .fade-transition--enter-active,
  .fade-transition--leave-active {
    transition: var(--duration) opacity ease-out;
    will-change: opacity;
  }

  .fade-transition--enter-active {
    --duration: v-bind("duration.enter + 'ms'");
  }

  .fade-transition--leave-active {
    --duration: v-bind("duration.leave + 'ms'");
  }
}
</style>
