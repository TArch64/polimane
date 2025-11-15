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
      v-model="destinationFolderId"
    >
      <option
        v-for="folder of foldersStore.folders"
        :key="folder.id"
        :value="folder.id"
      >
        {{ folder.name }}
      </option>
    </SelectField>
  </Modal>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { Modal, useActiveModal } from '@/components/modal';
import { useFoldersStore } from '@/modules/home/stores';
import { SelectField } from '@/components/form';
import { useAsyncAction } from '@/composables';

const props = defineProps<{
  schemaIds: string[];
  folderId: string | null;
}>();

const modal = useActiveModal();

const foldersStore = useFoldersStore();

const destinationFolderId = ref<string | null>(props.folderId);

const save = useAsyncAction(async () => {
  if (props.folderId === destinationFolderId.value) {
    return;
  }

  modal.close(null);
});
</script>
