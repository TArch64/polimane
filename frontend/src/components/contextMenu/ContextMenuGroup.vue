<template>
  <div class="context-menu-group__title-row">
    <Button icon @click="$emit('close-group')">
      <ArrowBackIcon />
    </Button>

    <ContextMenuTitle class="context-menu-group__title">
      {{ group.title }}
    </ContextMenuTitle>
  </div>

  <DropdownAction
    v-for="(item, index) of group.actions"
    :key="index"
    :title="item.title"
    :icon="item.icon"
    :danger="item.danger"
    @click="$emit('action', item)"
  />
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { DropdownAction } from '../dropdown';
import { Button } from '../button';
import { ArrowBackIcon } from '../icon';
import type { IContextMenuAction, IContextMenuGroup } from './ContextMenuModel';
import ContextMenuTitle from './ContextMenuTitle.vue';

const props = defineProps<{
  group: IContextMenuGroup;
}>();

defineEmits<{
  'action': [action: IContextMenuAction];
  'close-group': [];
}>();

const titleViewTransitionName = computed(() => `--context-menu-group-title-${props.group.id}`);
</script>

<style scoped>
@layer components {
  .context-menu-group__title-row {
    display: flex;
    align-items: center;
    gap: 4px;
    border-bottom: var(--divider);
    border-bottom-width: 4px;
    margin: 0 -4px;
    padding: 0 4px 4px;
    view-transition-name: v-bind("titleViewTransitionName");
  }

  .context-menu-group__title {
    padding: 0;
  }
}
</style>
