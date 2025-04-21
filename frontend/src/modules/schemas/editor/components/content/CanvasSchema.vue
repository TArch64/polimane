<template>
  <CanvasList
    :items="editorStore.schema.content"
    :position-iterator-class="PatternPositionIterator"
    v-slot="{ item }"
  >
    <CanvasPattern :pattern="item as ISchemaPattern" />
  </CanvasList>
</template>

<script setup lang="ts">
import type { ISchemaPattern } from '@/models';
import { useEditorStore } from '@/modules/schemas/editor/stores';
import { injectCanvas, provideObjectParent } from '@/modules/schemas/editor/composables';
import { CanvasPattern, PatternPositionIterator } from './pattern';
import CanvasList from './CanvasList.vue';

const editorStore = useEditorStore();
provideObjectParent(injectCanvas());
</script>
