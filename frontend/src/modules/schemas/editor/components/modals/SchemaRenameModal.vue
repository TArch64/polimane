<template>
  <Modal title="Перезназвати Схему" @save="save">
    <TextField
      variant="control"
      placeholder="Назва Схеми"
      v-model="editingName"
    />
  </Modal>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { Modal, useActiveModal } from '@/components/modal';
import { TextField } from '@/components/form';
import { useEditorStore } from '@/modules/schemas/editor/stores';

const editorStore = useEditorStore();
const modal = useActiveModal();
const editingName = ref(editorStore.schema.name);

function save(): void {
  editorStore.schema.name = editingName.value.trim();
  modal.close(null);
}
</script>
