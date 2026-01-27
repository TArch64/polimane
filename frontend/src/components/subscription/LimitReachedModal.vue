<template>
  <Modal :title :footer="false" :width="ModalWidth.LG">
    <div class="limit-reached__text">
      <slot :nextPlan name="description" />
    </div>

    <SubscriptionPlans
      embedded
      class="limit-reached__plans"
      @upgraded="modal.close(true)"
      v-if="nextPlan"
    />
  </Modal>
</template>

<script setup lang="ts">
import type { Slot } from 'vue';
import { Modal, ModalWidth, useActiveModal } from '@/components/modal';
import { usePlansStore, useSessionStore } from '@/stores';
import type { ISubscriptionPlan } from '@/models';
import SubscriptionPlans from './SubscriptionPlans.vue';

defineProps<{
  title: string;
}>();

defineSlots<{
  description: Slot<{
    nextPlan?: ISubscriptionPlan;
  }>;
}>();

const sessionStore = useSessionStore();
const plansStore = usePlansStore();

const modal = useActiveModal<boolean>();

// Need to be static to avoid deactivating listener when upgrading to last plan
const nextPlan = plansStore.plans.find((plan) => plan.tier > sessionStore.plan.tier);
</script>

<style scoped>
@layer components {
  .limit-reached__text {
    margin-bottom: 24px;

    &:deep(p:not(:last-child)) {
      margin-bottom: 8px;
    }
  }

  .limit-reached__plans {
    margin-bottom: 4px;
  }
}
</style>
