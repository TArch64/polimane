<template>
  <div class="color-list">
    <ColorItem
      v-for="(color, index) of colors"
      :key="index"
      :color
    />
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import ColorItem from './ColorItem.vue';

const props = defineProps<{
  colors: string[];
}>();

const columnCount = computed(() => Math.ceil(props.colors.length / 2));
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
