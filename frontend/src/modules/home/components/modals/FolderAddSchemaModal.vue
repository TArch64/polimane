<template>
  <Modal
    title="Перемістити в директорію"
    :loading="save.isActive"
    @save="save"
  >
    <SelectField
      label
      placeholder="Директорія"
      variant="control"
      v-model="form.folderId"
    >
      <option :value="NEW_FOLDER_ID">
        Нова Директорія
      </option>

      <option
        v-for="folder of foldersStore.folders"
        :key="folder.id"
        :value="folder.id"
      >
        {{ folder.name }}
      </option>
    </SelectField>

    <TextField
      label
      required
      ref="nameRef"
      placeholder="Назва Директорії"
      variant="control"
      class="folder-add-schema-modal__folder-name"
      v-model="form.folderName"
      v-if="isNewFolder"
    />
  </Modal>
</template>

<script setup lang="ts">
import { computed, reactive, ref } from 'vue';
import type { ComponentExposed } from 'vue-component-type-helpers';
import { Modal, useActiveModal } from '@/components/modal';
import { SelectField, TextField } from '@/components/form';
import { HttpError, HttpErrorReason, useAsyncAction } from '@/composables';
import { useHomeFoldersStore, useHomeStore } from '../../stores';

const props = defineProps<{
  schemaIds: string[];
  folderId: string | null;
}>();

const NEW_FOLDER_ID = 'new' as const;

const nameRef = ref<ComponentExposed<typeof TextField> | null>(null);
const modal = useActiveModal();

const homeStore = useHomeStore();
const foldersStore = useHomeFoldersStore();

const form = reactive({
  folderId: props.folderId || NEW_FOLDER_ID,
  folderName: '',
});

const isNewFolder = computed(() => form.folderId === NEW_FOLDER_ID);

const save = useAsyncAction(async () => {
  if (props.folderId === form.folderId) {
    return;
  }

  try {
    await homeStore.addSchemaToFolder.do({
      schemaIds: props.schemaIds,
      folderId: isNewFolder.value ? null : form.folderId,
      folderName: isNewFolder.value ? form.folderName : null,
    });

    modal.close(null);
  } catch (error) {
    if (HttpError.isReason(error, HttpErrorReason.NAME_ALREADY_IN_USE)) {
      nameRef.value?.setError('Ця назва вже використовується');
    }
  }
});
</script>

<style scoped>
@layer page {
  .folder-add-schema-modal__folder-name {
    margin-top: 16px;
  }
}
</style>
