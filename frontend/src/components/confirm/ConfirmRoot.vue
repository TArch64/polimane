<template>
  <Teleport to="body" v-if="openedConfirm">
    <Confirm :key="openedConfirm.id" :model="openedConfirm as ConfirmModel" />
  </Teleport>
</template>

<script setup lang="ts">
import { nextTick, ref, watch } from 'vue';
import { ConfirmPlugin } from './ConfirmPlugin';
import type { Confirm as ConfirmModel } from './Confirm';
import Confirm from './Confirm.vue';

const plugin = ConfirmPlugin.inject();
const openedConfirm = ref<ConfirmModel | null>(null);

watch(() => plugin.openedConfirm?.id, () => {
  document.startViewTransition(() => {
    openedConfirm.value = plugin.openedConfirm ?? null;
    return nextTick();
  });
});
</script>
