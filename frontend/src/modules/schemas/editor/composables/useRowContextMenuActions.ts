import { toRef } from '@vueuse/core';
import { computed, type MaybeRefOrGetter, nextTick, type Ref } from 'vue';
import type { ISchemaPattern, ISchemaRow } from '@/models';
import type { MaybeContextMenuAction } from '@/components/contextMenu';
import {
  ArrowDownwardIcon,
  ArrowUpwardIcon,
  ExpandIcon,
  PlusIcon,
  TrashIcon,
} from '@/components/icon';
import { useConfirm } from '@/components/confirm';
import { useRouteTransition } from '@/composables';
import { RowAddModal } from '@/modules/schemas/editor/components/modals';
import { useModal } from '@/components/modal';
import { useRowsStore } from '../stores';
import { useObjectParent } from '../models';
import { useRowTitle } from './useRowTitle';

export function useRowContextMenuActions(rowRef: MaybeRefOrGetter<ISchemaRow>): Ref<MaybeContextMenuAction[]> {
  const row = toRef(rowRef);
  const pattern = useObjectParent<ISchemaPattern>(rowRef);

  const rowsStore = useRowsStore(pattern);
  const routeTransition = useRouteTransition();
  const title = useRowTitle(row);
  const rowIndex = computed(() => rowsStore.rows.indexOf(row.value));

  const addModal = useModal(RowAddModal);

  function addRow(after: boolean): void {
    const index = rowsStore.rows.indexOf(row.value);
    const toIndex = after ? index + 1 : index;
    addModal.open({ pattern: pattern.value, toIndex });
  }

  const deleteConfirm = useConfirm({
    danger: true,
    control: false,
    message: `Ви впевнені, що хочете видалити '${title.value}'?`,
    acceptButton: 'Видалити',
  });

  return computed((): MaybeContextMenuAction[] => [
    {
      title: 'Додати Рядок',
      icon: PlusIcon,

      actions: [
        {
          title: 'Зверху',
          icon: ArrowUpwardIcon,
          onAction: () => addRow(false),
        },

        {
          title: 'Знизу',
          icon: ArrowDownwardIcon,
          onAction: () => addRow(true),
        },
      ],
    },

    {
      title: 'Перемістити Рядок',
      icon: ExpandIcon,
      disabled: rowIndex.value === 0 && rowIndex.value === rowsStore.rows.size - 1,

      actions: [
        {
          title: 'Вверх',
          icon: ArrowUpwardIcon,
          disabled: rowIndex.value === 0,
          onAction: () => rowsStore.moveRow(row.value, -1),
        },

        {
          title: 'Вниз',
          icon: ArrowDownwardIcon,
          disabled: rowIndex.value === rowsStore.rows.size - 1,
          onAction: () => rowsStore.moveRow(row.value, 1),
        },
      ],
    },

    {
      title: 'Видалити Рядок',
      icon: TrashIcon,
      danger: true,

      onAction: async (event) => {
        if (await deleteConfirm.ask({ virtualTarget: event.menuRect })) {
          routeTransition.start(async () => {
            rowsStore.deleteRow(row.value);
            await nextTick();
          });
        }
      },
    },
  ]);
}
