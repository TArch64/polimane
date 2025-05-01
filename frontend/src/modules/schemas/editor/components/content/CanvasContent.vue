<template>
  <KonvaLayer>
    <CanvasStackV ref="stackRef" :config="stackConfig" :gap="16">
      <CanvasPattern
        v-for="pattern of editorStore.schema.content"
        :key="pattern.id"
        :pattern
      />
    </CanvasStackV>
  </KonvaLayer>
</template>

<script setup lang="ts">
import Konva from 'konva';
import { useEditorStore } from '@/modules/schemas/editor/stores';
import { useNodeCentering, useNodeRef } from '../../composables';
import { CanvasStackV } from './base';
import { CanvasPattern } from './pattern';

const editorStore = useEditorStore();
const stackRef = useNodeRef<Konva.Group>();

const stackConfig = useNodeCentering(stackRef, {
  padding: {
    horizontal: 20,
    vertical: 16,
  },
});
</script>
