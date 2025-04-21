<template>
  <template v-for="item in items" :key="item.id">
    <slot :item />
  </template>
</template>

<script setup lang="ts" generic="I extends ISchemaObject">
import { computed, nextTick, onMounted, type Slot, watch } from 'vue';
import { type FabricObject, util } from 'fabric';
import type { ISchemaObject } from '@/models';
import {
  getCanvasObject,
  injectCanvas,
  useObjectParent,
} from '@/modules/schemas/editor/composables';
import type { PositionIterator } from './PositionIterator';

const props = defineProps<{
  items: I[];
  positionIteratorClass: typeof PositionIterator;
}>();

defineSlots<{
  default: Slot<{ item: I }>;
}>();

const canvas = injectCanvas();
const parent = useObjectParent();
const itemIds = computed(() => props.items.map((item) => item.id));

function createPositionIterator(): PositionIterator<FabricObject> {
  const objects = props.items.map((item) => getCanvasObject(item.id));
  // @ts-expect-error: The PositionIterator class is generic, but we are passing a specific type
  return new props.positionIteratorClass(parent, objects);
}

function tryAnimated(object: FabricObject, key: 'top' | 'left', value: number) {
  if (object[key] > 0) {
    object.animate({ [key]: value }, {
      duration: 150,
      onChange: () => canvas.requestRenderAll(),
      easing: util.ease.easeOutQuad,
    });
  } else {
    object.set(key, value);
  }
}

function updatePositions(): void {
  const iterator = createPositionIterator();

  for (const { object, top, left } of iterator) {
    if (object.top !== top) {
      tryAnimated(object, 'top', top);
    }

    if (object.left !== left) {
      tryAnimated(object, 'left', left);
    }
  }

  canvas.requestRenderAll();
}

onMounted(updatePositions);

watch(itemIds, async () => {
  await nextTick();
  updatePositions();
});
</script>
