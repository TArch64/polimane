<template>
  <KonvaImage ref="imageRef" :config />

  <Teleport :to="iconSource.hostEl">
    <slot />
  </Teleport>
</template>

<script setup lang="ts">
import { computed, type Slot, toRef, toRefs } from 'vue';
import Konva from 'konva';
import { useNodeRef } from '@/modules/schemas/editor/composables';
import { KonvaImage } from '../../konva';
import { useIconSource } from './useIconSource';
import { useIconImage } from './useIconImage';

const props = withDefaults(defineProps<{
  color?: string;
  size?: number | string;
}>(), {
  color: '',
  size: 20,
});

defineSlots<{
  default: Slot;
}>();

const imageRef = useNodeRef<Konva.Image>();

const iconSource = useIconSource(toRefs(props));

const iconImageEl = useIconImage({
  source: toRef(iconSource, 'source'),
});

const config = computed((): Konva.ImageConfig => ({
  image: iconImageEl,
  width: Number(props.size),
  height: Number(props.size),
}));

defineExpose({
  getNode: () => imageRef.value,
});
</script>
