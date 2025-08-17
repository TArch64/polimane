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
    :disabled="item.disabled"
    @click="$emit('action', item)"
  />
</template>

<script setup lang="ts">
import { DropdownAction } from '../dropdown';
import { Button } from '../button';
import { ArrowBackIcon } from '../icon';
import { ContextActionModel, type ContextGroupModel } from './model';
import ContextMenuTitle from './ContextMenuTitle.vue';

defineProps<{
  group: ContextGroupModel;
}>();

defineEmits<{
  'action': [action: ContextActionModel];
  'close-group': [];
}>();
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
  }

  .context-menu-group__title {
    padding: 0;
  }
}
</style>
