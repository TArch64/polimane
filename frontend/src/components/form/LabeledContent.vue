<template>
  <div class="labeled-content">
    <label class="labeled-content__label">
      <slot />

      <span>
        {{ label }}
      </span>
    </label>
  </div>
</template>

<script setup lang="ts">
import type { Slot } from 'vue';
import type { NumericString } from '@/types';

withDefaults(defineProps<{
  label: string;
  gap?: number | NumericString;
}>(), {
  gap: 16,
});

defineSlots<{
  default: Slot;
}>();
</script>

<style scoped>
@layer components {
  .labeled-content {
    display: flex;
  }

  .labeled-content__label {
    display: flex;
    align-items: center;
    gap: v-bind("gap + 'px'");
  }
}
</style>
