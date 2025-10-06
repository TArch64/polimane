<template>
  <ToolbarButton :active :class="classes">
    <slot />
  </ToolbarButton>
</template>

<script setup lang="ts">
import { computed, type Slot } from 'vue';
import ToolbarButton from '../ToolbarButton.vue';

const props = withDefaults(defineProps<{
  color: string;
  active?: boolean;
}>(), {
  active: false,
});

defineSlots<{
  default: Slot;
}>();

const classes = computed(() => {
  const modifier = props.color ? 'value' : 'empty';
  return `color-item--${modifier}`;
});
</script>

<style scoped>
@layer page {
  .color-item--value {
    --button-background: v-bind("color");
    --button-hover-background: color-mix(in srgb, v-bind("color"), white 20%);
  }

  .color-item--empty {
    background-image: url("@/assets/emptyColor.svg");
    background-size: 100% 100%;
  }
}
</style>
