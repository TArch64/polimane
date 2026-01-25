<template>
  <HomeListCard
    :to="folderRoute"
    :menu-title="folder.name"
    :menuActions
  >
    <HomeListScreenshot
      :path="folder.screenshotPath"
      :alt="`Скріншот схеми ${folder.name}`"
      :background-color="folder.backgroundColor || DEFAULT_SCHEMA_BACKGROUND"
    />

    {{ folder.name }}
  </HomeListCard>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import type { RouteLocationRaw } from 'vue-router';
import type { IListFolder } from '@/modules/home/stores';
import { DEFAULT_SCHEMA_BACKGROUND } from '@/config';
import { FolderRenameModal, HomeListCard, HomeListScreenshot } from '@/modules/home/components';
import type { MaybeContextMenuAction } from '@/components/contextMenu';
import { EditIcon, TrashIcon } from '@/components/icon';
import { useModal } from '@/components/modal';
import { useConfirm } from '@/components/confirm';
import { useAsyncAction, useProgressBar } from '@/composables';
import { useFoldersStore } from '../stores';

const props = defineProps<{
  folder: IListFolder;
}>();

const foldersStore = useFoldersStore();

const renameModal = useModal(FolderRenameModal);

const folderRoute = computed((): RouteLocationRaw => ({
  name: 'home-folder',
  params: { folderId: props.folder.id },
}));

const deleteConfirm = useConfirm({
  danger: true,
  message: 'Ви впевнені, що хочете видалити цю директорію?',
  acceptButton: 'Видалити',
  additionalCondition: 'Видалити всі схеми в директорії',
});

const deleteFolder = useAsyncAction(async (deleteSchemas: boolean) => {
  await foldersStore.deleteFolder(props.folder, {
    deleteSchemas,
  });
});

useProgressBar(deleteFolder);

const menuActions: MaybeContextMenuAction[] = [
  {
    title: 'Змінити Назву',
    icon: EditIcon,

    onAction() {
      renameModal.open({
        folder: props.folder,
      });
    },
  },
  {
    title: 'Видалити',
    icon: TrashIcon,
    danger: true,

    async onAction(event) {
      const confirmation = await deleteConfirm.ask({
        virtualTarget: event.menuRect,
      });

      if (confirmation.isAccepted) {
        await deleteFolder(confirmation.isSecondaryAccepted ?? false);
      }
    },
  },
];
</script>
