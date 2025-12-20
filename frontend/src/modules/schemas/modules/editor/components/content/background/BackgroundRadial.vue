<template>
  <BackgroundPattern :scale="2">
    <BackgroundCell
      v-for="(pos, index) in patternPositions"
      :key="index"
      :x="pos.x"
      :y="pos.y"
    />
  </BackgroundPattern>
</template>

<script setup lang="ts">
import { BEAD_SIZE } from '@editor/const';
import type { IPoint } from '@/models';
import BackgroundPattern from './BackgroundPattern.vue';
import BackgroundCell from './BackgroundCell.vue';

const patternPositions: IPoint[] = [
  { x: 0, y: 0 },
  { x: BEAD_SIZE, y: 0 },
  { x: BEAD_SIZE / 2, y: BEAD_SIZE },
  { x: BEAD_SIZE * 1.5, y: BEAD_SIZE },
  { x: -BEAD_SIZE / 2, y: BEAD_SIZE },
];
</script>
