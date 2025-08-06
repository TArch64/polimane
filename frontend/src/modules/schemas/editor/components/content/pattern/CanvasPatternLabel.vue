<template>
  <KonvaLabel :config>
    <KonvaTag :config="tagConfig" />
    <KonvaText ref="textRef" :config="textConfig" />
  </KonvaLabel>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import Konva from 'konva';
import type { ISchemaPattern } from '@/models';
import { SCREENSHOT_IGNORE, useNodeRef } from '@/modules/schemas/editor/composables';
import { useThemeVar } from '@/composables';

const props = defineProps<{
  pattern: ISchemaPattern;
}>();

const colorBackground1 = useThemeVar('--color-background-1');
const colorPrimary = useThemeVar('--color-primary');
const colorDivider = useThemeVar('--color-divider');
const roundedMd = useThemeVar('--rounded-sm');

const config: Partial<Konva.LabelConfig> = {
  name: SCREENSHOT_IGNORE,
  x: 12,
  y: -10,
};

const tagConfig: Partial<Konva.TagConfig> = computed(() => ({
  fill: colorBackground1.value,
  stroke: colorDivider.value,
  strokeWidth: 1,
  cornerRadius: roundedMd.value,
}));

const textRef = useNodeRef<Konva.Text | null>();

const labelTextWidth = computed(() => {
  return textRef.value?.measureSize(props.pattern.name).width;
});

const textConfig = computed((): Partial<Konva.TextConfig> => ({
  text: props.pattern.name,
  fill: colorPrimary.value,
  padding: 4,
  fontSize: 14,
  align: 'center',
  width: labelTextWidth.value ? labelTextWidth.value + 24 : 0,
}));
</script>
