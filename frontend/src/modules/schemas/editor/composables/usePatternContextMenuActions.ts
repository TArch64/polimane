import { type MaybeRefOrGetter, nextTick } from 'vue';
import { toRef } from '@vueuse/core';
import type { MaybeContextMenuAction } from '@/components/contextMenu';
import {
  ArrowDownwardIcon,
  ArrowUpwardIcon,
  EditIcon,
  PlusIcon,
  TrashIcon,
} from '@/components/icon';
import { useModal } from '@/components/modal';
import type { ISchemaPattern } from '@/models';
import { useRouteTransition } from '@/composables';
import { useConfirm } from '@/components/confirm';
import { PatternRenameModal, usePatternAddModal } from '../components/modals';
import { usePatternsStore } from '../stores';

export function usePatternContextMenuActions(patternRef: MaybeRefOrGetter<ISchemaPattern>): MaybeContextMenuAction[] {
  const pattern = toRef(patternRef);

  const routeTransition = useRouteTransition();
  const patternsStore = usePatternsStore();

  const renameModal = useModal<typeof PatternRenameModal, void>(PatternRenameModal);
  const addModal = usePatternAddModal();

  const deleteConfirm = useConfirm({
    danger: true,
    control: false,
    message: 'Ви впевнені, що хочете видалити цей паттерн?',
    acceptButton: 'Видалити',
  });

  function addPattern(after: boolean): void {
    const index = patternsStore.patterns.indexOf(pattern.value);
    const toIndex = after ? index + 1 : index;
    addModal.open({ toIndex });
  }

  return [
    {
      title: 'Переназвати Паттерн',
      icon: EditIcon,
      onAction: () => renameModal.open({ pattern: pattern.value }),
    },

    {
      title: 'Додати Паттерн',
      icon: PlusIcon,

      actions: [
        {
          title: 'Додати Зверху',
          icon: ArrowUpwardIcon,
          onAction: () => addPattern(false),
        },

        {
          title: 'Додати Знизу',
          icon: ArrowDownwardIcon,
          onAction: () => addPattern(true),
        },
      ],
    },

    {
      danger: true,
      title: 'Видалити Паттерн',
      icon: TrashIcon,

      onAction: async (event) => {
        if (await deleteConfirm.ask({ virtualTarget: event.menuRect })) {
          routeTransition.start(async () => {
            patternsStore.deletePattern(pattern.value);
            await nextTick();
          });
        }
      },
    },
  ];
}
