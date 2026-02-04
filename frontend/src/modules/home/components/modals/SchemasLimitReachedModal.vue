<template>
  <LimitReachedModal title="Перевищено Ліміт Схем">
    <template #description="{ nextPlan }">
      <p>
        Ви досягли максимальної кількості схем для вашої підписки
      </p>

      <p>
        Щоб продовжити, будь ласка, видаліть {{ overflowCount }} {{ schemaPlural }}
        <template v-if="nextPlan">або оновіть підписку</template>
      </p>
    </template>
  </LimitReachedModal>
</template>

<script setup lang="ts">
import { SCHEMA_PLURAL, usePluralFormatter } from '@/composables';
import { usePlansStore } from '@/stores';
import { LimitReachedModal } from '@/components/subscription';

const props = defineProps<{
  overflowCount: number;
}>();

defineOptions({
  async beforeModalOpen(): Promise<void> {
    await usePlansStore().load();
  },
});

const schemaPlural = usePluralFormatter(() => props.overflowCount, {
  ...SCHEMA_PLURAL,
  one: 'схему',
});
</script>
