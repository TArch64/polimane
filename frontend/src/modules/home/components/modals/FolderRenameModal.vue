<template>
  <Modal title="Змінити Назву Директорії" @save="save">
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
import type { IFolder } from '@/models';
import { useHomeStore } from '@/modules/home/stores';

const props = defineProps<{
  folder: IFolder;
}>();

const homeStore = useHomeStore();

const modal = useActiveModal();

const form = reactive({
  name: props.folder.name,
});

function save(): void {
  homeStore.updateFolder.do(props.folder, {
    name: form.name.trim(),
  });

  modal.close(null);
}
</script>
