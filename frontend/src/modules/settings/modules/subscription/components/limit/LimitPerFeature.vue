<template>
  <LimitView>
    <LimitInfo
      :title="limit.title"
      :value="max"
    />
  </LimitView>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import type { ISubscriptionLimit } from '../../stores';
import { LimitInfo, LimitView } from './base';

const props = defineProps<{
  limit: ISubscriptionLimit;
}>();

const max = computed(() => props.limit.max?.toString() ?? '∞');
</script>
