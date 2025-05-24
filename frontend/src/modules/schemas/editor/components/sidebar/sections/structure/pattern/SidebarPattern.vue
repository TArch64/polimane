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
import type { ISchemaPattern } from '@/models';
import { useModal } from '@/components/modal';
import { getPatternAddRowModal } from '@/modules/schemas/editor/components/modals';
import { usePatternContextMenuActions } from '@/modules/schemas/editor/composables';
import { SidebarStructureEmpty, SidebarStructureItem } from '../base';
import { SidebarRowList } from '../row';

const props = defineProps<{
  pattern: ISchemaPattern;
}>();

const DEPTH = 0;

const actions = usePatternContextMenuActions(() => props.pattern);

const addRowModal = useModal(getPatternAddRowModal(props.pattern));
const addRow = () => addRowModal.open({ pattern: props.pattern });
</script>
