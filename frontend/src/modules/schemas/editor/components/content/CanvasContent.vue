<template>
  <defs>
    <pattern
      x="0"
      y="0"
      id="editorEmptyBeads"
      patternUnits="userSpaceOnUse"
      :width="BEAD_SIZE"
      :height="BEAD_SIZE"
    >
      <rect
        x="0"
        y="0"
        :width="BEAD_SIZE"
        :height="BEAD_SIZE"
        :fill="beadsStore.emptyColor"
        fill-opacity="0.2"
        :stroke="beadsStore.emptyColor"
        stroke-width="0.5"
        class="canvas-content__background-bead"
      />
    </pattern>
  </defs>

  <g :transform class="canvas-content" v-on="listeners">
    <rect
      ref="backgroundRectRef"
      fill="url(#editorEmptyBeads)"
      :x="beadsGrid.size.minX"
      :y="beadsGrid.size.minY"
      :width="beadsGrid.size.width"
      :height="beadsGrid.size.height"
    />

    <CanvasBead
      v-for="bead of beadsGrid.beads"
      :key="bead.coord"
      :offset="bead.offset"
      :coord="bead.coord"
      :color="bead.color"
    />
  </g>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { BEAD_SIZE, type ICanvasZoom, useBeadPainting, useBeadsGrid } from '../../composables';
import { useBeadsStore, useEditorStore } from '../../stores';
import { CanvasBead } from './CanvasBead';

const props = defineProps<{
  wrapperRect: DOMRect;
  canvasZoom: ICanvasZoom;
}>();

const backgroundRectRef = ref<SVGRectElement>(null!);

const beadsStore = useBeadsStore();
const editorStore = useEditorStore();

const listeners = useBeadPainting({
  backgroundRectRef,
  canvasZoom: props.canvasZoom,
});

const beadsGrid = useBeadsGrid(() => editorStore.schema);

const transform = (() => {
  const y = (props.wrapperRect.height - beadsGrid.size.height) / 2;
  const x = (props.wrapperRect.width - beadsGrid.size.width) / 2;

  return `translate(${x}, ${y})`;
})();
</script>

<style scoped>
@layer page {
  .canvas-content:hover {
    cursor: crosshair;
  }

  .canvas-content__background-bead {
    transition: fill 0.15s ease-out;
  }
}
</style>
