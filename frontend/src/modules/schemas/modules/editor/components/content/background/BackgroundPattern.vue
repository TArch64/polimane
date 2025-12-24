<template>
  <pattern :width :height x="0" y="0" patternUnits="userSpaceOnUse">
    <slot />
  </pattern>
</template>

<script setup lang="ts">
import type { Slot } from 'vue';
import { BEAD_SIZE } from '@editor/const';

const props = defineProps<{
  scale: number;
}>();

const width = BEAD_SIZE * props.scale;
const height = BEAD_SIZE * props.scale;

defineSlots<{
  default: Slot;
}>();
</script>
