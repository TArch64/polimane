<template>
  <CanvasDefs>
    <Component :is="backgroundComponent" :id="patternId" />
  </CanvasDefs>

  <rect :fill="`url(#${patternId})`" v-bind="rect" />
</template>

<script setup lang="ts">
import { useEditorStore } from '@editor/stores';
import { type Component, markRaw, useId } from 'vue';
import type { INodeRect } from '@/models';
import { SchemaLayout } from '@/enums';
import CanvasDefs from '../CanvasDefs.vue';
import BackgroundLinear from './BackgroundLinear.vue';
import BackgroundRadial from './BackgroundRadial.vue';

defineProps<{
  rect: INodeRect;
}>();

const editorStore = useEditorStore();

const patternId = `canvas-background-${editorStore.schema.layout}-pattern-${useId()}`;

const backgroundComponents: Record<SchemaLayout, Component> = {
  [SchemaLayout.LINEAR]: markRaw(BackgroundLinear),
  [SchemaLayout.RADIAL]: markRaw(BackgroundRadial),
};

const backgroundComponent = backgroundComponents[editorStore.schema.layout];
</script>
