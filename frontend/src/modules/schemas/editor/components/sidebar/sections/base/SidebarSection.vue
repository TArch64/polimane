<template>
  <section class="sidebar__section" :class="classes">
    <slot />
  </section>
</template>

<script setup lang="ts">
import { computed, type Slot } from 'vue';

const props = withDefaults(defineProps<{
  padding?: boolean;
}>(), {
  padding: true,
});

defineSlots<{
  default: Slot;
}>();

const classes = computed(() => ({
  'sidebar__section--padding': props.padding,
}));
</script>

<style scoped>
@layer page {
  .sidebar__section:not(:last-of-type) {
    border-bottom: var(--divider);
  }

  .sidebar__section--padding {
    padding: 16px 12px;
  }
}
</style>
