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
import { useNodeRef } from '@/modules/schemas/editor/composables';

const props = defineProps<{
  pattern: ISchemaPattern;
}>();

const config: Partial<Konva.LabelConfig> = {
  offsetX: -12,
  offsetY: 10,
};

const tagConfig: Partial<Konva.TagConfig> = {
  fill: '#fff',
  stroke: 'rgba(0, 0, 0, 0.2)',
  strokeWidth: 1,
  cornerRadius: 4,
};

const textRef = useNodeRef<Konva.Text | null>();

const labelTextWidth = computed(() => {
  return textRef.value?.measureSize(props.pattern.name).width;
});

const textConfig = computed((): Partial<Konva.TextConfig> => ({
  text: props.pattern.name,
  fill: '#000',
  padding: 4,
  fontSize: 14,
  align: 'center',
  width: labelTextWidth.value ? labelTextWidth.value + 24 : 0,
}));
</script>
