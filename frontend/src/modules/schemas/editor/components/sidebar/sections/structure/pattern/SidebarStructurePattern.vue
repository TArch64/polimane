<template>
  <SidebarStructureItem :title="pattern.name" :active="isActive">
    <template #actions>
      <DropdownAction
        title="Переназвати Паттерн"
        :icon="EditIcon"
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
import { SidebarStructureItem } from '../base';

const props = defineProps<{
  pattern: ISchemaPattern;
}>();

const editorStore = useEditorStore();
const patternsStore = usePatternsStore();

const isActive = computed(() => editorStore.activePattern?.id === props.pattern.id);

function deletePattern(): void {
  patternsStore.deletePattern(props.pattern);
}
</script>
