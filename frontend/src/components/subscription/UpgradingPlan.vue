<template>
  <Card variant="inverted" class="upgrading-plan">
    <div class="upgrading-plan__main">
      <h3 class="upgrading-plan__name">
        {{ name }}
      </h3>

      <p class="upgrading-plan__price">
        {{ formattedPrice }}/місяць
      </p>

      <Button
        variant="primary"
        :loading="upgradePlan.isActive"
        @click="upgradePlan"
      >
        Підписатися
      </Button>
    </div>

    <div class="upgrading-plan__limits">
      <UpgradingPlanLimit
        v-for="(config, limit) in PLAN_LIMIT_CONFIGS"
        :key="limit"
        :plan
        :limit="limit as SubscriptionLimit"
        :config
      />
    </div>
  </Card>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { Card } from '@/components/card';
import type { ISubscriptionPlan } from '@/models';
import { getSubscriptionPlanName, SubscriptionLimit } from '@/enums';
import { Button } from '@/components/button';
import { useAsyncAction, useCurrencyFormatter } from '@/composables';
import { PLAN_LIMIT_CONFIGS } from '@/config';
import { useSessionStore } from '@/stores';
import UpgradingPlanLimit from './UpgradingPlanLimit.vue';

const props = defineProps<{
  plan: ISubscriptionPlan;
}>();

const emit = defineEmits<{
  upgraded: [];
}>();

const sessionStore = useSessionStore();

const name = computed(() => getSubscriptionPlanName(props.plan.id));
const formattedPrice = useCurrencyFormatter(() => props.plan.monthlyPrice);

const upgradePlan = useAsyncAction(async () => {
  await sessionStore.changePlan(props.plan.id);
  emit('upgraded');
});
</script>

<style scoped>
@layer components {
  .upgrading-plan {
    display: flex;
    align-items: flex-start;
    padding: 20px;
    gap: 16px;
  }

  .upgrading-plan__main {
    width: 40%;
  }

  .upgrading-plan__limits {
    flex-grow: 1;
    display: flex;
    flex-direction: column;
    gap: 4px;
  }

  .upgrading-plan__name {
    font-size: var(--font-lg);
    font-weight: 600;
    margin-bottom: 12px;
  }

  .upgrading-plan__price {
    margin-bottom: 16px;
  }
}
</style>
