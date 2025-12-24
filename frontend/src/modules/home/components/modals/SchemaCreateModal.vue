<template>
  <Modal
    title="Нова Схема"
    save-button="Додати"
    :loading="create.isActive"
    @save="create"
  >
    <TextField
      required
      placeholder="Назва"
      variant="control"
      class="schema-create__name"
      v-model="schema.name"
    />

    <RadioSelect
      :options="layoutOptions"
      v-model="schema.layout"
    >
      <template #option-linear>
        <p>
          Підходить для прикрас у вигляді стрічок, прямокутників та інших прикрас без вигинів
        </p>
      </template>

      <template #option-radial>
        <p>
          Підходить для прикрас з розширенням від центру, такий як силянки, кризи, та інші округлі
          прикраси
        </p>
      </template>
    </RadioSelect>
  </Modal>
</template>

<script setup lang="ts">
import { reactive } from 'vue';
import { useRouter } from 'vue-router';
import { Modal, useActiveModal } from '@/components/modal';
import { RadioSelect, type SelectOptions, TextField } from '@/components/form';
import { useAsyncAction } from '@/composables';
import { SchemaLayout } from '@/enums';
import { type ISchemaCreateRequest, useHomeStore } from '../../stores';

const homeStore = useHomeStore();

const router = useRouter();
const modal = useActiveModal();

const schema: ISchemaCreateRequest = reactive({
  name: '',
  layout: SchemaLayout.LINEAR,
});

const layoutOptions: SelectOptions<SchemaLayout> = [
  {
    value: SchemaLayout.LINEAR,
    label: 'Лінійна',
  },
  {
    value: SchemaLayout.RADIAL,
    label: 'Радіальна',
  },
];

const create = useAsyncAction(async () => {
  schema.name = schema.name.trim();
  const created = await homeStore.createSchema.do(schema);

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

<style scoped>
@layer page {
  .schema-create__name {
    margin-bottom: 24px;
  }
}
</style>
