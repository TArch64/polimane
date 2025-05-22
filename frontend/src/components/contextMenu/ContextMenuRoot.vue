<template>
  <Teleport to="body" v-if="activeMenu">
    <ContextMenu :key="activeMenu.id" :menu="activeMenu" />
  </Teleport>
</template>

<script setup lang="ts">
import { nextTick, shallowRef, watch } from 'vue';
import { useRouteTransition } from '@/composables';
import ContextMenu from './ContextMenu.vue';
import { ContextMenuPlugin } from './ContextMenuPlugin';
import { ContextMenu as MenuModel } from './ContextMenu';

const plugin = ContextMenuPlugin.inject();
const routeTransition = useRouteTransition();

const activeMenu = shallowRef<MenuModel | null>(null);

watch(plugin.activeMenu, (menu) => {
  routeTransition.start(async () => {
    activeMenu.value = menu;
    return nextTick();
  });
});
</script>
