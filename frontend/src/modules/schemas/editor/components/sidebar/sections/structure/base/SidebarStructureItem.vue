<template>
  <div
    ref="rootRef"
    class="sidebar-structure-item"
    :class="classes"
    @click.stop="activeObject.focus.toggle(ActiveObjectTrigger.SIDEBAR)"
    @mouseover.stop="activeObject.hover.activate(ActiveObjectTrigger.SIDEBAR)"
  >
    {{ title }}

    <Dropdown v-if="slots.actions">
      <template #activator="{ activatorStyle, open }">
        <Button
          icon
          class="sidebar-structure-item__more-actions"
          :style="mergeAnchorName(activatorStyle, moreActionsButtonStyle)"
          @click.stop="open"
        >
          <MoreHorizontalIcon />
        </Button>
      </template>

      <slot name="actions" />
    </Dropdown>
  </div>

  <Transition name="sidebar-structure-item__content-" :duration="150">
    <div class="sidebar-structure-item__content" v-if="activeObject.focus.isActive">
      <slot name="content" />
    </div>
  </Transition>
</template>

<script setup lang="ts">
import { computed, ref, type Slot } from 'vue';
import { MoreHorizontalIcon } from '@/components/icon';
import { Dropdown } from '@/components/dropdown';
import { Button } from '@/components/button';
import { mergeAnchorName } from '@/helpers';
import type { ISchemaObject } from '@/models';
import { ActiveObjectTrigger } from '@/modules/schemas/editor/stores';
import { useActiveObject } from '@/modules/schemas/editor/composables';

const props = withDefaults(defineProps<{
  object: ISchemaObject;
  title: string;
  depth: number;
  moreActionsButtonStyle?: Record<string, string> | Record<string, string>[];
}>(), {
  moreActionsButtonStyle: () => ({}),
});

const slots = defineSlots<{
  actions?: Slot;
  content?: Slot;
}>();

const rootRef = ref<HTMLElement>(null!);
const activeObject = useActiveObject(() => props.object);

const classes = computed(() => ({
  'sidebar-structure-item--hover': activeObject.hover.isExactActive && !activeObject.focus.isExactActive,
  'sidebar-structure-item--focus': activeObject.focus.isExactActive,
  'sidebar-structure-item--root': props.depth === 0,
}));

const leftPadding = computed(() => `${12 + props.depth * 8}px`);

activeObject.focus.onExactActive((trigger) => {
  if (trigger !== ActiveObjectTrigger.SIDEBAR) {
    rootRef.value.scrollIntoView({
      behavior: 'smooth',
    });
  }
});
</script>

<style scoped>
@layer page {
  .sidebar-structure-item {
    font-size: var(--font-sm);
    padding: 4px 8px 4px v-bind('leftPadding');
    display: flex;
    min-height: 40px;
    justify-content: space-between;
    align-items: center;
    background-color: var(--color-background-1);
    transition: background-color 0.15s ease-out;
    will-change: background-color;
  }

  .sidebar-structure-item--hover {
    background-color: var(--color-background-2);
  }

  .sidebar-structure-item--focus {
    background-color: var(--color-background-3);
  }

  .sidebar-structure-item--hover,
  .sidebar-structure-item--focus {
    .sidebar-structure-item__more-actions {
      opacity: 1;
    }
  }

  .sidebar-structure-item__more-actions {
    opacity: 0;
    transition: opacity 0.15s ease-out;
  }

  .sidebar-structure-item__content {
    position: relative;

    &::before {
      content: "";
      position: absolute;
      pointer-events: none;
      z-index: 1;
      left: v-bind('leftPadding');
      border-left: var(--divider);
      width: 0;
      height: 100%;
    }
  }

  .sidebar-structure-item__content--enter-active,
  .sidebar-structure-item__content--leave-active {
    overflow: hidden;
    transition: height ease-out 0.15s, opacity ease-out 0.15s;
  }

  .sidebar-structure-item__content--enter-from,
  .sidebar-structure-item__content--leave-to {
    height: 0;
    opacity: 0;
  }
}
</style>
