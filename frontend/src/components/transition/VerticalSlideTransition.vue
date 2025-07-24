<template>
  <Transition name="vertical-slice-transition-" :duration>
    <slot />
  </Transition>
</template>

<script setup lang="ts">
import type { Slot } from 'vue';

withDefaults(defineProps<{
  duration?: number;
  shift?: number;
}>(), {
  duration: 200,
  shift: 0,
});

defineSlots<{
  default: Slot;
}>();
</script>

<style scoped>
@layer components {
  .vertical-slice-transition--enter-from,
  .vertical-slice-transition--leave-to {
    height: 0 !important;
    padding: 0 !important;
    margin: v-bind("shift + 'px'") 0 0 !important;
    opacity: 0 !important;
  }

  .vertical-slice-transition--enter-active,
  .vertical-slice-transition--leave-active {
    --duration: v-bind("duration + 'ms'");

    overflow: hidden;
    transform-origin: center center;
    will-change: height, padding, margin, opacity;

    transition: var(--duration) height ease-out,
    var(--duration) padding ease-out,
    var(--duration) margin ease-out,
    var(--duration-2) opacity ease-out;
  }

  .vertical-slice-transition--enter-active {
    --duration-2: var(--duration);
  }

  .vertical-slice-transition--leave-active {
    --duration-2: calc(var(--duration) / 2);
  }
}
</style>
