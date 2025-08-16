<template>
  <RouterView />
  <ModalRoot />
  <ConfirmRoot />
  <ContextMenuRoot />
  <AppLoader v-if="isLoaderDisplaying" />
</template>

<script setup lang="ts">
import { RouterView, useRoute } from 'vue-router';
import { nextTick, ref, watch } from 'vue';
import { ModalRoot } from './components/modal';
import { ConfirmRoot } from './components/confirm';
import { ContextMenuRoot } from './components/contextMenu';
import { usePageClass, useRouteTransition } from './composables';
import { useLoaderStore } from './stores';
import AppLoader from './AppLoader.vue';

const loaderStore = useLoaderStore();
const route = useRoute();
const routeTransition = useRouteTransition();

const isLoaderDisplaying = ref(false);

usePageClass(() => route.name ? `app--${route.name}` : '');

watch(() => loaderStore.isDisplaying, (isDisplaying) => {
  routeTransition.start(() => {
    isLoaderDisplaying.value = isDisplaying;
    return nextTick();
  });
});
</script>
