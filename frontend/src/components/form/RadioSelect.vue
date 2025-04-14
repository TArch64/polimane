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
        :checked="option.value === model"
        @change="model = option.value"
      />

      <p>
        {{ option.label }}
      </p>
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
  }

  .radio-select__input {
    accent-color: var(--color-primary);
  }
}
</style>
