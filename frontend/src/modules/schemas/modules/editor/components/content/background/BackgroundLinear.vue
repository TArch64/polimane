<template>
  <CanvasDefs>
    <pattern
      :id
      x="0"
      y="0"
      patternUnits="userSpaceOnUse"
      :width="BEAD_SIZE"
      :height="BEAD_SIZE"
    >
      <BackgroundCell x="0" y="0" />
    </pattern>
  </CanvasDefs>

  <rect
    ref="backgroundRef"
    :fill="`url(#${id})`"
    v-bind="rect"
  />
</template>

<script setup lang="ts">
import { BEAD_SIZE } from '@editor/const';
import { useId } from 'vue';
import type { INodeRect } from '@/models';
import CanvasDefs from '../CanvasDefs.vue';
import BackgroundCell from './BackgroundCell.vue';

defineProps<{
  rect: INodeRect;
}>();

const id = `canvas-background-linear-pattern-${useId()}`;
</script>
