<template>
  <Transition name="selection-resize-handle-" appear :duration="200">
    <div
      class="selection-resize-handle"
      :class="classes"
      @mousedown="onMouseDown"
    />
  </Transition>
</template>

<script setup lang="ts">
import { computed, toRef } from 'vue';
import { useBackgroundContrast } from '@editor/composables';
import { useSelectionStore } from '@editor/stores';
import { Direction, isNegativeDirection, isVerticalDirection } from '@/enums';

const props = defineProps<{
  direction: Direction;
}>();

const overlay = defineModel<boolean>('overlay', { required: true });

const selectionStore = useSelectionStore();
const translation = toRef(selectionStore.resize.translation, props.direction);

const isVertical = isVerticalDirection(props.direction);
const isNegative = isNegativeDirection(props.direction);

const classes = computed(() => `selection-resize-handle--${props.direction}`);

const contrast = useBackgroundContrast('#FFF');
const backgroundColor = computed(() => contrast.isAA ? 'var(--color-primary)' : 'var(--color-white)');
const borderColor = computed(() => contrast.isAA ? 'var(--color-white)' : 'var(--color-primary)');

function onMouseMove(event: MouseEvent) {
  const axis = isVertical ? 'movementY' : 'movementX';
  const direction = isNegative ? -1 : 1;
  const value = translation.value + direction * event[axis];
  translation.value = Math.max(value, 0);
}

function onMouseUp() {
  overlay.value = false;
  selectionStore.resize.reset();
  removeEventListener('mousemove', onMouseMove);
}

function onMouseDown() {
  overlay.value = true;
  addEventListener('mousemove', onMouseMove);
  addEventListener('mouseup', onMouseUp, { once: true });
}
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
    will-change: background-color, border-color, v-bind("direction");
    --handle-main-size: max(calc(100% / 5), 40px);
    --handle-cross-size: 5px;
    --handle-offset: calc(-20px - v-bind("translation + 'px'"));
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

  .selection-resize-handle--enter-from,
  .selection-resize-handle--leave-to {
    opacity: 0;
    scale: 0.8;
  }

  .selection-resize-handle--enter-active,
  .selection-resize-handle--leave-active {
    transform-origin: center center;
    transition: opacity 0.2s ease-out, scale 0.2s ease-out;
    will-change: opacity, scale;
  }
}
</style>
