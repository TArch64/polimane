import { computed, type ComputedRef, type MaybeRefOrGetter, toValue } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import type { MaybeContextMenuAction } from '@/components/contextMenu';
import { CopyIcon, EditIcon, FolderIcon, PeopleIcon, TrashIcon } from '@/components/icon';
import { useAccessPermissions } from '@/composables';
import { useModal } from '@/components/modal';
import { SchemaRenameModal } from '@/modules/schemas/shared/modals/rename';
import {
  SchemaAccessEditModal,
  useSchemaUsersStore,
} from '@/modules/schemas/shared/modals/accessEdit';
import type { ISchema } from '@/models';
import { FolderAddSchemaModal } from '../components';
import { type ListSchema, useHomeFoldersStore, useHomeStore } from '../stores';

export function useSchemaMenuActions(schemaRef: MaybeRefOrGetter<ListSchema>): ComputedRef<MaybeContextMenuAction[]> {
  const schema = computed(() => toValue(schemaRef));

  const homeStore = useHomeStore();
  const homeFoldersStore = useHomeFoldersStore();
  const schemaUsersStore = useSchemaUsersStore();
  const { updateSchema, copySchema, deleteSchema } = homeStore;

  const router = useRouter();
  const route = useRoute();
  const permissions = useAccessPermissions(() => schema.value.access);

  const renameModal = useModal(SchemaRenameModal);
  const folderAddModal = useModal(FolderAddSchemaModal);
  const accessEditModal = useModal(SchemaAccessEditModal);

  return computed((): MaybeContextMenuAction[] => [
    permissions.write && {
      title: 'Змінити Назву',
      icon: EditIcon,

      onAction() {
        void renameModal.open({
          schema: schema.value as ISchema,
          updateSchema: (attrs) => updateSchema.do(schema.value, attrs),
        });
      },
    },

    {
      title: 'Перемістити в Директорію',
      icon: FolderIcon,

      async onAction() {
        await homeFoldersStore.load();

        void folderAddModal.open({
          schemaIds: [schema.value.id],
          folderId: null,
        });
      },
    },

    {
      title: 'Зробити Копію',
      icon: CopyIcon,

      async onAction() {
        const created = await copySchema.do(schema.value);

        await router.push({
          name: 'schema-editor',
          params: { schemaId: created.id },
          query: { from: route.path },
        });
      },
    },

    permissions.admin && {
      title: 'Редагувати Доступ',
      icon: PeopleIcon,

      async onAction() {
        await schemaUsersStore.load([schema.value.id]);
        void accessEditModal.open();
      },
    },

    permissions.admin && {
      danger: true,
      title: 'Видалити',
      icon: TrashIcon,

      async onAction(event) {
        await deleteSchema.do(schema.value);
      },
    },
  ]);
}
