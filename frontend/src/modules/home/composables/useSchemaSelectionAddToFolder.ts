import { computed, type Ref } from 'vue';
import type { MaybeContextMenuAction } from '@/components/contextMenu';
import { useModal } from '@/components/modal';
import { FolderAddSchemaModal } from '@/modules/home/components';
import { FolderIcon } from '@/components/icon';

export function useSchemaSelectionAddToFolder(actionIds: Ref<string[]>): Ref<MaybeContextMenuAction> {
  const folderAddModal = useModal(FolderAddSchemaModal);

  return computed(() => ({
    title: 'Перемістити в Директорію',
    icon: FolderIcon,

    onAction() {
      folderAddModal.open({
        schemaIds: actionIds.value,
        folderId: null,
      });
    },
  }));
}
