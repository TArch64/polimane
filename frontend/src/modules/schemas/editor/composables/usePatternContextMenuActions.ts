import { computed, type MaybeRefOrGetter, nextTick, type Ref } from 'vue';
import { toRef } from '@vueuse/core';
import type { MaybeContextMenuAction } from '@/components/contextMenu';
import {
  ArrowDownwardIcon,
  ArrowUpwardIcon,
  EditIcon,
  MoveIcon,
  PlusIcon,
  TrashIcon,
} from '@/components/icon';
import { useModal } from '@/components/modal';
import type { ISchemaPattern } from '@/models';
import { useRouteTransition } from '@/composables';
import { useConfirm } from '@/components/confirm';
import { PatternRenameModal, usePatternAddModal } from '../components/modals';
import { usePatternsStore } from '../stores';

export function usePatternContextMenuActions(patternRef: MaybeRefOrGetter<ISchemaPattern>): Ref<MaybeContextMenuAction[]> {
  const pattern = toRef(patternRef);

  const routeTransition = useRouteTransition();
  const patternsStore = usePatternsStore();
  const patternIndex = computed(() => patternsStore.patterns.indexOf(pattern.value));

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

  return computed(() => [
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
      title: 'Перемістити Паттерн',
      icon: MoveIcon,

      actions: [
        {
          title: 'Вверх',
          icon: ArrowUpwardIcon,
          disabled: patternIndex.value === 0,
          onAction: () => patternsStore.movePattern(pattern.value, -1),
        },

        {
          title: 'Вниз',
          icon: ArrowDownwardIcon,
          disabled: patternIndex.value === patternsStore.patterns.size - 1,
          onAction: () => patternsStore.movePattern(pattern.value, 1),
        },
      ],
    },

    {
      danger: true,
      title: 'Видалити Паттерн',
      icon: TrashIcon,

      onAction: () => {
        routeTransition.start(async () => {
          patternsStore.deletePattern(pattern.value);
          await nextTick();
        });
      },
    },
  ]);
}
