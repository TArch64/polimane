<template>
  <DropdownMenu ref="menuRef" class="dropdown" :class="classes">
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
import { computed, onMounted } from 'vue';
import { NodeRect } from '@/models';
import { useDomRef } from '@/composables';
import { DropdownAction, DropdownMenu } from '../dropdown';
import type { ContextMenuModel } from './ContextMenuModel';

const props = defineProps<{
  menu: ContextMenuModel;
}>();

const menuRef = useDomRef();

const classes = computed(() => ({
  'dropdown--initial': !props.menu.menuRect,
}));

onMounted(() => {
  let rect = new NodeRect(menuRef.value!.getBoundingClientRect())
    .delta({ y: -4 });

  if (rect.right >= window.innerWidth - 4) {
    rect = rect.with({ x: window.innerWidth - rect.width - 4 });
  }

  if (rect.bottom >= window.innerHeight) {
    rect = rect.with({ y: window.innerHeight - rect.height });
  }

  // eslint-disable-next-line vue/no-mutating-props
  props.menu.menuRect = rect;
});
</script>

<style scoped>
@layer components {
  .dropdown {
    z-index: 9999;
    margin: 4px 0 0;

    &:not(.dropdown--initial) {
      position-anchor: v-bind("menu.anchorVar");
      position-area: bottom center;
    }
  }

  .dropdown--initial {
    top: v-bind("menu.position.y + 'px'");
    left: v-bind("menu.position.x + 'px'");
  }
}
</style>
