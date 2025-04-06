<template>
  <Modal
    ref="modalRef"
    title="Створення Схеми"
    save-button="Створити"
    @save="create.call"
  >
    <template #activator="ctx">
      <slot v-bind="ctx" />
    </template>

    <TextField
      required
      placeholder="Назва"
      variant="control"
      v-model="schema.name"
    />
  </Modal>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue';
import type { ComponentExposed } from 'vue-component-type-helpers';
import { Modal, type ModalActivatorSlot } from '@/components/modal';
import { TextField } from '@/components/form';
import { useAsyncAction } from '@/composables';
import { type ICreateSchemaInput, useSchemasStore } from '../../stores';

defineSlots<{
  default: ModalActivatorSlot;
}>();

const modalRef = ref<ComponentExposed<typeof Modal>>(null!);
const schemasStore = useSchemasStore();

const schema: ICreateSchemaInput = reactive({
  name: '',
});

const create = useAsyncAction(async () => {
  await schemasStore.createSchema(schema);
  modalRef.value.close();
});
</script>
