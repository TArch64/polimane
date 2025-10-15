<template>
  <div class="selection-area">
    <slot />
  </div>
</template>

<script lang="ts" setup>
import { useCanvasPosition } from '@editor/composables';
import type { Slot } from 'vue';

const props = defineProps<{
  selectionRef: SVGElement;
}>();

defineSlots<{
  default: Slot;
}>();

const position = useCanvasPosition(() => props.selectionRef);
</script>

<style scoped>
@layer page {
  .selection-area {
    z-index: 1000;
    position: fixed;
    top: 0;
    left: 0;
    width: v-bind("position.width + 'px'");
    height: v-bind("position.height + 'px'");
    translate: v-bind("position.x + 'px'") v-bind("position.y + 'px'");
    pointer-events: none;
  }
}
</style>
