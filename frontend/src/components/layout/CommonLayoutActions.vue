<template>
  <div ref="actionsRef" class="common-layout-top-bar__actions">
    <slot />
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref, type Slot } from 'vue';
import { useMutationObserver } from '@vueuse/core';

const hasActions = defineModel<boolean>('has-actions', {
  required: true,
});

defineSlots<{
  default: Slot;
}>();

const actionsRef = ref<HTMLElement>(null!);

function update() {
  const wrapper = actionsRef.value.querySelector('[data-bar-actions-wrapper]')
    || actionsRef.value;

  hasActions.value = wrapper.children.length > 0;
}

onMounted(update);

useMutationObserver(actionsRef, update, {
  childList: true,
  subtree: true,
});
</script>

<style scoped>
@layer page {
  .common-layout-top-bar__actions {
    margin-left: auto;
    margin-right: 2px;
    display: flex;
    gap: 8px;
    align-items: center;
  }
}
</style>
