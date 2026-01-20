<template>
  <div class="home-banners" v-if="schemasCreatedLimit.isReached">
    <SchemasCreatedLimitReached v-if="schemasCreatedLimit.isReached" />
  </div>
</template>

<script setup lang="ts">
import { useSchemasCreatedLimit } from '@/composables/subscription';
import SchemasCreatedLimitReached from './SchemasCreatedLimitReached.vue';

const schemasCreatedLimit = useSchemasCreatedLimit();
</script>

<style scoped>
@layer page {
  .home-banners {
    padding: 0 12px;
  }
}
</style>
