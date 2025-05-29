<template>
  <SidebarStructureItem
    :actions
    :title
    :depth
    :object="row"
  />
</template>

<script setup lang="ts">
import type { ISchemaPattern, ISchemaRow } from '@/models';
import { useRowContextMenuActions, useRowTitle } from '@/modules/schemas/editor/composables';
import { SidebarStructureItem } from '../base';

const props = defineProps<{
  depth: number;
  row: ISchemaRow;
  pattern: ISchemaPattern;
}>();

const title = useRowTitle(() => props.row);
const actions = useRowContextMenuActions(() => props.row);
</script>
