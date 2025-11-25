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
import { EditIcon } from '@/components/icon';
import { useModal } from '@/components/modal';

const props = defineProps<{
  folder: IListFolder;
}>();

const renameModal = useModal(FolderRenameModal);

const folderRoute = computed((): RouteLocationRaw => ({
  name: 'home-folder',
  params: { folderId: props.folder.id },
}));

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
];
</script>
