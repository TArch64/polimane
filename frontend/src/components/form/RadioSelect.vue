<template>
  <ul class="radio-select">
    <li class="radio-select__option" v-for="option of options" :key="option.value">
      <LabeledContent
        gap="12"
        :label="option.label"
        class="radio-select__option-label"
        :class="getOptionLabelClasses(option)"
      >
        <input
          type="radio"
          class="radio-select__input"
          :name="id"
          :value="option.value"
          :disabled="option.disabled"
          v-model="model"
        >
      </LabeledContent>

      <VerticalSlideTransition :duration="150">
        <slot
          :name="`option-${option.value}`"
          v-if="isSlotVisible(option)"
        />
      </VerticalSlideTransition>
    </li>
  </ul>
</template>

<script setup lang="ts" generic="V extends SelectValue">
import { type Slot, useId } from 'vue';
import { VerticalSlideTransition } from '@/components/transition';
import LabeledContent from './LabeledContent.vue';
import type { ISelectOption, SelectOptions, SelectValue } from './ISelectOption';

defineProps<{
  options: SelectOptions<V>;
}>();

const model = defineModel<V>({ required: true });

type OptionSlots = Record<`option-${string}`, Slot>;
const slots = defineSlots<OptionSlots>();

const id = useId();

function isSlotVisible(option: ISelectOption<V>) {
  return !!slots[`option-${option.value}`] && model.value === option.value;
}

const getOptionLabelClasses = (option: ISelectOption<V>) => ({
  'radio-select__option-label--with-slot': isSlotVisible(option),
});
</script>

<style scoped>
@layer components {
  .radio-select {
    list-style-type: none;
    padding: 0 0 0 4px;
    margin: 0;
    display: flex;
    flex-direction: column;
    gap: 8px;
  }

  .radio-select__input {
    accent-color: var(--color-primary);
    width: 16px;
    height: 16px;
    cursor: pointer;
  }

  .radio-select__option,
  .radio-select__option-label {
    transition: margin 150ms ease-out;
    will-change: margin;
  }

  .radio-select__option:has(.radio-select__option-label--with-slot) + .radio-select__option {
    margin-top: 8px;
  }

  .radio-select__option-label--with-slot {
    margin-bottom: 8px;
  }
}
</style>
