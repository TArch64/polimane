<template>
  <Card as="div" class="subscription-plan" :variant="cardVariant">
    <div class="subscription-plan__column">
      <h3 class="subscription-plan__name">
        {{ name }}
      </h3>

      <div class="subscription-plan__limits">
        <PlanLimit
          v-for="(config, limit) in PLAN_LIMIT_CONFIGS"
          :key="limit"
          :plan
          :limit="limit as SubscriptionLimit"
          :config
        />
      </div>

      <p class="subscription-plan__price">
        {{ formattedPrice }}/місяць
      </p>

      <p v-if="isActive">
        Активна Підписка
      </p>

      <Button
        variant="primary"
        size="lg"
        class="subscription-plan__activate"
        :loading="changePlan.isActive"
        :style="downgradePlanConfirm.anchorStyle"
        @click="changePlanIntent"
        v-else
      >
        Підписатися
      </Button>
    </div>
  </Card>
</template>

<script setup lang="ts">
import { computed, toRef } from 'vue';
import type { ISubscriptionPlan } from '@/models';
import { getSubscriptionPlanName, SubscriptionLimit, SubscriptionPlanId } from '@/enums';
import { Card } from '@/components/card';
import { useAsyncAction, useCurrencyFormatter } from '@/composables';
import { Button } from '@/components/button';
import { useSessionStore } from '@/stores';
import type { ComponentVariant } from '@/types';
import { PLAN_LIMIT_CONFIGS } from '@/config';
import { useDowngradePlanConfirm } from '@/composables/subscription';
import PlanLimit from './PlanLimit.vue';

const props = defineProps<{
  plan: ISubscriptionPlan;
}>();

const emit = defineEmits<{
  upgraded: [];
}>();

const sessionStore = useSessionStore();

const downgradePlanConfirm = useDowngradePlanConfirm(toRef(props, 'plan'));

const name = computed(() => getSubscriptionPlanName(props.plan.id));

const isControl = computed(() => props.plan.id === SubscriptionPlanId.PRO);
const cardVariant = computed((): ComponentVariant => isControl.value ? 'inverted' : 'main');

const isActive = computed(() => props.plan.id === sessionStore.plan.id);
const isLowerPlan = computed(() => props.plan.tier < sessionStore.plan.tier);

const formattedPrice = useCurrencyFormatter(() => props.plan.monthlyPrice);

const changePlan = useAsyncAction(async () => {
  await sessionStore.changePlan(props.plan.id);
});

async function changePlanIntent(): Promise<void> {
  const wasLowerPlan = isLowerPlan.value;

  if (wasLowerPlan) {
    const result = await downgradePlanConfirm.ask();
    if (!result.isAccepted) return;
  }

  await changePlan();

  if (!wasLowerPlan) {
    emit('upgraded');
  }
}
</script>

<style scoped>
@layer page {
  .subscription-plan {
    border: none;
    padding: 100px 40px;
    border-radius: calc(var(--rounded-lg) - 1px);
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
