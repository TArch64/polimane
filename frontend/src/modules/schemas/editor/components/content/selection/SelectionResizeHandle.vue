<template>
  <div class="selection-resize-handle" :class="classes" />
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { useBackgroundContrast } from '@editor/composables';

const props = defineProps<{
  position: 'top' | 'bottom' | 'left' | 'right';
}>();

const classes = computed(() => `selection-resize-handle--${props.position}`);

const contrast = useBackgroundContrast('#FFF');
const backgroundColor = computed(() => contrast.isAA ? 'var(--color-primary)' : 'var(--color-white)');
const borderColor = computed(() => contrast.isAA ? 'var(--color-white)' : 'var(--color-primary)');
</script>

<style scoped>
@layer page {
  .selection-resize-handle {
    pointer-events: initial;
    position: absolute;
    border-radius: var(--rounded-full);
    background-color: v-bind("backgroundColor");
    border: 1px solid v-bind("borderColor");
    transition: background-color 0.15s ease-out, border-color 0.15s ease-out;
    --handle-main-size: max(calc(100% / 5), 40px);
    --handle-cross-size: 8px;
    --handle-offset: -20px;
  }

  .selection-resize-handle--top,
  .selection-resize-handle--bottom {
    left: 50%;
    translate: -50%;
    width: var(--handle-main-size);
    height: var(--handle-cross-size);
    cursor: ns-resize;
  }

  .selection-resize-handle--top {
    top: var(--handle-offset);
  }

  .selection-resize-handle--bottom {
    bottom: var(--handle-offset);
  }

  .selection-resize-handle--left,
  .selection-resize-handle--right {
    top: 50%;
    translate: 0 -50%;
    width: var(--handle-cross-size);
    height: var(--handle-main-size);
    cursor: ew-resize;
  }

  .selection-resize-handle--left {
    left: var(--handle-offset);
  }

  .selection-resize-handle--right {
    right: var(--handle-offset);
  }
}
</style>
