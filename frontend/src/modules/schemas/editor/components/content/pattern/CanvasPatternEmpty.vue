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
import { computed } from 'vue';
import { PlusIcon } from '@/components/icon';
import { useThemeVar } from '@/composables';
import { CanvasIcon, CanvasStackH, GroupRenderer } from '../base';

const colorDivider = useThemeVar('--color-divider');
const colorPrimary = useThemeVar('--color-primary');
const roundedMd = useThemeVar('--rounded-md');

const backgroundConfig: Partial<Konva.GroupConfig> = computed(() => ({
  stroke: colorDivider.value,
  strokeWidth: 1,
  cornerRadius: roundedMd.value,
  width: 156,
  height: 32,
}));

const stackConfig: Partial<Konva.GroupConfig> = {
  x: 12,
};

const labelTextConfig: Partial<Konva.TextConfig> = computed(() => ({
  text: 'Додати Рядок',
  fill: colorPrimary.value,
  fontSize: 15,
  height: 32,
  offsetY: -1,
  verticalAlign: 'middle',
}));
</script>
