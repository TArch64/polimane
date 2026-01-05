<template>
  <LimitView
    :title="limit.title"
    :subtitle="`${limit.max}`"
    :progress="0"
  />
</template>

<script setup lang="ts">
import type { ISubscriptionLimit } from '../../stores';
import LimitView from './LimitView.vue';

defineProps<{
  limit: ISubscriptionLimit;
}>();
</script>
