<template>
  <div inert hidden ref="anchorRef" />

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
import { computed, nextTick, type Ref, ref, type Slot, type VNodeRef } from 'vue';
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

const anchorRef = ref<HTMLElement>(null!);
const currentEl = computed(() => anchorRef.value.parentElement!);

const SCROLL_OFFSET = 200;
const SCROLL_STEP = 10;

const area = ref<NodeRect>(NodeRect.BLANK.clone());
const visibleArea = computed(() => area.value.normalized);

const selected: Ref<Set<I['id']>> = defineModel({ required: true });
const registry: SelectionListRegistry<I> = new Map();

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

function isSafeZone(area: NodeRect) {
  return Math.abs(area.width) < 10 && Math.abs(area.height) < 10;
}

function onMouseMove(downEvent: MouseEvent): (event: MouseEvent) => void {
  return (event) => {
    clearInterval(scrollInterval);
    area.value.width += event.movementX;
    area.value.height += event.movementY;

    if (!downEvent.defaultPrevented && !isSafeZone(area.value)) {
      downEvent.preventDefault();
      area.value.x = downEvent.clientX;
      area.value.y = downEvent.clientY + window.scrollY;
    }

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
  if (isSafeZone(area)) {
    return;
  }

  for (const item of registry.values()) {
    const rect = new NodeRect(item.el.getBoundingClientRect()).delta({ y: window.scrollY });

    if (area.isIntersecting(rect)) {
      selected.value.add(item.data.id);
    }
  }
}

useEventListener('mousedown', async (event: MouseEvent) => {
  const target = event.target as Element;

  const canSelect = event.buttons === 1 && (
    currentEl.value.contains(target)
    || target.matches('#app')
    || target.closest('[data-cursor-selection-include]')
  );

  if (!canSelect) return;

  await nextTick();
  selected.value.clear();

  if (event.buttons !== 1) {
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
