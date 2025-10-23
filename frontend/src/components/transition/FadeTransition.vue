<template>
  <Transition
    name="fade-transition-"
    :duration
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
}>(), {
  duration: () => ({ enter: 150, leave: 100 }),
  state: undefined,
});

defineSlots<{
  default: Slot;
}>();

const duration = computed(() => normalizeDuration(props.duration));
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
