<template>
  <LimitView>
    <LimitInfo :title="limit.title" :value="formattedValue" />
  </LimitView>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { useNumberFormatter } from '@/composables';
import { INFINITY_SYMBOL } from '@/config';
import type { ISubscriptionLimit } from '../../stores';
import { LimitInfo, LimitView } from './base';

const props = defineProps<{
  limit: ISubscriptionLimit;
}>();

const formattedMax = useNumberFormatter(() => props.limit.max);
const formattedValue = computed(() => formattedMax.value || INFINITY_SYMBOL);
</script>
