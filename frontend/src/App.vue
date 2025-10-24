<template>
  <RouterView />
  <ModalRoot />
  <ConfirmRoot />
  <ContextMenuRoot />
</template>

<script setup lang="ts">
import { RouterView, useRoute } from 'vue-router';
import { ref } from 'vue';
import { ModalRoot } from '@/components/modal';
import { ConfirmRoot } from '@/components/confirm';
import { ContextMenuRoot } from '@/components/contextMenu';
import { usePageClass } from '@/composables';
import { provideTopElement } from '@/composables/useTopElement';

const route = useRoute();
usePageClass(() => route.name ? `app--${route.name}` : '');
provideTopElement(ref(document.body));
</script>
