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
import { type ISchemaCreateRequest, useHomeStore } from '../../stores';

const router = useRouter();
const modal = useActiveModal();

const homeStore = useHomeStore();
const createSchema = homeStore.createSchema!;

const schema: ISchemaCreateRequest = reactive({
  name: '',
});

const create = useAsyncAction(async () => {
  schema.name = schema.name.trim();
  const created = await createSchema.do(schema);

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
