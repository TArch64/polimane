<template>
  <ContextMenuTitle class="context-menu__title">
    {{ menu.title }}
  </ContextMenuTitle>

  <template v-for="(item, index) of menu.actions" :key="index">
    <DropdownAction
      :title="item.title"
      :icon="item.icon"
      :danger="item.danger"
      @click="$emit('action', item)"
      v-if="isContextMenuAction(item)"
    />

    <DropdownAction
      :title="item.title"
      :icon="item.icon"
      data-context-menu-group
      @click.stop="openGroup($event, item)"
      v-else
    />
  </template>
</template>

<script setup lang="ts">
import { DropdownAction } from '../dropdown';
import {
  ContextActionModel,
  ContextGroupModel,
  ContextMenuModel,
  isContextMenuAction,
} from './model';
import ContextMenuTitle from './ContextMenuTitle.vue';

defineProps<{
  menu: ContextMenuModel;
}>();

const emit = defineEmits<{
  'action': [action: ContextActionModel];
  'open-group': [group: ContextGroupModel];
}>();

function openGroup(event: Event, group: ContextGroupModel): void {
  const target = event.target as HTMLElement;
  target.style.viewTransitionName = `--context-menu-group-title`;
  emit('open-group', group);
}
</script>

<style scoped>
@layer components {
  .context-menu__title {
    border-bottom: var(--divider);
  }
}
</style>
