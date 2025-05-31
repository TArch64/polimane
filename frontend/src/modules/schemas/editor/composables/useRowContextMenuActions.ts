import { computed, type MaybeRefOrGetter, nextTick, toValue } from 'vue';
import type { ISchemaPattern, ISchemaRow } from '@/models';
import type { MaybeContextMenuAction } from '@/components/contextMenu';
import { ArrowDownwardIcon, ArrowUpwardIcon, TrashIcon } from '@/components/icon';
import { useConfirm } from '@/components/confirm';
import { useRouteTransition } from '@/composables';
import { useModal } from '@/components/modal';
import { getPatternAddRowModal } from '@/modules/schemas/editor/components/modals';
import { useRowsStore } from '../stores';
import { getObjectParent } from '../models';
import { useRowTitle } from './useRowTitle';

export function useRowContextMenuActions(rowRef: MaybeRefOrGetter<ISchemaRow>): MaybeContextMenuAction[] {
  const row = computed(() => toValue(rowRef));
  const pattern = computed(() => getObjectParent<ISchemaPattern>(row.value)!);

  const rowsStore = useRowsStore(pattern);
  const routeTransition = useRouteTransition();
  const title = useRowTitle(row);

  const addModal = useModal(getPatternAddRowModal(pattern.value));

  function addRow(after: boolean): void {
    const index = rowsStore.rows.indexOf(row.value);
    const toIndex = after ? index + 1 : index;
    addModal.open({ pattern: pattern.value, toIndex });
  }

  const deleteConfirm = useConfirm({
    danger: true,
    message: `Ви впевнені, що хочете видалити '${title.value}'?`,
    acceptButton: 'Видалити',
  });

  return [
    {
      title: 'Додати Зверху',
      icon: ArrowUpwardIcon,
      onAction: () => addRow(false),
    },

    {
      title: 'Додати Знизу',
      icon: ArrowDownwardIcon,
      onAction: () => addRow(true),
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
  ];
}
