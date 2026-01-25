<template>
  <HomeBarRouteActions>
    <HomeBarCreateSchema />

    <Button icon variant="secondary" @click="openRenameModal">
      <EditIcon />
    </Button>

    <Button
      icon
      danger
      variant="secondary"
      :loading="deleteFolder.isActive"
      :style="deleteConfirm.anchorStyle"
      @click="deleteIntent"
    >
      <TrashIcon />
    </Button>
  </HomeBarRouteActions>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';
import {
  FolderRenameModal,
  HomeBarCreateSchema,
  HomeBarRouteActions,
} from '@/modules/home/components';
import { Button } from '@/components/button';
import { EditIcon, TrashIcon } from '@/components/icon';
import { useModal } from '@/components/modal';
import { useConfirm } from '@/components/confirm';
import { useAsyncAction, useProgressBar } from '@/composables';
import { useFolderStore } from './stores';

const folderStore = useFolderStore();

const router = useRouter();
const renameModal = useModal(FolderRenameModal);

const deleteConfirm = useConfirm({
  danger: true,
  message: 'Ви впевнені, що хочете видалити цю директорію?',
  acceptButton: 'Видалити',
  additionalCondition: 'Видалити всі схеми в директорії',
});

const openRenameModal = () => renameModal.open({
  folder: folderStore.folder,
});

const deleteFolder = useAsyncAction(async (deleteSchemas: boolean) => {
  await folderStore.delete({ deleteSchemas });
  await router.push({ name: 'home' });
});

useProgressBar(deleteFolder);

async function deleteIntent(): Promise<void> {
  const confirmation = await deleteConfirm.ask();

  if (confirmation.isAccepted) {
    await deleteFolder(confirmation.isSecondaryAccepted ?? false);
  }
}
</script>
