import { computed, type MaybeRefOrGetter, type Ref, toValue } from 'vue';
import type { MaybeContextMenuAction } from '@/components/contextMenu';
import { type ListSchema, useHomeFoldersStore } from '@/modules/home/stores';
import { FolderIcon } from '@/components/icon';
import { useModal } from '@/components/modal';
import { FolderAddSchemaModal } from '@/modules/home/components';

export function useSchemaMenuAddToDirectory(
  schemaRef: MaybeRefOrGetter<ListSchema>,
  folderId: string | null = null,
): Ref<MaybeContextMenuAction> {
  const schema = computed(() => toValue(schemaRef));

  const homeFoldersStore = useHomeFoldersStore();

  const folderAddModal = useModal(FolderAddSchemaModal);

  return computed(() => ({
    title: 'Перемістити в Директорію',
    icon: FolderIcon,

    async onAction() {
      await homeFoldersStore.load();

      void folderAddModal.open({
        schemaIds: [schema.value.id],
        folderId,
      });
    },
  }));
}
