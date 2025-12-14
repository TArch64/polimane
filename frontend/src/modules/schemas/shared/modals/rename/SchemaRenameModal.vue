<template>
  <Modal
    title="Змінити Назву Схеми"
    :loading="save.isActive"
    @save="save"
  >
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
import type { ISchema, SchemaUpdate } from '@/models';
import type { MaybePromise } from '@/types';
import { useAsyncAction } from '@/composables';

const props = defineProps<{
  schema: ISchema;
  updateSchema: (attrs: SchemaUpdate) => MaybePromise<void>;
}>();

const modal = useActiveModal();

const form = reactive({
  name: props.schema.name,
});

const save = useAsyncAction(async () => {
  form.name = form.name.trim();

  if (form.name === props.schema.name) {
    return modal.close(null);
  }

  await props.updateSchema({
    name: form.name.trim(),
  });

  modal.close(null);
});
</script>
