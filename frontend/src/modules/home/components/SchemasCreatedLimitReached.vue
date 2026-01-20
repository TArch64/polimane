<template>
  <Banner type="danger" :prepend-icon="LockIcon" v-if="limit.isReached">
    <p>
      Ви досягли максимальної кількості створених схем для вашого поточного плану підписки.
      Будь ласка, оновіть свій план, щоб створювати більше схем, або видаліть існуючі схеми, щоб
      звільнити місце для нових.
    </p>

    <p>
      Наразі ви маєте <b>{{ limit.current }}</b> схем із максимально дозволених <b>{{
        limit.max
      }}</b>
    </p>
  </Banner>
</template>

<script setup lang="ts">
import { useSchemasCreatedLimit } from '@/composables/subscription';
import { Banner } from '@/components/banner';
import { LockIcon } from '@/components/icon';

const limit = useSchemasCreatedLimit();
</script>
