<template>
  <CursorSelectionItem
    v-for="item of list"
    :key="item.id"
    :item
    :registry
    v-slot="{ itemRef }"
  >
    <slot :item :itemRef />
  </CursorSelectionItem>

  <template v-if="!area.isBlank">
    <Teleport to="body">
      <div class="cursor-selection-overlay" />
    </Teleport>

    <div class="cursor-selection" :style="selectionStyles" />
  </template>
</template>

<script setup lang="ts" generic="I extends SelectionItem">
import { computed, type Ref, ref, type Slot, type VNodeRef } from 'vue';
import { useEventListener } from '@vueuse/core';
import { NodeRect } from '@/models';
import CursorSelectionItem from './CursorSelectionItem.vue';
import type { SelectionItem, SelectionListRegistry } from './ListRegistry';

defineProps<{
  list: I[];
}>();

defineSlots<{
  default: Slot<{
    item: I;
    itemRef: VNodeRef;
  }>;
}>();

const SCROLL_OFFSET = 200;
const SCROLL_STEP = 10;
const area = ref<NodeRect>(NodeRect.BLANK.clone());

const selected: Ref<Set<I['id']>> = defineModel({ required: true });
const registry: SelectionListRegistry<I> = new Map();

const visibleArea = computed(() => ({
  x: Math.min(area.value.x, area.value.x + area.value.width),
  y: Math.min(area.value.y, area.value.y + area.value.height),
  width: Math.abs(area.value.width),
  height: Math.abs(area.value.height),
}));

const selectionStyles = computed(() => ({
  top: `${visibleArea.value.y}px`,
  left: `${visibleArea.value.x}px`,
  width: `${visibleArea.value.width}px`,
  height: `${visibleArea.value.height}px`,
}));

let scrollInterval: number | undefined = undefined;

function scrollingSelect(step: number) {
  area.value.height += step;
  window.scrollTo({ top: window.scrollY + step });
}

function onMouseMove(downEvent: MouseEvent): (event: MouseEvent) => void {
  return (event) => {
    if (!downEvent.defaultPrevented) {
      downEvent.preventDefault();
      area.value.x = downEvent.clientX;
      area.value.y = downEvent.clientY + window.scrollY;
    }

    clearInterval(scrollInterval);
    area.value.width += event.movementX;
    area.value.height += event.movementY;

    if (event.clientY < SCROLL_OFFSET) {
      if (window.scrollY === 0) {
        return;
      }
      scrollingSelect(-SCROLL_STEP);
      scrollInterval = window.setInterval(() => scrollingSelect(-SCROLL_STEP), 30);
    } else if (window.innerHeight - event.clientY < SCROLL_OFFSET) {
      if (window.scrollY + window.innerHeight === document.body.scrollHeight) {
        return;
      }
      scrollingSelect(SCROLL_STEP);
      scrollInterval = window.setInterval(() => scrollingSelect(SCROLL_STEP), 30);
    }
  };
}

function selectItems(area: NodeRect) {
  if (area.width < 5 && area.height < 5) {
    return;
  }

  const newSelection = new Set<I['id']>();

  for (const item of registry.values()) {
    const rect = new NodeRect(item.el.getBoundingClientRect()).delta({ y: window.scrollY });

    if (area.isIntersecting(rect)) {
      newSelection.add(item.data.id);
    }
  }

  selected.value = newSelection;
}

useEventListener('mousedown', (event: MouseEvent) => {
  if ((event.target as Element).closest('[data-cursor-selection-ignore]')) {
    return;
  }

  const handler = onMouseMove(event);
  addEventListener('mousemove', handler);

  addEventListener('mouseup', () => {
    clearInterval(scrollInterval);
    removeEventListener('mousemove', handler);
    selectItems(new NodeRect(visibleArea.value));
    area.value = NodeRect.BLANK.clone();
  }, { once: true });
});
</script>

<style scoped>
@layer components {
  :global(body:has(.cursor-selection)) {
    user-select: none;
  }

  .cursor-selection {
    position: absolute;
    border: var(--divider);
    border-radius: var(--rounded-sm);
    background-color: color-mix(in srgb, var(--color-primary), transparent 80%);
    will-change: top, left, width, height;
  }

  .cursor-selection-overlay {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
  }
}
</style>
