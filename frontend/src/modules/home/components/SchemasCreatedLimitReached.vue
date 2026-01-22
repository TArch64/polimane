<template>
  <Banner type="danger" :prepend-icon="LockIcon" v-if="counter.isReached">
    <p>Ліміт схем перевищено: <b>{{ counter.current }}</b> / <b>{{ counter.max }}</b></p>
    <p>Оновіть підписку або видаліть схеми</p>
  </Banner>
</template>

<script setup lang="ts">
import { useSchemasCreatedCounter } from '@/composables/subscription';
import { Banner } from '@/components/banner';
import { LockIcon } from '@/components/icon';

const counter = useSchemasCreatedCounter();
</script>
