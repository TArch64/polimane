<template>
  <Modal title="Додати Схему" save-button="Додати">
    <template #activator="ctx">
      <slot v-bind="ctx" />
    </template>
  </Modal>
</template>

<script setup lang="ts">
import { Modal, type ModalActivatorSlot } from '@/components/modal';

defineSlots<{
  default: ModalActivatorSlot;
}>();
</script>
