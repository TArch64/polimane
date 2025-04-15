<template>
  <Modal
    title="Створення Схеми"
    save-button="Створити"
    @save="create.call"
  >
    <TextField
      required
      placeholder="Назва"
      variant="control"
      v-model="schema.name"
    />
  </Modal>
</template>

<script setup lang="ts">
import { reactive } from 'vue';
import { Modal, useActiveModal } from '@/components/modal';
import { TextField } from '@/components/form';
import { useAsyncAction } from '@/composables';
import { type ICreateSchemaInput, useSchemasStore } from '../../stores';

const schemasStore = useSchemasStore();
const modal = useActiveModal();

const schema: ICreateSchemaInput = reactive({
  name: '',
});

const create = useAsyncAction(async () => {
  await schemasStore.createSchema(schema);
  modal.close();
});
</script>
