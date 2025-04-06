<template>
  <HomeTopBar />
  <HomeSchemasEmpty />
</template>

<script setup lang="ts">
import { definePreload } from '@/router/define';
import { HomeSchemasEmpty, HomeTopBar } from './components';
import { useSchemasStore } from './stores';

defineOptions({
  beforeRouteEnter: definePreload<'welcome'>(async () => {
    const store = useSchemasStore();
    await store.schemas.load();
  }),
});
</script>
