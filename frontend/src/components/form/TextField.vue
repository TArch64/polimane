<template>
  <label class="text-field">
    <input
      :type
      :placeholder
      class="text-field__input"
      v-model="model"
    >
  </label>
</template>

<script setup lang="ts">
defineProps<{
  placeholder: string;
  type?: 'text' | 'password';
}>();

const model = defineModel<string>({ required: true });
</script>

<style scoped>
@layer components {
  .text-field {
    background-color: var(--color-background-1);
    border: 1px solid var(--color-divider);
    border-radius: var(--rounded-md);
    padding: 8px;
    transition: border-color 0.15s ease-out;

    &:focus-within {
      border-color: var(--color-primary);
    }
  }

  .text-field__input {
    background-color: transparent;
    border: none;
    width: 100%;
    outline: none;
  }
}
</style>
