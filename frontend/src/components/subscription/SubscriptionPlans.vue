<template>
  <Card class="subscription-plans" :class="classes">
    <SubscriptionPlan
      v-for="plan of plansStore.plans"
      :key="plan.id"
      :plan="plan"
      class="subscription-plans__item"
      @upgraded="$emit('upgraded')"
    />
  </Card>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { Card } from '@/components/card';
import { usePlansStore } from '@/stores';
import SubscriptionPlan from './SubscriptionPlan.vue';

const props = withDefaults(defineProps<{
  embedded?: boolean;
}>(), {
  embedded: false,
});

defineEmits<{
  upgraded: [];
}>();

const plansStore = usePlansStore();

const classes = computed(() => ({
  'subscription-plans--embedded': props.embedded,
}));
</script>

<style scoped>
@layer page {
  .subscription-plans {
    padding: 0;
    display: flex;
    min-width: 0;
  }

  .subscription-plans--embedded,
  .subscription-plans--embedded:deep(.card:not(.card--inverted)) {
    border: none;
    box-shadow: none;
  }

  .subscription-plans__item {
    flex-basis: 0;
    flex-grow: 1;
    min-width: 0;
  }

  @media (max-width: 768px) {
    .subscription-plans {
      flex-direction: column-reverse;
    }
  }
}
</style>
