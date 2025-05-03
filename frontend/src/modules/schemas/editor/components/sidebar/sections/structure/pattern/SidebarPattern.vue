<template>
  <SidebarStructureItem
    :depth="DEPTH"
    :object="pattern"
    :title="pattern.name"
    :more-actions-button-style="deleteConfirm.anchorStyle"
  >
    <template #actions>
      <DropdownAction
        title="Переназвати Паттерн"
        :icon="EditIcon"
        @click="renamePattern"
      />

      <DropdownAction
        danger
        title="Видалити Паттерн"
        :icon="TrashIcon"
        @click="deletePattern"
      />
    </template>

    <template #content>
      <SidebarRowList :pattern :depth="DEPTH + 1" />
    </template>
  </SidebarStructureItem>
</template>

<script setup lang="ts">
import { nextTick } from 'vue';
import type { ISchemaPattern } from '@/models';
import { usePatternsStore } from '@/modules/schemas/editor/stores';
import { DropdownAction } from '@/components/dropdown';
import { EditIcon, TrashIcon } from '@/components/icon';
import { useModal } from '@/components/modal';
import { useConfirm } from '@/components/confirm';
import { useRouteTransition } from '@/composables';
import { SidebarStructureItem } from '../base';
import { SidebarRowList } from '../row';
import PatternRenameModal from './PatternRenameModal.vue';

const props = defineProps<{
  pattern: ISchemaPattern;
}>();

const DEPTH = 0;

const patternsStore = usePatternsStore();
const renameModal = useModal(PatternRenameModal);
const routeTransition = useRouteTransition();

const deleteConfirm = useConfirm({
  danger: true,
  message: 'Ви впевнені, що хочете видалити цей паттерн?',
  acceptButton: 'Видалити',
});

async function deletePattern(): Promise<void> {
  if (await deleteConfirm.ask()) {
    routeTransition.start(async () => {
      patternsStore.deletePattern(props.pattern);
      await nextTick();
    });
  }
}

function renamePattern(): void {
  renameModal.open({ pattern: props.pattern });
}
</script>
