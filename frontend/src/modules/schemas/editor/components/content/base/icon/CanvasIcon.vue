<template>
  <KonvaImage :config />

  <Teleport :to="iconSource.hostEl">
    <slot />
  </Teleport>
</template>

<script setup lang="ts">
import { computed, type Slot, toRef, toRefs } from 'vue';
import Konva from 'konva';
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

const iconSource = useIconSource(toRefs(props));

const iconImageEl = useIconImage({
  source: toRef(iconSource, 'source'),
});

const config = computed((): Partial<Konva.ImageConfig> => ({
  width: Number(props.size),
  height: Number(props.size),
  image: iconImageEl,
}));
</script>
