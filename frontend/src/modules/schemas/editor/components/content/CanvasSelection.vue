<template>
  <rect
    v-bind="visibleSelection"
    :fill="selectionFill"
    stroke="var(--color-divider)"
    class="canvas-selection"
  />
</template>

<script setup lang="ts">
import { computed, reactive } from 'vue';
import { useEventListener } from '@vueuse/core';
import { useCanvasStore, useEditorStore } from '@editor/stores';
import { useContrast } from '@editor/composables';
import type { INodeRect } from '@/models';

const props = defineProps<{
  canvasRef: SVGSVGElement;
}>();

const editorStore = useEditorStore();
const canvasStore = useCanvasStore();

const selection = reactive<INodeRect>({
  x: 0,
  y: 0,
  width: 0,
  height: 0,
});

const visibleSelection = computed(() => ({
  x: selection.width < 0 ? selection.x + selection.width : selection.x,
  y: selection.height < 0 ? selection.y + selection.height : selection.y,
  width: Math.abs(selection.width) * canvasStore.scale,
  height: Math.abs(selection.height) * canvasStore.scale,
}));

let abortController: AbortController;

useEventListener(props.canvasRef, 'mousedown', (event: MouseEvent) => {
  selection.x = event.clientX;
  selection.y = event.clientY;
  abortController = new AbortController();

  addEventListener('mousemove', (moveEvent: MouseEvent) => {
    selection.width += moveEvent.movementX;
    selection.height += moveEvent.movementY;
  }, { signal: abortController.signal });

  addEventListener('mouseup', () => {
    abortController.abort();
    selection.x = 0;
    selection.y = 0;
    selection.width = 0;
    selection.height = 0;
  }, { signal: abortController.signal });
});

const contrast = useContrast(() => editorStore.schema.backgroundColor, '#000');
const selectionColor = computed(() => contrast.value > 4.5 ? 'var(--color-black)' : 'var(--color-white)');
const selectionFill = computed(() => `color-mix(in srgb, ${selectionColor.value}, transparent 80%)`);
</script>
