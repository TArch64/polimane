<template>
  <SidebarStructureItem
    :actions
    :depth="DEPTH"
    :object="pattern"
    :title="pattern.name"
  >
    <template #content>
      <SidebarRowList
        :pattern
        :depth="DEPTH + 1"
        v-if="pattern.content.length"
      />

      <SidebarStructureEmpty
        button-text="Додати Рядок"
        @click="addRow"
        v-else
      />
    </template>
  </SidebarStructureItem>
</template>

<script setup lang="ts">
import { nextTick } from 'vue';
import type { ISchemaPattern } from '@/models';
import { usePatternsStore } from '@/modules/schemas/editor/stores';
import { ArrowDownwardIcon, ArrowUpwardIcon, EditIcon, TrashIcon } from '@/components/icon';
import { useModal } from '@/components/modal';
import { useConfirm } from '@/components/confirm';
import { useRouteTransition } from '@/composables';
import {
  getPatternAddRowModal,
  PatternAddModal,
  PatternRenameModal,
} from '@/modules/schemas/editor/components/modals';
import type { MaybeContextMenuAction } from '@/components/contextMenu';
import { SidebarStructureEmpty, SidebarStructureItem } from '../base';
import { SidebarRowList } from '../row';

const props = defineProps<{
  pattern: ISchemaPattern;
}>();

const DEPTH = 0;

const patternsStore = usePatternsStore();
const renameModal = useModal(PatternRenameModal);
const addModal = useModal(PatternAddModal);
const addRowModal = useModal(getPatternAddRowModal(props.pattern));
const routeTransition = useRouteTransition();

const deleteConfirm = useConfirm({
  danger: true,
  message: 'Ви впевнені, що хочете видалити цей паттерн?',
  acceptButton: 'Видалити',
});

const actions: MaybeContextMenuAction[] = [
  {
    title: 'Переназвати Паттерн',
    icon: EditIcon,
    onAction: () => renameModal.open({ pattern: props.pattern }),
  },

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

  {
    danger: true,
    title: 'Видалити Паттерн',
    icon: TrashIcon,

    onAction: async (event) => {
      if (await deleteConfirm.ask({ virtualTarget: event.menuRect })) {
        routeTransition.start(async () => {
          patternsStore.deletePattern(props.pattern);
          await nextTick();
        });
      }
    },
  },
];

function addPattern(after: boolean): void {
  const index = patternsStore.patterns.indexOf(props.pattern);
  const toIndex = after ? index + 1 : index;
  addModal.open({ toIndex });
}

function addRow(): void {
  addRowModal.open({ pattern: props.pattern });
}
</script>
