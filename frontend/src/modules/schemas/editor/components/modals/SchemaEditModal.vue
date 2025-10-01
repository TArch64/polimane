<template>
  <Modal title="Налаштування Схеми" @save="save">
    <TextField
      required
      class="form-field"
      variant="control"
      placeholder="Назва"
      v-model="form.name"
    />

    <ColorPicker
      label="Колір Фону"
      v-model="form.backgroundColor"
    />
  </Modal>
</template>

<script setup lang="ts">
import { reactive } from 'vue';
import { Modal, useActiveModal } from '@/components/modal';
import { ColorPicker, TextField } from '@/components/form';
import { useEditorStore } from '@/modules/schemas/editor/stores';

const editorStore = useEditorStore();
const modal = useActiveModal();

const form = reactive({
  name: editorStore.schema.name,
  backgroundColor: editorStore.schema.backgroundColor,
});

function save(): void {
  editorStore.schema.name = form.name.trim();
  editorStore.schema.backgroundColor = form.backgroundColor;
  modal.close(null);
}
</script>

<style scoped>
@layer page {
  .form-field {
    margin-bottom: 16px;
  }
}
</style>
