<template>
  <HomeListCard
    :disabled
    :menuActions
    :to="editorRoute"
    :selected="isSelected"
    :menu-title="schema.name"
  >
    <HomeListScreenshot
      :path="schema.screenshotPath"
      :alt="`Скріншот схеми ${schema.name}`"
      :background-color="schema.backgroundColor"
    />

    {{ schema.name }}
  </HomeListCard>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { type RouteLocationRaw } from 'vue-router';
import type { MaybeContextMenuAction } from '@/components/contextMenu';
import { type ListSchema, useHomeStore } from '../../stores';
import HomeListCard from './HomeListCard.vue';
import HomeListScreenshot from './HomeListScreenshot.vue';

const props = defineProps<{
  schema: ListSchema;
  disabled?: boolean;
  menuActions: MaybeContextMenuAction[];
}>();

const homeStore = useHomeStore();

const isSelected = computed(() => homeStore.selection.ids.has(props.schema.id));

const editorRoute = computed((): RouteLocationRaw => ({
  name: 'schema-editor',
  params: { schemaId: props.schema.id },
}));
</script>
