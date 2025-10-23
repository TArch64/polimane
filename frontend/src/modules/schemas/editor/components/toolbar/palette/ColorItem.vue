<template>
  <ToolbarButton class="color-item" :active :class="classes">
    <slot />
  </ToolbarButton>
</template>

<script setup lang="ts">
import { computed, type Slot } from 'vue';
import { useContrast } from '@editor/composables';
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

const valueContrast = useContrast(() => props.color || '#FFF', '#FFF');
const hoverVar = computed(() => valueContrast.value > 4.5 ? '#fff' : '#999');
</script>

<style scoped>
@layer page {
  .color-item {
    border: var(--divider);
  }

  .button--active:not(.color-item--empty) {
    border-color: var(--color-white);
    outline: solid 1px var(--color-primary);
    outline-offset: 1px;
  }

  .color-item--value {
    --button-background: v-bind("color");
    --button-hover-background: color-mix(in srgb, v-bind("color"), v-bind("hoverVar") 20%);
  }

  .color-item--empty {
    background-image: url("@/assets/emptyColor.svg");
    background-size: 100% 100%;
  }
}
</style>
