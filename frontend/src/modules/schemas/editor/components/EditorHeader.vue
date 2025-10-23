<template>
  <Card as="header" class="editor-header">
    <Button
      icon
      class="editor-header-back"
      :to="{ name: 'home' }"
      :prepend-icon="ArrowBackIcon"
    >
      Едітор
    </Button>

    <template v-if="editorStore.canEdit">
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
        @click="historyStore.undo"
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
        @click="accessEditModal.open()"
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
        @click="deleteSchema"
        v-if="editorStore.canDelete"
      />
    </Dropdown>
  </Card>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';
import { computed, toRef } from 'vue';
import { useHotKeys } from '@editor/composables';
import { Button } from '@/components/button';
import {
  ArrowBackIcon,
  CheckmarkCircleIcon,
  CornerUpLeftIcon,
  CornerUpRightIcon,
  EditIcon,
  FileTextIcon,
  type IconComponent,
  LoaderIcon,
  MoreHorizontalIcon,
  PeopleIcon,
  SaveIcon,
  TrashIcon,
} from '@/components/icon';
import { useAsyncAction, useProgressBar } from '@/composables';
import { useConfirm } from '@/components/confirm';
import { Dropdown, DropdownAction } from '@/components/dropdown';
import { mergeAnchorName } from '@/helpers';
import { Card } from '@/components/card';
import { useModal } from '@/components/modal';
import { useEditorStore, useHistoryStore } from '../stores';
import { AccessEditModal, SchemaExportModal, SchemaRenameModal } from './modals';

const router = useRouter();
const historyStore = useHistoryStore();
const editorStore = useEditorStore();

const renameModal = useModal(SchemaRenameModal);
const exportModal = useModal(SchemaExportModal);
const accessEditModal = useModal(AccessEditModal);

const SavingIcon = computed((): IconComponent => {
  if (editorStore.isSaving) {
    return LoaderIcon;
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

const deleteConfirm = useConfirm({
  danger: true,
  control: false,
  message: 'Ви впевнені, що хочете видалити цю схему?',
  acceptButton: 'Видалити',
});

const deleteSchema = useAsyncAction(async () => {
  if (await deleteConfirm.ask()) {
    await editorStore.deleteSchema();
    await router.push({ name: 'home' });
  }
});

useHotKeys({
  Meta_Z: historyStore.undo,
  Meta_Shift_Z: historyStore.redo,
}, {
  isActive: toRef(editorStore, 'canEdit'),
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

  .editor-header-back {
    margin-right: 40px;
  }
}
</style>
