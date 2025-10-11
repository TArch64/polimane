<template>
  <rect
    v-bind="selectionStore.selection"
    stroke="var(--color-divider)"
    :fill="selectionFill"
  />
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { useEditorStore, useSelectionStore } from '@editor/stores';
import { useContrast } from '@editor/composables';

const editorStore = useEditorStore();
const selectionStore = useSelectionStore();

const contrast = useContrast(() => editorStore.schema.backgroundColor, '#000');
const selectionColor = computed(() => contrast.value > 4.5 ? 'var(--color-black)' : 'var(--color-white)');
const selectionFill = computed(() => `color-mix(in srgb, ${selectionColor.value}, transparent 80%)`);
</script>
