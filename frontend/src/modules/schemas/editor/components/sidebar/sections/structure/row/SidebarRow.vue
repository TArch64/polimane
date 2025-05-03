<template>
  <SidebarStructureItem
    :title
    :depth
    :object="row"
    :more-actions-button-style="deleteConfirm.anchorStyle"
  >
    <template #actions>
      <DropdownAction
        danger
        title="Видалити Рядок"
        :icon="TrashIcon"
        @click="deleteRow"
      />
    </template>

    <template #content>
      <SidebarBeadList :row :depth="depth + 1" />
    </template>
  </SidebarStructureItem>
</template>

<script setup lang="ts">
import { computed, nextTick } from 'vue';
import type { ISchemaPattern, ISchemaRow } from '@/models';
import { useRowsStore } from '@/modules/schemas/editor/stores';
import { DropdownAction } from '@/components/dropdown';
import { TrashIcon } from '@/components/icon';
import { useConfirm } from '@/components/confirm';
import { useRouteTransition } from '@/composables';
import { SidebarStructureItem } from '../base';
import { SidebarBeadList } from '../bead';

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

async function deleteRow(): Promise<void> {
  if (await deleteConfirm.ask()) {
    routeTransition.start(async () => {
      rowsStore.deleteRow(props.row);
      await nextTick();
    });
  }
}
</script>
