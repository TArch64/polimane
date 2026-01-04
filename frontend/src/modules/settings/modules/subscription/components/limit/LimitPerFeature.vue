<template>
  <p>
    {{ limit.key }} - ({{ limit.max }})
  </p>
</template>

<script setup lang="ts">
import type { ISubscriptionLimit } from '../../stores';

defineProps<{
  limit: ISubscriptionLimit;
}>();
</script>
