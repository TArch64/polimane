<template>
  <ColorItem title="Колір Фону" :color="model">
    <input
      ref="pickerRef"
      type="color"
      class="toolbar-background-color__picker"
      v-model="model"
    >
  </ColorItem>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { useEditorStore } from '@editor/stores';
import { ColorItem } from './palette';

const editorStore = useEditorStore();

const model = computed({
  get: () => editorStore.schema.backgroundColor,
  set: (color) => editorStore.schema.backgroundColor = color,
});
</script>

<style scoped>
@layer page {
  .toolbar-background-color__picker {
    opacity: 0;
    width: 100%;
    height: 100%;
    cursor: pointer;
  }
}
</style>
