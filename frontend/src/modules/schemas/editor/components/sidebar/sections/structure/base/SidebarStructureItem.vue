<template>
  <div
    class="sidebar-structure-item"
    :class="classes"
    @click.stop="activeObject.focus.activate(ActiveObjectTrigger.SIDEBAR)"
    @mouseover.stop="activeObject.hover.activate(ActiveObjectTrigger.SIDEBAR)"
  >
    {{ title }}

    <Dropdown>
      <template #activator="{ activatorStyle, open }">
        <Button
          icon
          class="sidebar-structure-item__more-actions"
          :style="mergeAnchorName(activatorStyle, moreActionsButtonStyle)"
          @click="open"
        >
          <MoreHorizontalIcon />
        </Button>
      </template>

      <slot name="actions" />
    </Dropdown>
  </div>

  <slot name="content" v-if="activeObject.focus.isActive" />
</template>

<script setup lang="ts">
import { computed, type Slot } from 'vue';
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
  moreActionsButtonStyle?: Record<string, string> | Record<string, string>[];
}>(), {
  moreActionsButtonStyle: () => ({}),
});

defineSlots<{
  actions: Slot;
  content: Slot;
}>();

const activeObject = useActiveObject(() => props.object);

const classes = computed(() => ({
  'sidebar-structure-item--hover': activeObject.hover.isActive && !activeObject.focus.isActive,
  'sidebar-structure-item--focus': activeObject.focus.isActive,
}));
</script>

<style scoped>
@layer page {
  .sidebar-structure-item {
    font-size: var(--font-sm);
    padding: 4px 8px 4px 12px;
    display: flex;
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
}
</style>
