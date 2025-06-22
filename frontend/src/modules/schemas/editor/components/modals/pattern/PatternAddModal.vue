<template>
  <Modal title="Додати Паттерн" save-button="Додати" @save="addPattern">
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
import { getPatternTitle, PatternKindValues, PatternType } from '@/enums';
import { usePatternsStore } from '@/modules/schemas/editor/stores';
import type { IAddingPattern } from './IAddingPattern';

const props = withDefaults(defineProps<{
  toIndex?: number;
}>(), {
  toIndex: -1,
});

const patternsStore = usePatternsStore();
const modal = useActiveModal<IAddingPattern>();

const selectedType = ref(PatternType.SQUARE);

const options = PatternKindValues.map((type): ISelectOption<PatternType> => ({
  value: type,
  label: getPatternTitle(type),
  disabled: type === PatternType.DIAMOND,
}));

function addPattern(): void {
  const pattern = patternsStore.createPattern(selectedType.value);
  modal.close({ pattern, toIndex: props.toIndex });
}
</script>

<style scoped>
@layer page {
  .add-pattern-modal__description {
    margin-bottom: 20px;
  }
}
</style>
