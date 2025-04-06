<template>
  <Modal title="Створення Схеми">
    <template #activator="ctx">
      <slot v-bind="ctx" />
    </template>

    <TextField
      placeholder="Назва"
      variant="control"
      v-model="schema.name"
    />
  </Modal>
</template>

<script setup lang="ts">
import { reactive } from 'vue';
import { Modal, type ModalActivatorSlot } from '@/components/modal';
import { TextField } from '@/components/form';

defineSlots<{
  default: ModalActivatorSlot;
}>();

const schema = reactive({
  name: '',
});
</script>
