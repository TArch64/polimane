<template>
  <GroupRenderer :config="rootConfig">
    <KonvaRect :config="backgroundConfig" />

    <CanvasStackH ref="contentRef" align="center" :gap="8">
      <CanvasIcon size="22" color="--color-primary">
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
import {
  SCREENSHOT_IGNORE,
  useNodeConfigs,
  useNodeFiller,
  useNodeRef,
} from '@/modules/schemas/editor/composables';
import { CanvasIcon, CanvasStackH, GroupRenderer } from '../base';

const colorDivider = useThemeVar('--color-divider');
const colorPrimary = useThemeVar('--color-primary');
const colorBackground1 = useThemeVar('--color-background-1');
const roundedMd = useThemeVar('--rounded-md');

const contentRef = useNodeRef<Konva.Group>();

const rootConfig: Partial<Konva.GroupConfig> = {
  name: SCREENSHOT_IGNORE,
};

const backgroundConfig = useNodeConfigs<Konva.RectConfig>([
  () => ({
    fill: colorBackground1.value,
    stroke: colorDivider.value,
    strokeWidth: 1,
    cornerRadius: roundedMd.value,
    offsetY: 3,
    offsetX: 16,
  }),

  useNodeFiller(contentRef, {
    padding: {
      vertical: 3,
      horizontal: 12,
    },
  }),
]);

const labelTextConfig: Partial<Konva.TextConfig> = computed(() => ({
  text: 'Додати Рядок',
  fill: colorPrimary.value,
  fontSize: 15,
  verticalAlign: 'middle',
}));
</script>
