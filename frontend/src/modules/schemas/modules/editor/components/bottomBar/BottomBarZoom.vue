<template>
  <BottomBarMetric
    label="Маштаб"
    :value="formattedValue"
    :max-value="formattedMax"
  />
</template>

<script setup lang="ts">
import { MAX_SCALE, useCanvasStore } from '@editor/stores';
import { useIntPercentageFormatter } from '@/composables';
import BottomBarMetric from './BottomBarMetric.vue';

const canvasStore = useCanvasStore();

const formattedValue = useIntPercentageFormatter(() => canvasStore.scale);
const formattedMax = useIntPercentageFormatter(() => MAX_SCALE);
</script>
