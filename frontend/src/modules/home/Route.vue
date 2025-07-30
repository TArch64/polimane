<template>
  <CommonLayout>
    <template #top-bar-actions>
      <HomeTopBarActions />
    </template>

    <HomeSchemasList v-if="schemasStore.hasSchemas" />
    <HomeSchemasEmpty v-else />
  </CommonLayout>
</template>

<script setup lang="ts">
import { definePreload } from '@/router/define';
import { CommonLayout } from '@/components/layout';
import { HomeSchemasEmpty, HomeSchemasList, HomeTopBarActions } from './components';
import { useSchemasStore } from './stores';

defineOptions({
  beforeRouteEnter: definePreload<'home'>(async () => {
    const store = useSchemasStore();
    await store.schemas.load();
  }),
});

const schemasStore = useSchemasStore();
</script>
