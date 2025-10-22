<template>
  <DropdownMenu
    ref="menuRef"
    class="dropdown-menu"
    view-transition-name="context-menu"
    :control="menu.control"
    :class="classes"
    v-popover-shift
  >
    <ContextMenuGroup
      v-if="menu.openedGroup"
      :group="menu.openedGroup"
      @action="executeAction"
      @close-group="closeGroup"
    />

    <ContextMenuTop
      :menu
      @action="executeAction"
      @open-group="openGroup"
      v-else
    />
  </DropdownMenu>
</template>

<script setup lang="ts">
import { computed, nextTick, onMounted } from 'vue';
import { onClickOutside } from '@vueuse/core';
import { NodeRect } from '@/models';
import { useDomRef, useRouteTransition } from '@/composables';
import { vPopoverShift } from '@/directives';
import { DropdownMenu } from '../dropdown';
import type { ContextActionModel, ContextGroupModel, ContextMenuModel } from './model';
import ContextMenuTop from './ContextMenuTop.vue';
import ContextMenuGroup from './ContextMenuGroup.vue';

const props = defineProps<{
  menu: ContextMenuModel;
}>();

const emit = defineEmits<{
  close: [];
}>();

const routeTransition = useRouteTransition();
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

  props.menu.setMenuRect(rect);
});

function executeAction(action: ContextActionModel): void {
  emit('close');
  props.menu.executeAction(action);
}

function openGroup(group: ContextGroupModel): void {
  routeTransition.start(() => {
    props.menu.openGroup(group);
    return nextTick();
  });
}

function closeGroup(): void {
  routeTransition.start(() => {
    props.menu.closeGroup();
    return nextTick();
  });
}

onClickOutside(menuRef, () => emit('close'));
</script>

<style scoped>
@layer components {
  .dropdown-menu {
    z-index: 9999;
    margin: 4px 0 0;
    min-width: 200px;

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
