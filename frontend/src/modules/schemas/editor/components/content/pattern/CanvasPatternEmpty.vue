<template>
  <KonvaGroup :config="groupConfig">
    <KonvaLabel ref="labelRef">
      <KonvaTag :config="labelTagConfig" />

      <CanvasStackH align="center" :gap="8" :config="stackConfig">
        <CanvasIcon size="22" color="rgba(0, 0, 0, 0.8)">
          <PlusIcon />
        </CanvasIcon>

        <KonvaText :config="labelTextConfig" />
      </CanvasStackH>
    </KonvaLabel>
  </KonvaGroup>
</template>

<script setup lang="ts">
import Konva from 'konva';
import { computed, type ComputedRef, onMounted, ref } from 'vue';
import type { KonvaLabel } from 'vue-konva';
import { PlusIcon } from '@/components/icon';
import { useNodeCentering } from '@/modules/schemas/editor/composables';
import { CanvasIcon, CanvasStackH } from '../base';

const labelRef = ref<InstanceType<KonvaLabel> | null>(null);
const labelNode: ComputedRef<Konva.Label> = computed(() => labelRef.value?.getNode() ?? null);

const groupConfig: Partial<Konva.GroupConfig> = {
  width: 1000,
  height: 100,
};

useNodeCentering(labelNode);

const labelTagConfig: Partial<Konva.TagConfig> = {
  stroke: 'rgba(0, 0, 0, 0.2)',
  strokeWidth: 1,
  cornerRadius: 8,
};

const stackConfig: Partial<Konva.GroupConfig> = {
  x: 12,
};

const labelTextConfig: Partial<Konva.TextConfig> = {
  width: 156,
  text: 'Додати Рядок',
  fill: 'rgba(0, 0, 0, 0.8)',
  fontSize: 15,
  height: 32,
  offsetY: -1,
  verticalAlign: 'middle',
};

onMounted(() => {

});
</script>
