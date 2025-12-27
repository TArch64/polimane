<template>
  <div class="home-list-empty">
    <div class="home-list-empty__description">
      <slot name="description" />
    </div>

    <div class="home-list-empty__actions" v-if="slots.actions">
      <slot name="actions" />
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Slot } from 'vue';

const slots = defineSlots<{
  description: Slot;
  actions?: Slot;
}>();
</script>

<style scoped>
@layer page {
  .home-list-empty {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    flex-grow: 1;
  }

  .home-list-empty__description {
    margin-bottom: 12px;
    text-align: center;

    :deep(p:not(:first-child)) {
      margin-top: 8px;
    }
  }

  .home-list-empty__actions {
    display: flex;
    gap: 8px;
    align-items: center;
  }
}
</style>
