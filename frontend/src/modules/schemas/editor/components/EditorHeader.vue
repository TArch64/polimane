<template>
  <Card as="header" class="editor-header">
    <Button
      icon
      class="editor-header__back"
      :to="{ name: 'home' }"
      :prepend-icon="ArrowBackIcon"
    >
      Едітор
    </Button>

    <template v-if="editorStore.canEdit && !isMobile">
      <Button
        icon
        :disabled="isSaveDisabled"
        :title="savingTitle"
        @click="editorStore.save"
      >
        <SavingIcon />
      </Button>

      <Button
        icon
        :disabled="!historyStore.canUndo"
        title="Відмінити зміни"
        @click="undo"
      >
        <CornerUpLeftIcon />
      </Button>

      <Button
        icon
        :disabled="!historyStore.canRedo"
        title="Повернути назад зміни"
        @click="historyStore.redo"
      >
        <CornerUpRightIcon />
      </Button>
    </template>

    <Dropdown>
      <template #activator="{ open, isOpened, activatorStyle }">
        <Button
          icon
          :active="isOpened"
          :style="mergeAnchorName(activatorStyle, deleteConfirm.anchorStyle)"
          @click="open"
        >
          <MoreHorizontalIcon />
        </Button>
      </template>

      <DropdownAction
        title="Переназвати"
        :icon="EditIcon"
        @click="renameModal.open()"
        v-if="editorStore.canEdit"
      />

      <DropdownAction
        title="Редагувати Доступ"
        :icon="PeopleIcon"
        :disabled="openAccessEditModal.isActive"
        @click="openAccessEditModal"
        v-if="editorStore.canEditAccess"
      />

      <DropdownAction
        title="Зберегти як PDF"
        :icon="FileTextIcon"
        @click="exportModal.open()"
      />

      <DropdownAction
        danger
        title="Видалити"
        :icon="TrashIcon"
        @click="deleteSchemaIntent"
        v-if="editorStore.canDelete"
      />
    </Dropdown>
  </Card>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';
import { type Component, computed } from 'vue';
import { useHotKeys } from '@editor/composables';
import { Button } from '@/components/button';
import {
  ArrowBackIcon,
  CheckmarkCircleIcon,
  CornerUpLeftIcon,
  CornerUpRightIcon,
  EditIcon,
  FileTextIcon,
  MoreHorizontalIcon,
  PeopleIcon,
  SaveIcon,
  TrashIcon,
} from '@/components/icon';
import Spinner from '@/components/Spinner.vue';
import { useAsyncAction, useMobileScreen, useProgressBar } from '@/composables';
import { useConfirm } from '@/components/confirm';
import { Dropdown, DropdownAction } from '@/components/dropdown';
import { mergeAnchorName } from '@/helpers';
import { Card } from '@/components/card';
import { useModal } from '@/components/modal';
import { useEditorStore, useHistoryStore, useSchemaUsersStore, useSelectionStore } from '../stores';
import { AccessEditModal, SchemaExportModal, SchemaRenameModal } from './modals';

const router = useRouter();
const historyStore = useHistoryStore();
const editorStore = useEditorStore();
const schemaUsersStore = useSchemaUsersStore();
const selectionStore = useSelectionStore();

const isMobile = useMobileScreen();

const renameModal = useModal(SchemaRenameModal);
const exportModal = useModal(SchemaExportModal);
const accessEditModal = useModal(AccessEditModal);

const SavingIcon = computed((): Component => {
  if (editorStore.isSaving) {
    return Spinner;
  }
  if (editorStore.hasUnsavedChanges) {
    return SaveIcon;
  }
  return CheckmarkCircleIcon;
});

const savingTitle = computed((): string => {
  if (editorStore.isSaving) {
    return 'Зміни зберігаються прямо зараз';
  }
  if (editorStore.hasUnsavedChanges) {
    return 'Є незбережені зміни';
  }
  return 'Всі зміни збережено';
});

const isSaveDisabled = computed(() => {
  return !editorStore.hasUnsavedChanges || editorStore.isSaving;
});

const openAccessEditModal = useAsyncAction(async () => {
  await schemaUsersStore.load();
  accessEditModal.open();
});

const deleteConfirm = useConfirm({
  danger: true,
  control: false,
  message: 'Ви впевнені, що хочете видалити цю схему?',
  acceptButton: 'Видалити',
});

const deleteSchema = useAsyncAction(async () => {
  await editorStore.deleteSchema();
  await router.push({ name: 'home' });
});

async function deleteSchemaIntent() {
  if (await deleteConfirm.ask()) {
    await deleteSchema();
  }
}

function undo() {
  selectionStore.reset();
  historyStore.undo();
}

useHotKeys({
  Meta_KeyZ: undo,
  Meta_Shift_KeyZ: historyStore.redo,
}, {
  isActive: () => editorStore.canEdit && !isMobile.value,
});

useProgressBar(deleteSchema);
</script>

<style scoped>
@layer page {
  .editor-header {
    position: fixed;
    top: var(--editor-ui-padding);
    left: var(--editor-ui-padding);
    z-index: 10;
    display: flex;
    align-items: center;
    padding: 6px 8px;
    gap: 4px;
  }

  .editor-header__back {
    margin-right: 40px;
  }
}
</style>
