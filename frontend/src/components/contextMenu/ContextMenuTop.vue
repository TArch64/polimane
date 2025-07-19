<template>
  <ContextMenuTitle class="context-menu__title">
    {{ menu.title }}
  </ContextMenuTitle>

  <template v-for="(item, index) of menu.actions" :key="index">
    <DropdownAction
      :title="item.title"
      :icon="item.icon"
      :danger="item.danger"
      :disabled="item.disabled"
      @click="$emit('action', item)"
      v-if="isContextMenuAction(item)"
    />

    <DropdownAction
      :title="item.title"
      :icon="item.icon"
      :disabled="item.disabled"
      data-context-menu-group
      @click.stop="$emit('open-group', item)"
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

defineEmits<{
  'action': [action: ContextActionModel];
  'open-group': [group: ContextGroupModel];
}>();
</script>

<style scoped>
@layer components {
  .context-menu__title {
    border-bottom: var(--divider);
    margin: 0 -4px;
    padding: 8px 4px 10px 12px;
  }
}
</style>
