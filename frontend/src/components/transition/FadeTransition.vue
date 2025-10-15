<template>
  <Transition name="fade-transition-" :duration @after-leave="$emit('after-leave')">
    <slot />
  </Transition>
</template>

<script setup lang="ts">
import { computed, type Slot } from 'vue';
import { normalizeDuration, type TransitionDuration } from './TransitionDuration';

const props = withDefaults(defineProps<{
  duration?: TransitionDuration;
}>(), {
  duration: () => ({ enter: 150, leave: 100 }),
});

defineEmits<{
  'after-leave': [];
}>();

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
