<template>
  <Component :is="component" :limit />
</template>

<script setup lang="ts">
import { type Component, computed, markRaw } from 'vue';
import { type ISubscriptionLimit, SubscriptionLimitType } from '../../stores';
import LimitUser from './LimitUser.vue';
import LimitFeature from './LimitFeature.vue';

const props = defineProps<{
  limit: ISubscriptionLimit;
}>();

type LimitComponent = Component<{
  limit: ISubscriptionLimit;
}>;

const components: Record<SubscriptionLimitType, LimitComponent> = {
  [SubscriptionLimitType.USER]: markRaw(LimitUser),
  [SubscriptionLimitType.FEATURE]: markRaw(LimitFeature),
};

const component = computed(() => components[props.limit.type]);
</script>
