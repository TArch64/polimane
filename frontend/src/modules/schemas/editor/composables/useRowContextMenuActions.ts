import { computed, type MaybeRefOrGetter, nextTick, toValue } from 'vue';
import type { ISchemaPattern, ISchemaRow } from '@/models';
import type { MaybeContextMenuAction } from '@/components/contextMenu';
import { TrashIcon } from '@/components/icon';
import { useConfirm } from '@/components/confirm';
import { useRouteTransition } from '@/composables';
import { useRowsStore } from '../stores';
import { getObjectParent } from '../models';
import { useRowTitle } from './useRowTitle';

export function useRowContextMenuActions(rowRef: MaybeRefOrGetter<ISchemaRow>): MaybeContextMenuAction[] {
  const row = computed(() => toValue(rowRef));
  const pattern = computed(() => getObjectParent<ISchemaPattern>(row.value));

  const rowsStore = useRowsStore(pattern);
  const routeTransition = useRouteTransition();
  const title = useRowTitle(row);

  const deleteConfirm = useConfirm({
    danger: true,
    message: `Ви впевнені, що хочете видалити '${title.value}'?`,
    acceptButton: 'Видалити',
  });

  return [
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
