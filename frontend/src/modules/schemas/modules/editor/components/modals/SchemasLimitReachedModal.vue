<template>
  <LimitReachedModal title="Перевищено Ліміт Схем">
    <template #description="{ nextPlan }">
      <p>
        Ви досягли максимальної кількості схем для вашої підписки
      </p>

      <p>
        Щоб відновити редагування схем, будь ласка, видаліть {{ counter.overflowed }}
        {{ schemaPlural }}
        <template v-if="nextPlan">або оновіть підписку</template>
      </p>
    </template>
  </LimitReachedModal>
</template>

<script setup lang="ts">
import { SCHEMA_PLURAL, usePluralFormatter } from '@/composables';
import { usePlansStore } from '@/stores';
import { LimitReachedModal } from '@/components/subscription';
import { useSchemasCreatedCounter } from '@/composables/subscription';

defineOptions({
  async beforeModalOpen(): Promise<void> {
    await usePlansStore().load();
  },
});

const counter = useSchemasCreatedCounter();

const schemaPlural = usePluralFormatter(() => counter.overflowed, {
  ...SCHEMA_PLURAL,
  one: 'схему',
});
</script>
