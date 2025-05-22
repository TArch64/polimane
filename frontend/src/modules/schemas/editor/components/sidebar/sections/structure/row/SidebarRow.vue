<template>
  <SidebarStructureItem
    :actions
    :title
    :depth
    :object="row"
  />
</template>

<script setup lang="ts">
import { computed, nextTick } from 'vue';
import type { ISchemaPattern, ISchemaRow } from '@/models';
import { useRowsStore } from '@/modules/schemas/editor/stores';
import { TrashIcon } from '@/components/icon';
import { useConfirm } from '@/components/confirm';
import { useRouteTransition } from '@/composables';
import type { MaybeContextMenuAction } from '@/components/contextMenu';
import { SidebarStructureItem } from '../base';

const props = defineProps<{
  depth: number;
  row: ISchemaRow;
  pattern: ISchemaPattern;
}>();

const rowsStore = useRowsStore(() => props.pattern);
const routeTransition = useRouteTransition();

const index = computed(() => rowsStore.rows.indexOf(props.row));
const title = computed(() => `Рядок #${index.value + 1}`);

const deleteConfirm = useConfirm({
  danger: true,
  message: `Ви впевнені, що хочете видалити '${title.value}'?`,
  acceptButton: 'Видалити',
});

const actions: MaybeContextMenuAction[] = [
  {
    title: 'Видалити Рядок',
    icon: TrashIcon,
    danger: true,
    onAction: async () => {
      if (await deleteConfirm.ask()) {
        routeTransition.start(async () => {
          rowsStore.deleteRow(props.row);
          await nextTick();
        });
      }
    },
  },
];
</script>
