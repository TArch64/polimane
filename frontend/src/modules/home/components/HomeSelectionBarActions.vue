<template>
  <template v-if="adminActionIds.size">
    <Button
      icon
      size="md"
      variant="secondary"
      title="Редагувати доступ"
      @click="openAccessEditModal"
    >
      <PeopleIcon />
    </Button>

    <Button
      icon
      size="md"
      variant="secondary"
      title="Видалити обрані схеми"
      :style="deleteConfirm.anchorStyle"
      @click="deleteIntent"
    >
      <TrashIcon />
    </Button>
  </template>

  <p v-else>
    Немає доступних дій для обраних схем
  </p>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { Button } from '@/components/button';
import { PeopleIcon, TrashIcon } from '@/components/icon';
import { useConfirm } from '@/components/confirm';
import { useSchemasStore } from '@/modules/home/stores';
import { useAsyncAction } from '@/composables';
import { useModal } from '@/components/modal';
import {
  SchemaAccessEditModal,
  useSchemaUsersStore,
} from '@/modules/schemas/shared/modals/accessEdit';
import { AccessLevel } from '@/enums';

const schemasStore = useSchemasStore();
const schemaUsersStore = useSchemaUsersStore();

const accessEditModal = useModal(SchemaAccessEditModal);

const adminActionIds = computed(() => {
  return schemasStore.filterIdsByAccess(schemasStore.selected, AccessLevel.ADMIN);
});

const deleteConfirm = useConfirm({
  danger: true,
  control: false,
  message: () => `Ви впевнені, що хочете видалити ${adminActionIds.value.size} схеми?`,
  acceptButton: 'Видалити',
});

const deleteSchemas = useAsyncAction(async () => {
  await schemasStore.deleteMany(adminActionIds.value);
  schemasStore.clearSelection();
});

async function deleteIntent(): Promise<void> {
  if (await deleteConfirm.ask()) {
    await deleteSchemas();
  }
}

async function openAccessEditModal(): Promise<void> {
  await schemaUsersStore.load([...adminActionIds.value]);
  accessEditModal.open();
}
</script>
