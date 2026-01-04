<template>
  <Component :is="component" :limit />
</template>

<script setup lang="ts">
import { type Component, computed, markRaw } from 'vue';
import { type ISubscriptionLimit, SubscriptionLimitType } from '../../stores';
import LimitCounter from './LimitCounter.vue';
import LimitPerFeature from './LimitPerFeature.vue';

const props = defineProps<{
  limit: ISubscriptionLimit;
}>();

type LimitComponent = Component<{
  limit: ISubscriptionLimit;
}>;

const components: Record<SubscriptionLimitType, LimitComponent> = {
  [SubscriptionLimitType.COUNTER]: markRaw(LimitCounter),
  [SubscriptionLimitType.PER_FEATURE]: markRaw(LimitPerFeature),
};

const component = computed(() => components[props.limit.type]);
</script>
