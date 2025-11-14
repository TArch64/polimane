import { computed, reactive } from 'vue';
import type { MaybeContextMenuAction } from '@/components/contextMenu';
import { AccessLevel } from '@/enums';
import {
  SchemaAccessEditModal,
  useSchemaUsersStore,
} from '@/modules/schemas/shared/modals/accessEdit';
import { useModal } from '@/components/modal';
import { useConfirm } from '@/components/confirm';
import { useAsyncAction } from '@/composables';
import { PeopleIcon, TrashIcon } from '@/components/icon';
import { useSchemasStore } from '../stores';

export interface ISchemasSelection {
  count: number;
  title: string;
  actions: MaybeContextMenuAction[];
}

export function useSchemasSelection(): ISchemasSelection {
  const schemasStore = useSchemasStore();
  const schemaUsersStore = useSchemaUsersStore();

  const accessEditModal = useModal(SchemaAccessEditModal);

  const count = computed(() => schemasStore.selected.size);
  const title = computed(() => `Обрано ${count.value} схем`);

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
    if (!adminActionIds.value.length) {
      return [];
    }

    return [
      {
        title: 'Редагувати доступ',
        icon: PeopleIcon,

        async onAction() {
          await schemaUsersStore.load(adminActionIds.value);
          accessEditModal.open();
        },
      },

      {
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
    count,
    title,
    actions,
  });
}
