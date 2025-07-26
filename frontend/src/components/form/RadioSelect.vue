<template>
  <div class="radio-select">
    <label
      v-for="option of options"
      :key="option.value"
      class="radio-select__option"
    >
      <input
        class="radio-select__input"
        type="radio"
        :name
        :required
        :value="option.value"
        :disabled="option.disabled"
        :checked="option.value === model"
        @change="model = option.value"
      />

      <span class="radio-select__label">
        {{ option.label }}
      </span>
    </label>
  </div>
</template>

<script setup lang="ts" generic="V extends string">
import type { ISelectOption } from './ISelectOption';

withDefaults(defineProps<{
  name: string;
  options: ISelectOption<V>[];
  required?: boolean;
}>(), {
  required: false,
});

const model = defineModel<V>({ required: true });
</script>

<style scoped>
@layer components {
  .radio-select {
    display: flex;
    flex-direction: column;
    gap: 8px;
  }

  .radio-select__option {
    display: flex;
    gap: 8px;

    &:has(input:disabled) {
        color: var(--color-text-3);
    }
  }

  .radio-select__input {
    accent-color: var(--color-primary);
  }
}
</style>
