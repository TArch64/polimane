<template>
  <Teleport to="body" v-if="openedModal">
    <ModalRender :key="openedModal.id" :modal="openedModal as ModalModel" />
  </Teleport>
</template>

<script setup lang="ts">
import { nextTick, type Ref, ref, watch } from 'vue';
import { useRouteTransition } from '@/composables';
import { ModalPlugin } from './ModalPlugin';
import { ModalRender } from './ModalRender';
import type { ModalModel } from './ModalModel';

const plugin = ModalPlugin.inject();
const routeTransition = useRouteTransition();
const openedModal: Ref<ModalModel | null> = ref(null);

watch(() => plugin.openedModal?.id, () => {
  routeTransition.start(async () => {
    const previousModal = openedModal.value;
    openedModal.value = plugin.openedModal;

    await previousModal?.completeClose();

    // Catch next possible modal
    await nextTick();
    await nextTick();
    await nextTick();
  });
});
</script>
