<template>
  <Modal
    title="Створення Схеми"
    save-button="Створити"
    :loading="create.isActive"
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
import { type ICreateSchemaRequest, useSchemasStore } from '@/modules/home/stores';

const router = useRouter();
const schemasStore = useSchemasStore();
const modal = useActiveModal();

const schema: ICreateSchemaRequest = reactive({
  name: '',
});

const create = useAsyncAction(async () => {
  schema.name = schema.name.trim();
  const created = await schemasStore.createSchema(schema);

  modal.close(null, async () => {
    await router.push({
      name: 'schema-editor',
      params: {
        schemaId: created.id,
      },
    });
  });
});
</script>
