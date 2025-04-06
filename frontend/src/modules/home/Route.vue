<template>
  <HomeTopBar />
</template>

<script setup lang="ts">
import { definePreload } from '@/router/define';
import { HomeTopBar } from './components';
import { useSchemasStore } from './stores';

defineOptions({
  beforeRouteEnter: definePreload<'welcome'>(async () => {
    const store = useSchemasStore();
    await store.schemas.load();
  }),
});
</script>
