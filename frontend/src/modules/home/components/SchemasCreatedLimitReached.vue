<template>
  <Banner type="danger" :prepend-icon="LockIcon" v-if="limit.isReached">
    <p>Ліміт схем перевищено: <b>{{ limit.current }}</b> / <b>{{ limit.max }}</b></p>
    <p>Оновіть підписку або видаліть схеми</p>
  </Banner>
</template>

<script setup lang="ts">
import { useSchemasCreatedLimit } from '@/composables/subscription';
import { Banner } from '@/components/banner';
import { LockIcon } from '@/components/icon';

const limit = useSchemasCreatedLimit();
</script>
