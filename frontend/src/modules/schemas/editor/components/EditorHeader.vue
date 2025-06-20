<template>
  <Card as="header" class="editor-header">
    <Button icon class="editor-header-back" :to="{ name: 'home' }">
      <ArrowBackIcon />
      Едітор
    </Button>

    <Button icon :disabled="!editorStore.canUndo" @click="editorStore.undo">
      <CornerUpLeftIcon />
    </Button>

    <Button icon :disabled="!editorStore.canRedo" @click="editorStore.redo">
      <CornerUpRightIcon />
    </Button>

    <Dropdown>
      <template #activator="{ open, activatorStyle }">
        <Button
          icon
          :style="mergeAnchorName(activatorStyle, deleteConfirm.anchorStyle)"
          @click="open"
        >
          <MoreHorizontalIcon />
        </Button>
      </template>

      <DropdownAction
        title="Переназвати Схему"
        :icon="EditIcon"
        @click="renameModal.open()"
      />

      <DropdownAction
        danger
        title="Видалити Схему"
        :icon="TrashIcon"
        @click="deleteSchema"
      />
    </Dropdown>
  </Card>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';
import { Button } from '@/components/button';
import {
  ArrowBackIcon,
  CornerUpLeftIcon,
  CornerUpRightIcon,
  EditIcon,
  MoreHorizontalIcon,
  TrashIcon,
} from '@/components/icon';
import { useAsyncAction } from '@/composables';
import { useConfirm } from '@/components/confirm';
import { Dropdown, DropdownAction } from '@/components/dropdown';
import { mergeAnchorName } from '@/helpers';
import { Card } from '@/components/card';
import { useModal } from '@/components/modal';
import { useEditorStore } from '../stores';
import { SchemaRenameModal } from './modals';

const router = useRouter();
const editorStore = useEditorStore();

const renameModal = useModal(SchemaRenameModal);

const deleteConfirm = useConfirm({
  danger: true,
  message: 'Ви впевнені, що хочете видалити цю схему?',
  acceptButton: 'Видалити',
});

const deleteSchema = useAsyncAction(async () => {
  if (await deleteConfirm.ask()) {
    await editorStore.deleteSchema();
    await router.push({ name: 'home' });
  }
});
</script>

<style scoped>
@layer page {
  .editor-header {
    position: fixed;
    top: 8px;
    left: 8px;
    width: 250px;
    z-index: 10;
    display: flex;
    align-items: center;
    padding: 6px 8px;
    gap: 4px;
  }

  .editor-header-back {
    gap: 8px;
    margin-right: auto;
  }
}
</style>
