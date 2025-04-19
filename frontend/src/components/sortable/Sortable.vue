<template>
  <TransitionGroup
    name="sortable__item-"
    :duration="200"
    @before-leave="onBeforeItemLeave"
  >
    <SortableItem
      v-for="(item, index) of list"
      :key="item.id"
      :as="itemAs"
      :group
      :item
      :list
      :gap
      :direction
      :sortableAnchorVar
    >
      <template #default="ctx">
        <slot :item v-bind="ctx" :index />
      </template>

      <template #preview="ctx" v-if="slots.preview">
        <slot name="preview" v-bind="ctx" :index />
      </template>
    </SortableItem>
  </TransitionGroup>
</template>

<script setup lang="ts" generic="I extends ISortableEntity">
import { inject, ref, type Slot, watch } from 'vue';
import { reorder } from '@atlaskit/pragmatic-drag-and-drop/reorder';
import { extractClosestEdge } from '@atlaskit/pragmatic-drag-and-drop-hitbox/closest-edge';
import type { Edge } from '@atlaskit/pragmatic-drag-and-drop-hitbox/types';
import type { ComponentAs } from '@/types';
import { TOKEN_SCROLLER } from '@/InjectionToken';
import { newId } from '@/helpers';
import SortableItem from './SortableItem.vue';
import { useAutoScroll, useMonitor } from './composables';
import type { ISortableEntity } from './ISortableEntity';
import type { IMoveEvent } from './IMoveEvent';
import { type DragDirection, getAfterDirection, getBeforeDirection } from './DragDirection';

const props = withDefaults(defineProps<{
  group: string;
  list: I[];
  gap: number;
  itemAs: ComponentAs;
  direction?: DragDirection;
}>(), {
  direction: 'horizontal',
});

const emit = defineEmits<{
  move: [event: IMoveEvent<I>];
}>();

const slots = defineSlots<{
  default: Slot<{ item: I; index: number }>;
  preview?: Slot<{ item: I; index: number }>;
}>();

const scrollerEl = inject(TOKEN_SCROLLER)!;
const isDragging = ref(false);

const sortableAnchorVar = `--sortable-${newId()}`;

watch(isDragging, (value) => {
  document.body.classList.toggle('m-cursor--grabbing', value);
});

function getEdgedIndexes(from: number, to: number, edge: Edge): number {
  if (getBeforeDirection(props.direction) && to < from) return to;
  if (getAfterDirection(props.direction) && to > from) return to;
  return edge === getBeforeDirection(props.direction) ? to - 1 : to + 1;
}

useMonitor({
  onDragStart() {
    isDragging.value = true;
  },

  onDrop({ location, source }) {
    isDragging.value = false;

    const target = location.current.dropTargets[0]?.data.item as I;
    const activeItem = source.data.item as I;
    if (!target || activeItem.id === target.id) return;

    const edge = extractClosestEdge(location.current.dropTargets[0].data);
    const fromIndex = props.list.findIndex((item) => item.id === activeItem.id);
    const toIndex = props.list.findIndex((item) => item.id === target.id);
    const toIndexEdged = getEdgedIndexes(fromIndex, toIndex, edge!);

    if (fromIndex === toIndexEdged) return;

    const updated = reorder<I>({
      list: props.list,
      startIndex: fromIndex,
      finishIndex: toIndexEdged,
    });

    emit('move', {
      fromIndex,
      toIndex: toIndexEdged,
      item: activeItem,
      updated,
    });
  },

  canMonitor(args) {
    return args.source.data.group === props.group;
  },
});

useAutoScroll({
  isEnabled: isDragging,
  element: scrollerEl,
});

function onBeforeItemLeave(input: Element): void {
  const el = input as HTMLElement;
  const rect = el.getBoundingClientRect();

  el.style.position = 'fixed';
  el.style.left = `${rect.left}px`;
  el.style.top = `${rect.top}px`;
  el.style.width = `${rect.width}px`;
  el.style.height = `${rect.height}px`;

  if (el.dataset.tableSubgrid === 'true') {
    el.style.gridTemplateColumns = 'inherit';
    el.style.gridTemplateAreas = 'inherit';
  }
}
</script>

<style scoped>
@layer components {
  .sortable__item--move,
  .sortable__item--enter-active,
  .sortable__item--leave-active {
    transition: all 100ms ease-out;
  }

  .sortable__item--enter-from,
  .sortable__item--leave-to {
    opacity: 0;
  }
}
</style>
