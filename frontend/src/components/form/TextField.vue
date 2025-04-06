<template>
  <label class="text-field" :class="labelClasses">
    <input
      :type
      :placeholder
      class="text-field__input"
      v-model="model"
    >
  </label>
</template>

<script setup lang="ts">
import { computed } from 'vue';

const props = withDefaults(defineProps<{
  placeholder: string;
  type?: 'text' | 'password';
  variant?: 'main' | 'control';
}>(), {
  variant: 'main',
});

const model = defineModel<string>({ required: true });

const labelClasses = computed(() => `text-field--variant-${props.variant}`);
</script>

<style scoped>
@layer components {
  .text-field {
    border: 1px solid var(--color-divider);
    border-radius: var(--rounded-md);
    padding: 4px 8px;
    transition: border-color 0.15s ease-out;
    display: block;
    max-width: 100%;

    &:focus-within {
      border-color: var(--color-primary);
    }
  }

  .text-field--variant-main {
    background-color: var(--color-background-1);
  }

  .text-field--variant-control {
    background-color: var(--color-background-2);
  }

  .text-field__input {
    background-color: transparent;
    border: none;
    width: 100%;
    outline: none;
    font-size: var(--font-md);
    line-height: 20px;
  }
}
</style>
