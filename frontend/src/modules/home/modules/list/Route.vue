<template>
  <HomeSchemasList v-if="schemasStore.hasSchemas" />
  <HomeSchemasEmpty v-else />

  <Teleport to="[data-home-top-bar-actions]">
    <HomeTopBarActions />
  </Teleport>
</template>

<script setup lang="ts">
import { definePreload } from '@/router/define';
import { useHomeListStore, useSchemasStore } from './stores';
import { useSchemasSelection } from './composables';
import { HomeSchemasEmpty, HomeSchemasList, HomeTopBarActions } from './components';

defineOptions({
  beforeRouteEnter: definePreload<'home'>(async () => {
    await useHomeListStore().load();
  }),
});

const schemasStore = useSchemasStore();
useSchemasSelection();
</script>
