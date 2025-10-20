<template>
  <Teleport to="body" :key="activeMenu.id" v-if="activeMenu">
    <VirtualTarget :menu="activeMenu" v-if="activeMenu.menuRect" />
    <ContextMenu :menu="activeMenu" @close="plugin.hide()" />
  </Teleport>
</template>

<script setup lang="ts">
import { type FunctionalComponent, h, nextTick, shallowRef, watch } from 'vue';
import { useRouteTransition } from '@/composables';
import ContextMenu from './ContextMenu.vue';
import { ContextMenuPlugin } from './ContextMenuPlugin';
import { ContextMenuModel } from './model';

const plugin = ContextMenuPlugin.inject();
const routeTransition = useRouteTransition();

const activeMenu = shallowRef<ContextMenuModel | null>(null);

watch(plugin.activeMenu, (menu) => {
  routeTransition.start(async () => {
    activeMenu.value = menu;
    return nextTick();
  });
});

const VirtualTarget: FunctionalComponent<{ menu: ContextMenuModel }> = (props) => h('div', {
  inert: true,

  style: {
    position: 'fixed',
    anchorName: props.menu.anchorVar,
    top: `${props.menu.menuRect!.top}px`,
    left: `${props.menu.menuRect!.left}px`,
    width: `${props.menu.menuRect!.width}px`,
  },
});

VirtualTarget.displayName = 'VirtualTarget';
</script>
