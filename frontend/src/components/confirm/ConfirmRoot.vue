<template>
  <Teleport to="body" v-if="openedConfirm">
    <Confirm :key="openedConfirm.id" :model="openedConfirm as ConfirmModel" />
  </Teleport>
</template>

<script setup lang="ts">
import { nextTick, ref, watch } from 'vue';
import { useRouteTransition } from '@/composables';
import { ConfirmPlugin } from './ConfirmPlugin';
import type { Confirm as ConfirmModel } from './Confirm';
import Confirm from './Confirm.vue';

const plugin = ConfirmPlugin.inject();
const routeTransition = useRouteTransition();
const openedConfirm = ref<ConfirmModel | null>(null);

watch(() => plugin.openedConfirm?.id, () => {
  routeTransition.start(async () => {
    openedConfirm.value = plugin.openedConfirm ?? null;
    await nextTick();
  });
});
</script>
