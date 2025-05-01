<template>
  <GroupRenderer>
    <KonvaRect :config="backgroundConfig" />

    <CanvasStackH align="center" :gap="8" :config="stackConfig">
      <CanvasIcon size="22" color="rgba(0, 0, 0, 0.8)">
        <PlusIcon />
      </CanvasIcon>

      <KonvaText :config="labelTextConfig" />
    </CanvasStackH>
  </GroupRenderer>
</template>

<script setup lang="ts">
import Konva from 'konva';
import { PlusIcon } from '@/components/icon';
import { CanvasIcon, CanvasStackH, GroupRenderer } from '../base';

const backgroundConfig: Partial<Konva.GroupConfig> = {
  stroke: 'rgba(0, 0, 0, 0.2)',
  strokeWidth: 1,
  cornerRadius: 8,
  width: 156,
  height: 32,
};

const stackConfig: Partial<Konva.GroupConfig> = {
  x: 12,
};

const labelTextConfig: Partial<Konva.TextConfig> = {
  text: 'Додати Рядок',
  fill: 'rgba(0, 0, 0, 0.8)',
  fontSize: 15,
  height: 32,
  offsetY: -1,
  verticalAlign: 'middle',
};
</script>
