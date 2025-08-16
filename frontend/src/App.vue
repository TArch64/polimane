<template>
  <RouterView />
  <ModalRoot />
  <ConfirmRoot />
  <ContextMenuRoot />
</template>

<script setup lang="ts">
import { RouterView, useRoute } from 'vue-router';
import { ModalRoot } from '@/components/modal';
import { ConfirmRoot } from '@/components/confirm';
import { ContextMenuRoot } from '@/components/contextMenu';
import { usePageClass } from '@/composables';

const route = useRoute();
usePageClass(() => route.name ? `app--${route.name}` : '');
</script>
