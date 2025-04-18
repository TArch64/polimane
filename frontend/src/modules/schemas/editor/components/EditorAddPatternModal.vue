<template>
  <Modal
    title="Додати Паттерн"
    save-button="Додати"
    @save="addPattern"
  >
    <p class="add-pattern-modal__description">
      Виберіть паттерн, який ви хочете додати до вашої схеми
    </p>

    <RadioSelect
      required
      name="pattern-type"
      :options
      v-model="selectedType"
    />
  </Modal>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { Modal, useActiveModal } from '@/components/modal';
import { type ISelectOption, RadioSelect } from '@/components/form';
import { getPatternTitle, PatternType, PatternTypeValues } from '@/enums';
import { usePatternsStore } from '../stores';

const patternsStore = usePatternsStore();
const modal = useActiveModal();

const selectedType = ref(PatternType.SQUARE);

const options = PatternTypeValues.map((type): ISelectOption<PatternType> => ({
  value: type,
  label: getPatternTitle(type),
}));

function addPattern(): void {
  patternsStore.addPattern(selectedType.value);
  modal.close();
}
</script>

<style scoped>
@layer page {
  .add-pattern-modal__description {
    margin-bottom: 20px;
  }
}
</style>
