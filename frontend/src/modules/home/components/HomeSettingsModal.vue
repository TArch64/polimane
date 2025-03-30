<template>
  <Modal title="Налаштування">
    <template #activator="ctx">
      <slot v-bind="ctx" />
    </template>

    <Button>
      close
    </Button>
  </Modal>
</template>

<script setup lang="ts">
import { Modal, type ModalActivatorSlot } from '@/components/modal';
import { Button } from '@/components/button';

defineSlots<{
  default: ModalActivatorSlot;
}>();
</script>
