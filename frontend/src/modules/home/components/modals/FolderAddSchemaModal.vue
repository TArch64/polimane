<template>
  <Modal
    title="Додати в директорію"
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
      placeholder="Назва Директорії"
      variant="control"
      class="folder-add-schema-modal__folder-name"
      v-model="form.folderName"
      v-if="isNewFolder"
    />
  </Modal>
</template>

<script setup lang="ts">
import { computed, reactive } from 'vue';
import { Modal, useActiveModal } from '@/components/modal';
import { useFoldersStore } from '@/modules/home/stores';
import { SelectField, TextField } from '@/components/form';
import { useAsyncAction } from '@/composables';

const props = defineProps<{
  schemaIds: string[];
  folderId: string | null;
}>();

const NEW_FOLDER_ID = 'new' as const;

const modal = useActiveModal();

const foldersStore = useFoldersStore();

const form = reactive({
  folderId: props.folderId || NEW_FOLDER_ID,
  folderName: '',
});

const isNewFolder = computed(() => form.folderId === NEW_FOLDER_ID);

const save = useAsyncAction(async () => {
  if (props.folderId === form.folderId) {
    return;
  }

  modal.close(null);
});
</script>

<style scoped>
@layer page {
  .folder-add-schema-modal__folder-name {
    margin-top: 16px;
  }
}
</style>
