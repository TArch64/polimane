<template>
  <LimitView>
    <LimitInfo :title="limit.title" :value="formattedValue" />
    <LimitProgress :value="usage" />
  </LimitView>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { useNumberFormatter } from '@/composables';
import type { ISubscriptionLimit } from '../../stores';
import { LimitInfo, LimitProgress, LimitView } from './base';

const props = defineProps<{
  limit: ISubscriptionLimit;
}>();

const used = computed(() => props.limit.used!);
const max = computed(() => props.limit.max);

const usage = computed(() => {
  if (max.value === null) return 0;
  return Math.min((used.value / max.value) * 100, 100);
});

const formattedUsed = useNumberFormatter(used);
const formattedMax = useNumberFormatter(max);
const formattedValue = computed(() => `${formattedUsed.value} / ${formattedMax.value || '∞'}`);
</script>
