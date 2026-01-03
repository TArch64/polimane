<template>
  <p>
    subscription
  </p>
</template>

<script setup lang="ts">
import { definePreload } from '@/router/define';

defineOptions({
  beforeRouteEnter: definePreload<'settings-subscription'>(async () => {
  }),
});
</script>
