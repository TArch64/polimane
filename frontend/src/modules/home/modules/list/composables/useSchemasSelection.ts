import { computed, reactive, toRef } from 'vue';
import type { MaybeContextMenuAction } from '@/components/contextMenu';
import { AccessLevel } from '@/enums';
import {
  SchemaAccessEditModal,
  useSchemaUsersStore,
} from '@/modules/schemas/shared/modals/accessEdit';
import { useModal } from '@/components/modal';
import { useConfirm } from '@/components/confirm';
import { useAsyncAction } from '@/composables';
import { FolderIcon, PeopleIcon, TrashIcon } from '@/components/icon';
import { type IHomeSelectionState } from '@/modules/home/stores';
import { FolderAddSchemaModal } from '@/modules/home/components/modals';
import { useSchemasStore } from '../stores';

export function useSchemasSelection(): IHomeSelectionState {
  const schemasStore = useSchemasStore();
  const schemaUsersStore = useSchemaUsersStore();

  const folderAddModal = useModal(FolderAddSchemaModal);
  const accessEditModal = useModal(SchemaAccessEditModal);

  const count = computed(() => schemasStore.selected.size);
  const title = computed(() => `Обрано ${count.value} схем`);

  const allActionIds = computed(() => {
    return [...schemasStore.selected];
  });

  const adminActionIds = computed(() => {
    return schemasStore.filterIdsByAccess(schemasStore.selected, AccessLevel.ADMIN);
  });

  const deleteConfirm = useConfirm({
    danger: true,
    control: false,
    message: () => `Ви впевнені, що хочете видалити ${adminActionIds.value.length} схеми?`,
    acceptButton: 'Видалити',
  });

  const deleteSchemas = useAsyncAction(async () => {
    await schemasStore.deleteMany(adminActionIds.value);
    schemasStore.clearSelection();
  });

  const actions = computed((): MaybeContextMenuAction[] => {
    if (!allActionIds.value.length) {
      return [];
    }

    return [
      {
        title: 'Перемістити в Директорію',
        icon: FolderIcon,

        onAction() {
          folderAddModal.open({
            schemaIds: allActionIds.value,
            folderId: null,
          });
        },
      },

      !!adminActionIds.value.length && {
        title: 'Редагувати доступ',
        icon: PeopleIcon,

        async onAction() {
          await schemaUsersStore.load(adminActionIds.value);
          void accessEditModal.open();
        },
      },

      !!adminActionIds.value.length && {
        title: 'Видалити схеми',
        icon: TrashIcon,
        danger: true,

        async onAction(event) {
          if (await deleteConfirm.ask({ virtualTarget: event.menuRect })) {
            await deleteSchemas();
          }
        },
      },
    ];
  });

  return reactive({
    ids: toRef(schemasStore, 'selected'),
    title,
    actions,
    onClear: schemasStore.clearSelection,
  });
}
