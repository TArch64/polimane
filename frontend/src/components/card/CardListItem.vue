<template>
  <li class="card-list-item">
    <p class="card-list-item__title">
      {{ title }}
    </p>

    <div class="card-list-item__actions" v-if="slots.actions">
      <slot name="actions" />
    </div>
  </li>
</template>

<script setup lang="ts">
import type { Slot } from 'vue';

defineProps<{
  title: string;
}>();

const slots = defineSlots<{
  default: Slot;
  actions?: Slot;
}>();
</script>

<style scoped>
@layer components {
  .card-list-item {
    padding-left: 6px;
    min-height: 40px;
    display: flex;
    align-items: center;
    min-width: 0;

    &:not(:last-child) {
      border-bottom: var(--divider)
    }
  }

  .card-list-item__title {
    flex: 1;
    min-width: 0;
    white-space: nowrap;
    text-overflow: ellipsis;
    overflow: hidden;
    margin-right: 40px;
  }

  .card-list-item__actions {
    margin-left: auto;
  }
}
</style>
