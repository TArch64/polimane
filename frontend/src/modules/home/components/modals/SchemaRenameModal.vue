<template>
  <Modal title="Переназвати Схему" @save="save">
    <TextField
      required
      variant="control"
      placeholder="Назва"
      v-model="form.name"
    />
  </Modal>
</template>

<script setup lang="ts">
import { reactive } from 'vue';
import { Modal, useActiveModal } from '@/components/modal';
import { TextField } from '@/components/form';
import type { ISchema } from '@/models';
import { useSchemasStore } from '../../stores';

const props = defineProps<{
  schema: ISchema;
}>();

const schemasStore = useSchemasStore();
const modal = useActiveModal();

const form = reactive({
  name: props.schema.name,
});

function save(): void {
  schemasStore.updateSchema(props.schema, {
    name: form.name.trim(),
  });

  modal.close(null);
}
</script>
