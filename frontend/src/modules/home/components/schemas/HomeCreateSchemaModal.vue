<template>
  <Modal
    title="Створення Схеми"
    save-button="Створити"
    @save="create"
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
import { useRouter } from 'vue-router';
import { Modal, useActiveModal } from '@/components/modal';
import { TextField } from '@/components/form';
import { useAsyncAction } from '@/composables';
import { type ICreateSchemaInput, useSchemasStore } from '@/modules/home/stores';

const router = useRouter();
const schemasStore = useSchemasStore();
const modal = useActiveModal();

const schema: ICreateSchemaInput = reactive({
  name: '',
});

const create = useAsyncAction(async () => {
  const created = await schemasStore.createSchema(schema);

  modal.close(async () => {
    await router.push({
      name: 'schema-editor',
      params: {
        schemaId: created.id,
      },
    });
  });
});
</script>
