<template>
  <Transition name="vertical-slide-transition-" :duration>
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
  .vertical-slide-transition--enter-from,
  .vertical-slide-transition--leave-to {
    height: 0 !important;
    padding: 0 !important;
    margin: v-bind("shift + 'px'") 0 0 !important;
    opacity: 0 !important;

    :deep(.vertical-slide-transition__item) {
      translate: 0 calc(0px - v-bind("shift + 'px'"));
    }
  }

  .vertical-slide-transition--enter-active,
  .vertical-slide-transition--leave-active {
    --duration: v-bind("duration + 'ms'");

    overflow: hidden;
    will-change: height, padding, margin, opacity;

    transition: var(--duration) height ease-out,
    var(--duration) padding ease-out,
    var(--duration) margin ease-out,
    var(--duration-2) opacity ease-out;

    :deep(.vertical-slide-transition__item) {
      transition: var(--duration) translate ease-out;
    }
  }

  .vertical-slide-transition--enter-active {
    --duration-2: var(--duration);
  }

  .vertical-slide-transition--leave-active {
    --duration-2: calc(var(--duration) / 2);
  }
}
</style>
