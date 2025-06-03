<template>
  <div class="color-list">
    <ColorItem
      v-for="(color, index) of model"
      :key="index"
      :model-value="color"
      :active="activeIndex === index"
      @update:model-value="model[index] = $event"
      @update:active="activeIndex = index"
    />
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import ColorItem from './ColorItem.vue';

const model = defineModel<string[]>({ required: true });
const activeIndex = defineModel<number>('active-index', { required: true });

const columnCount = computed(() => Math.ceil(model.value.length / 2));
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
