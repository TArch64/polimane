<template>
  <Card as="div" class="subscription-plan" :class="classes">
    <div class="subscription-plan__column">
      <h3 class="subscription-plan__name">
        {{ name }}
      </h3>

      <div class="subscription-plan__limits">
        <PlanLimit
          v-for="(config, limit) in PLAN_LIMIT_CONFIGS"
          :key="limit"
          :plan
          :limit
          :config
        />
      </div>

      <p class="subscription-plan__price">
        {{ formattedMonthlyPrice }}/місяць
      </p>

      <p v-if="isActive">
        Активна Підписка
      </p>

      <Button
        variant="primary"
        size="lg"
        class="subscription-plan__activate"
        :loading="changePlan.isActive"
        @click="changePlan"
        v-else
      >
        Обрати Підписку
      </Button>
    </div>
  </Card>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import type { ISubscriptionPlan } from '@/models';
import { getSubscriptionPlanName, SubscriptionPlanId } from '@/enums';
import { Card } from '@/components/card';
import { useAsyncAction, useCurrencyFormatter } from '@/composables';
import { Button } from '@/components/button';
import { PLAN_LIMIT_CONFIGS, useSubscriptionStore } from '../../stores';
import PlanLimit from './PlanLimit.vue';

const props = defineProps<{
  plan: ISubscriptionPlan;
}>();

const subscriptionStore = useSubscriptionStore();

const name = computed(() => getSubscriptionPlanName(props.plan.id));

const isControl = computed(() => props.plan.id === SubscriptionPlanId.PRO);
const isActive = computed(() => props.plan.id === subscriptionStore.subscription.planId);

const classes = computed(() => ({
  'subscription-plan--control': isControl.value,
}));

const formattedMonthlyPrice = useCurrencyFormatter(() => {
  return props.plan.monthlyPrice;
});

const changePlan = useAsyncAction(async () => {
  await subscriptionStore.changePlan(props.plan.id);
});
</script>

<style scoped>
@layer page {
  .subscription-plan {
    border: none;
    padding: 100px 40px;
    border-radius: calc(var(--rounded-lg) - 1px);
  }

  .subscription-plan--control {
    --card-background: color-mix(in srgb, var(--color-primary), var(--color-white) 10%);
    --button-base-color: var(--color-white);
    color: var(--color-white);

    .subscription-plan__activate {
      --button-foreground: var(--color-primary);
    }
  }

  .subscription-plan__column {
    max-width: 280px;
    margin: 0 auto;
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 24px;
  }

  .subscription-plan__name {
    font-size: var(--font-xlg);
    font-weight: 600;
    margin-bottom: 4px;
  }

  .subscription-plan__limits {
    display: flex;
    flex-direction: column;
    width: 100%;
    gap: 4px;
  }

  .subscription-plan__price {
    font-weight: 500;
    margin-bottom: 4px;
  }

  @media (max-width: 768px) {
    .subscription-plan {
      padding: 40px;
    }
  }
}
</style>
