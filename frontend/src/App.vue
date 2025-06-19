<template>
  <RouterView />
  <ModalRoot />
  <ConfirmRoot />
  <ContextMenuRoot />
</template>

<script setup lang="ts">
import { watch } from 'vue';
import { RouterView, useRoute } from 'vue-router';
import { ModalRoot } from '@/components/modal';
import { ConfirmRoot } from '@/components/confirm';
import { ContextMenuRoot } from '@/components/contextMenu';

const route = useRoute();

watch(() => route.name, (name, oldName) => {
  if (name) document.scrollingElement!.classList.add(`app--${name}`);
  if (oldName) document.scrollingElement!.classList.remove(`app--${oldName}`);
}, { immediate: true });
</script>
