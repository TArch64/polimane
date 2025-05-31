<template>
  <div class="color-list">
    <ColorItem
      v-for="(color, index) of model"
      :key="index"
      :model-value="color"
      @update:model-value="model[index] = $event"
    />
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import ColorItem from './ColorItem.vue';

const model = defineModel<string[]>({ required: true });

const columnCount = computed(() => Math.ceil(model.value.length / 2));
</script>

<style scoped>
@layer page {
  .color-list {
    display: grid;
    gap: 4px;
    grid-template-columns: repeat(v-bind("columnCount"), 20px);
    grid-template-rows: repeat(2, 20px);
  }
}
</style>
