import { computed, type Ref } from 'vue';
import type { MaybeContextMenuAction } from '@/components/contextMenu';
import { useModal } from '@/components/modal';
import { FolderIcon } from '@/components/icon';
import { FolderAddSchemaModal } from '../../components';
import { useHomeFoldersStore } from '../../stores';

export function useSchemaSelectionAddToFolder(actionIds: Ref<string[]>): Ref<MaybeContextMenuAction> {
  const homeFoldersStore = useHomeFoldersStore();

  const folderAddModal = useModal(FolderAddSchemaModal);

  return computed(() => ({
    title: 'Перемістити в Директорію',
    icon: FolderIcon,

    async onAction() {
      await homeFoldersStore.load();

      void folderAddModal.open({
        schemaIds: actionIds.value,
        folderId: null,
      });
    },
  }));
}
