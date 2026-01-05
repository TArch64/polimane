<template>
  <LimitView>
    <LimitInfo
      :title="limit.title"
      :value="`${limit.max}`"
    />
  </LimitView>
</template>

<script setup lang="ts">
import type { ISubscriptionLimit } from '../../stores';
import { LimitInfo, LimitView } from './base';

defineProps<{
  limit: ISubscriptionLimit;
}>();
</script>
