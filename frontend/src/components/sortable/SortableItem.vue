<template>
  <Component
    :is="as"
    :ref="setRef"
    class="sortable-item"
    :class="classes"
    :style="styles"
  >
    <slot />

    <SortableIndicator
      :gap
      :elRef
      :direction
      :sortableAnchorVar
      :edge="closestEdge"
      v-if="closestEdge && elRef"
    />

    <Teleport :to="previewContainer" v-if="previewContainer">
      <slot name="preview" :item="item as I" />
    </Teleport>
  </Component>
</template>

<script setup lang="ts" generic="I extends ISortableEntity">
import { computed, type Ref, ref, type Slot, type VNodeRef } from 'vue';
import { type MaybeComputedElementRef, unrefElement } from '@vueuse/core';
import {
  attachClosestEdge,
  extractClosestEdge,
} from '@atlaskit/pragmatic-drag-and-drop-hitbox/closest-edge';
import type { Edge } from '@atlaskit/pragmatic-drag-and-drop-hitbox/types';
import {
  setCustomNativeDragPreview,
} from '@atlaskit/pragmatic-drag-and-drop/element/set-custom-native-drag-preview';
import type { ComponentAs } from '@/types';
import { useDraggable, useDropTarget } from './composables';
import SortableIndicator from './SortableIndicator.vue';
import type { ISortableEntity } from './ISortableEntity';
import type { DragDirection } from './DragDirection';

const props = defineProps<{
  group: string;
  item: I;
  list: I[];
  gap: number;
  as: ComponentAs;
  direction: DragDirection;
  sortableAnchorVar: string;
}>();

const slots = defineSlots<{
  default: Slot;
  preview?: Slot<{ item: I }>;
}>();

const elRef = ref<HTMLElement | null>(null);
const closestEdge = ref<Edge | null>(null);
const previewContainer = ref<HTMLElement | null>(null);
const isDragging = ref(false);

const setRef: VNodeRef = (el) => {
  elRef.value = unrefElement(el as MaybeComputedElementRef) as HTMLElement;
};

const classes = computed(() => ({
  'sortable-item--dragging': isDragging.value,
}));

const styles = computed(() => closestEdge.value ? { anchorName: props.sortableAnchorVar } : {});

useDraggable({
  element: elRef as Ref<HTMLElement>,

  getInitialData: () => ({
    item: props.item,
    group: props.group,
  }),

  dragHandle: computed(() => elRef.value!.querySelector('[data-drag-handle]') || elRef.value!),

  onDragStart: () => isDragging.value = true,
  onDrop: () => isDragging.value = false,

  onGenerateDragPreview({ nativeSetDragImage }) {
    setCustomNativeDragPreview({
      nativeSetDragImage,

      getOffset: slots.preview
        ? undefined
        : () => ({
            x: elRef.value!.offsetWidth / 4,
            y: elRef.value!.offsetHeight / 4,
          }),

      render({ container }) {
        if (!slots.preview) {
          const previewEl = elRef.value!.cloneNode(true) as HTMLElement;
          previewEl.style.width = `${elRef.value!.offsetWidth}px`;
          previewEl.style.scale = '0.5';
          container.append(previewEl);
          return;
        }
        previewContainer.value = container;
        return () => previewContainer.value = null;
      },
    });
  },
});

useDropTarget({
  element: elRef as Ref<HTMLElement>,
  getData({ input }) {
    return attachClosestEdge({ item: props.item }, {
      element: elRef.value!,
      input,
      allowedEdges: props.direction === 'horizontal' ? ['left', 'right'] : ['top', 'bottom'],
    });
  },
  onDrag: ({ self, source }) => {
    const isSource = source.element === elRef.value;

    if (isSource) {
      closestEdge.value = null;
      return;
    }

    closestEdge.value = extractClosestEdge(self.data);
  },
  onDragLeave: () => closestEdge.value = null,
  onDrop: () => closestEdge.value = null,

  canDrop: (args) => {
    return args.source.data.group === props.group;
  },
});
</script>

<style scoped>
@layer components {
  .sortable-item {
    transition: opacity 150ms ease-out;
  }

  .sortable-item--dragging {
    opacity: 0.5;
  }
}
</style>
