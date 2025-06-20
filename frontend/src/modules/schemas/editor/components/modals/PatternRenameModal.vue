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
import { set } from '@vueuse/core';
import { Modal, useActiveModal } from '@/components/modal';
import { TextField } from '@/components/form';
import type { ISchemaPattern } from '@/models';

const props = defineProps<{
  pattern: ISchemaPattern;
}>();

const modal = useActiveModal();
const editingName = ref(props.pattern.name);

function save(): void {
  set(props.pattern, 'name', editingName.value.trim());
  modal.close();
}
</script>
