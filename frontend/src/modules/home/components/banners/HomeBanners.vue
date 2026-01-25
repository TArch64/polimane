<template>
  <div class="home-banners" v-if="schemasCreatedCounter.isReached">
    <SchemasCreatedLimitReachedBanner v-if="schemasCreatedCounter.isReached" />
  </div>
</template>

<script setup lang="ts">
import { useSchemasCreatedCounter } from '@/composables/subscription';
import SchemasCreatedLimitReachedBanner from './SchemasCreatedLimitReachedBanner.vue';

const schemasCreatedCounter = useSchemasCreatedCounter();
</script>

<style scoped>
@layer page {
  .home-banners {
    padding: 0 12px;
  }
}
</style>
