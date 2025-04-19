<template>
  <SidebarStructureItem
    :title="pattern.name"
    :active="isActive"
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
  </SidebarStructureItem>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import type { ISchemaPattern } from '@/models';
import { useEditorStore, usePatternsStore } from '@/modules/schemas/editor/stores';
import { DropdownAction } from '@/components/dropdown';
import { EditIcon, TrashIcon } from '@/components/icon';
import { useModal } from '@/components/modal';
import { useConfirm } from '@/components/confirm';
import { SidebarStructureItem } from '../base';
import PatternRenameModal from './PatternRenameModal.vue';

const props = defineProps<{
  pattern: ISchemaPattern;
}>();

const editorStore = useEditorStore();
const patternsStore = usePatternsStore();
const renameModal = useModal(PatternRenameModal);

const isActive = computed(() => editorStore.activePattern?.id === props.pattern.id);

const deleteConfirm = useConfirm({
  message: 'Ви впевнені, що хочете видалити цей паттерн?',

  acceptButton: {
    text: 'Видалити',
    danger: true,
  },
});

async function deletePattern(): Promise<void> {
  if (await deleteConfirm.ask()) {
    patternsStore.deletePattern(props.pattern);
  }
}

function renamePattern(): void {
  renameModal.open({ pattern: props.pattern });
}
</script>
