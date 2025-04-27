<template>
  <KonvaLayer>
    <CanvasStackV ref="stackRef" :gap="16">
      <CanvasPattern
        v-for="pattern of editorStore.schema.content"
        :key="pattern.id"
        :pattern
      />
    </CanvasStackV>
  </KonvaLayer>
</template>

<script setup lang="ts">
import { computed, type ComputedRef, ref } from 'vue';
import type { ComponentExposed } from 'vue-component-type-helpers';
import Konva from 'konva';
import { useEditorStore } from '@/modules/schemas/editor/stores';
import { useNodeCentering } from '../../composables';
import { CanvasStackV } from './base';
import { CanvasPattern } from './pattern';

const editorStore = useEditorStore();
const stackRef = ref<ComponentExposed<typeof CanvasStackV> | null>(null);
const stackNode: ComputedRef<Konva.Group | null> = computed(() => stackRef.value?.groupNode as Konva.Group);

useNodeCentering(stackNode, {
  padding: {
    horizontal: 20,
    vertical: 16,
  },

  trigger: () => editorStore.schema.content,
});

</script>
