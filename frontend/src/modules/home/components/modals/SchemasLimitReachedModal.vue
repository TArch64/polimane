<template>
  <Modal title="Перевищено Ліміт Схем" :footer="false">
    <p class="schemas-limit-reached__text">
      Ви досягли максимальної кількості схем для вашої підписки
    </p>

    <p class="schemas-limit-reached__text">
      Щоб {{ actionTitle }}, будь ласка, видаліть {{ overflowCount }} {{ schemaPlural }}
      <template v-if="nextPlan">або оновіть підписку</template>
    </p>

    <UpgradingPlan
      :plan="nextPlan"
      v-if="nextPlan"
    />
  </Modal>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { Modal } from '@/components/modal';
import { SCHEMA_PLURAL, usePluralFormatter } from '@/composables';
import { usePlansStore, useSessionStore } from '@/stores';
import { UpgradingPlan } from '@/components/subscription';

const props = defineProps<{
  actionTitle: string;
  overflowCount: number;
}>();

defineOptions({
  async beforeModalOpen(): Promise<void> {
    await usePlansStore().load();
  },
});

const sessionStore = useSessionStore();
const plansStore = usePlansStore();

const schemaPlural = usePluralFormatter(() => props.overflowCount, {
  ...SCHEMA_PLURAL,
  one: 'схему',
});

const nextPlan = computed(() => {
  return plansStore.plans.find((plan) => plan.tier > sessionStore.plan.tier);
});
</script>

<style scoped>
@layer page {
  .schemas-limit-reached__text {
    margin-bottom: 8px;
  }
}
</style>
