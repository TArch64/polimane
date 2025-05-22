<template>
  <DropdownMenu ref="menuRef" class="dropdown">
    <DropdownAction
      v-for="(action, index) of menu.actions"
      :key="`${action.title} ${index}`"
      :title="action.title"
      :icon="action.icon"
      :danger="action.danger"
      @click="menu.executeAction(action)"
    />
  </DropdownMenu>
</template>

<script setup lang="ts">
import { onMounted } from 'vue';
import { NodeRect } from '@/models';
import { useDomRef } from '@/composables';
import { DropdownAction, DropdownMenu } from '../dropdown';
import type { ContextMenu } from './ContextMenu';

const props = defineProps<{
  menu: ContextMenu;
}>();

const menuRef = useDomRef();

onMounted(() => {
  const rect = new NodeRect(menuRef.value!.getBoundingClientRect());
  props.menu.setMenuRect(rect.with({ y: rect.y - 4 }));
});
</script>

<style scoped>
@layer components {
  .dropdown {
    position: fixed;
    z-index: 9999;
    margin: 4px 0 0;
    top: v-bind("menu.position.y + 'px'");
    left: v-bind("menu.position.x + 'px'");
  }
}
</style>
