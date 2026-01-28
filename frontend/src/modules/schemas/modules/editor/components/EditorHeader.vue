<template>
  <Card as="header" class="editor-header">
    <Button
      class="editor-header__back"
      :to="backLocation"
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
        :title="undoTooltip"
        @click="undo"
      >
        <CornerUpLeftIcon />
      </Button>

      <Button
        icon
        :disabled="!historyStore.canRedo"
        :title="redoTooltip"
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
          title="Змінити назву"
        :icon="EditIcon"
        @click="openRenameModal"
          v-if="editorStore.canEditName"
      />

      <DropdownAction
        title="Редагувати Доступ"
        :icon="editAccessIcon"
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
import { type RouteLocationRaw, useRoute, useRouter } from 'vue-router';
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
  LockIcon,
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
import { SchemaRenameModal } from '@/modules/schemas/shared/modals/rename';
import { SchemaAccessEditModal } from '@/modules/schemas/shared/modals/accessEdit';
import { useSessionStore } from '@/stores';
import { SubscriptionLimit } from '@/enums';
import { UpgradePlanModal } from '@/components/subscription';
import { useEditorStore, useHistoryStore, useSelectionStore } from '../stores';
import { SchemaExportModal } from './modals';

const router = useRouter();
const route = useRoute<'schema-editor'>();

const sessionStore = useSessionStore();
const historyStore = useHistoryStore();
const editorStore = useEditorStore();
const selectionStore = useSelectionStore();

const isMobile = useMobileScreen();

const renameModal = useModal(SchemaRenameModal);
const exportModal = useModal(SchemaExportModal);
const accessEditModal = useModal(SchemaAccessEditModal);
const upgradePlanModal = useModal(UpgradePlanModal);

const sharedAccessLimit = computed(() => sessionStore.getLimit(SubscriptionLimit.SHARED_ACCESS));
const isSharedAccessAvailable = computed(() => sharedAccessLimit.value! > 1);
const editAccessIcon = computed(() => isSharedAccessAvailable.value ? PeopleIcon : LockIcon);

const backLocation = computed((): RouteLocationRaw => {
  const from = route.query.from as string | undefined;
  return from || { name: 'home' };
});

const SavingIcon = computed((): Component => {
  if (editorStore.isSaving) {
    return Spinner;
  }
  if (editorStore.hasUnsavedChanges) {
    return SaveIcon;
  }
  return CheckmarkCircleIcon;
});

const isSaveDisabled = computed(() => {
  return !editorStore.hasUnsavedChanges || editorStore.isSaving;
});

const openRenameModal = () => renameModal.open({
  schema: editorStore.schema,
  updateSchema: (attrs) => void Object.assign(editorStore.schema, attrs),
});

async function openAccessEditModal(): Promise<void> {
  if (!isSharedAccessAvailable.value) {
    const isUpgraded = await upgradePlanModal.open();
    if (!isUpgraded) return;
  }

  void accessEditModal.open({
    schemaIds: [editorStore.schema.id],
  });
}

const deleteConfirm = useConfirm({
  danger: true,
  message: 'Ви впевнені, що хочете видалити цю схему?',
  acceptButton: 'Видалити',
});

const deleteSchema = useAsyncAction(async () => {
  await editorStore.deleteSchema();
  await router.push({ name: 'home' });
});

async function deleteSchemaIntent() {
  const confirmation = await deleteConfirm.ask();
  if (confirmation.isAccepted) await deleteSchema();
}

function undo() {
  selectionStore.reset();
  historyStore.undo();
}

const hotKeys = useHotKeys({
  mac: {
    Meta_KeyZ: undo,
    Meta_Shift_KeyZ: historyStore.redo,
    Meta_KeyS: editorStore.save,
  },
  win: {
    Ctrl_KeyZ: undo,
    Ctrl_Shift_KeyZ: historyStore.redo,
    Ctrl_KeyS: editorStore.save,
  },
}, {
  isActive: () => editorStore.canEdit && !isMobile.value,
});

const undoTooltip = computed(() => {
  const hotKey = hotKeys.titles.Meta_KeyZ || hotKeys.titles.Ctrl_KeyZ;
  return `Відмінити зміни (${hotKey})`;
});

const redoTooltip = computed(() => {
  const hotKey = hotKeys.titles.Meta_Shift_KeyZ || hotKeys.titles.Ctrl_Shift_KeyZ;
  return `Повернути назад зміни (${hotKey})`;
});

const savingTitle = computed((): string => {
  const hotKey = hotKeys.titles.Meta_KeyS || hotKeys.titles.Ctrl_KeyS;

  if (editorStore.isSaving) {
    return 'Зміни зберігаються прямо зараз';
  }

  if (editorStore.hasUnsavedChanges) {
    return `Зберегти зміни (${hotKey})`;
  }

  return 'Всі зміни збережено';
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
