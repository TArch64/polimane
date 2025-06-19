<template>
  <div class="color-list">
    <slot />
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';

const props = defineProps<{
  size: number;
}>();

const columnCount = computed(() => Math.ceil(props.size / 2));
</script>

<style scoped>
@layer page {
  .color-list {
    display: grid;
    gap: 4px;
    grid-template-columns: repeat(v-bind("columnCount"), var(--color-button-size));
    grid-template-rows: repeat(2, var(--color-button-size));
  }
}
</style>
