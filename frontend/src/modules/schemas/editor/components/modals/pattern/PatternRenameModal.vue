<template>
  <Modal title="Перезназвати Паттерн" @save="save">
    <TextField
      variant="control"
      placeholder="Назва Паттерна"
      v-model="editingName"
    />
  </Modal>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { Modal, useActiveModal } from '@/components/modal';
import { TextField } from '@/components/form';
import type { ISchemaPattern } from '@/models';
import { usePatternsStore } from '@/modules/schemas/editor/stores';

const props = defineProps<{
  pattern: ISchemaPattern;
}>();

const patternsStore = usePatternsStore();
const modal = useActiveModal();
const editingName = ref(props.pattern.name);

function save(): void {
  modal.close(null, () => {
    patternsStore.updatePattern(props.pattern, {
      name: editingName.value.trim(),
    });
  });
}
</script>
