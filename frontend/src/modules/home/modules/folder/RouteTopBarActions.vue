<template>
  <HomeBarRouteActions>
    <Button icon variant="secondary" @click="openRenameModal">
      <EditIcon />
    </Button>

    <Button icon danger variant="secondary">
      <TrashIcon />
    </Button>
  </HomeBarRouteActions>
</template>

<script setup lang="ts">
import { FolderRenameModal, HomeBarRouteActions } from '@/modules/home/components';
import { Button } from '@/components/button';
import { EditIcon, TrashIcon } from '@/components/icon';
import { useModal } from '@/components/modal';
import { useFolderStore } from './stores';

const folderStore = useFolderStore();

const renameModal = useModal(FolderRenameModal);

const openRenameModal = () => renameModal.open({
  folder: folderStore.folder,
});
</script>
