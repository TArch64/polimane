<template>
  <FormField :variant :placeholder :label class="search-field">
    <select class="select-field__input" v-model="model">
      <option :value="null" disabled hidden v-if="placeholder">
        {{ placeholder }}
      </option>

      <slot />
    </select>
  </FormField>
</template>

<script setup lang="ts" generic="V">
import FormField from './FormField.vue';

withDefaults(defineProps<{
  placeholder?: string;
  label?: boolean;
  required?: boolean;
  disabled?: boolean;
  variant?: 'main' | 'control';
}>(), {
  placeholder: '',
  label: false,
  variant: 'main',
});

const model = defineModel<V>({ required: true });
</script>

<style scoped>
@layer components {
  .search-field :deep(.form-field__container) {
    padding: 0 8px 0 0;
    overflow: clip;
  }

  .select-field__input {
    padding: 4px 0 4px 8px;
    border: 0;
    outline: none;
    background-color: transparent;
  }
}
</style>
