<template>
  <Modal
    title="Додати Паттерн"
    save-button="Додати"
    @save="addPattern"
  >
    <template #activator="ctx">
      <slot v-bind="ctx" />
    </template>

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
import { Modal, type ModalActivatorSlot } from '@/components/modal';
import { type ISelectOption, RadioSelect } from '@/components/form';
import { PatternType } from '@/models';
import { usePatternsStore } from '../stores';

defineSlots<{
  default: ModalActivatorSlot;
}>();

const patternsStore = usePatternsStore();

const selectedType = ref(PatternType.SQUARE);

const options: ISelectOption<PatternType>[] = [
  {
    value: PatternType.SQUARE,
    label: 'Квадратна Сітка',
  },
  {
    value: PatternType.DIAMOND,
    label: 'Ромбова Сітка',
  },
];

function addPattern(): void {
  patternsStore.addPattern(selectedType.value);
}
</script>

<style scoped>
@layer page {
  .add-pattern-modal__description {
    margin-bottom: 20px;
  }
}
</style>
