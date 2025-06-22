<template>
  <div role="menu" popover="manual" class="dropdown">
    <slot />
  </div>
</template>

<script setup lang="ts">
import { computed, type Slot } from 'vue';

const props = withDefaults(defineProps<{
  control?: boolean;
  viewTransitionName?: string;
}>(), {
  control: false,
  viewTransitionName: 'dropdown',
});

defineSlots<{
  default: Slot;
}>();

const backgroundColor = computed(() => {
  return props.control ? 'var(--color-background-2)' : 'var(--color-background-1)';
});
</script>

<style scoped>
@layer components {
  .dropdown {
    background-color: v-bind("backgroundColor");
    border: var(--divider);
    border-radius: var(--rounded-md);
    box-shadow: var(--box-shadow);
    display: flex;
    flex-direction: column;
    gap: 4px;
    width: max-content;
    view-transition-name: v-bind("viewTransitionName");
  }
}
</style>
