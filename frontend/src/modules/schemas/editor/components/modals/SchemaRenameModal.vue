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
import { useEditorStore } from '@editor/stores';
import { Modal, useActiveModal } from '@/components/modal';
import { TextField } from '@/components/form';

const editorStore = useEditorStore();
const modal = useActiveModal();

const form = reactive({
  name: editorStore.schema.name,
});

function save(): void {
  editorStore.schema.name = form.name.trim();
  modal.close(null);
}
</script>
