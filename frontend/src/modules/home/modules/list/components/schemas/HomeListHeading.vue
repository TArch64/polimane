<template>
  <h2 class="home-list-heading">
    <slot />

    <span class="home-list-heading__action" v-if="slots.action">
      <slot name="action" />
    </span>
  </h2>
</template>

<script setup lang="ts">
import type { Slot } from 'vue';

const slots = defineSlots<{
  default: Slot;
  action?: Slot;
}>();
</script>

<style scoped>
@layer page {
  .home-list-heading {
    padding-left: 16px;
    padding-right: 8px;
    margin-bottom: 8px;
    font-weight: 450;
    font-size: var(--font-lg);
    color: var(--color-text-2);
    display: flex;
    align-items: center;

    &:first-child {
      margin-top: 20px;
    }
  }

  .home-list-heading__action {
    margin-left: auto;
  }
}
</style>
