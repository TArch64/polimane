<template>
  <HomeTopBar />
  <HomeSchemasList v-if="schemasStore.hasSchemas" />
  <HomeSchemasEmpty v-else />
</template>

<script setup lang="ts">
import { definePreload } from '@/router/define';
import { HomeSchemasEmpty, HomeSchemasList, HomeTopBar } from './components';
import { useSchemasStore } from './stores';

defineOptions({
  beforeRouteEnter: definePreload<'welcome'>(async () => {
    const store = useSchemasStore();
    await store.schemas.load();
  }),
});

const schemasStore = useSchemasStore();
</script>

<style>
@layer page {
  .app--home {
    background-color: var(--color-background-2);
  }
}
</style>
