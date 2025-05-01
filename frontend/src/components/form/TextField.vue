<template>
  <label class="text-field">
    <span class="text-field__label" v-if="label">
      {{ placeholder }}
    </span>

    <span class="text-field__container" :class="containerClasses">
      <input
        :type
        :placeholder
        :required
        ref="inputRef"
        class="text-field__input"
        v-bind="inputAttrs"
        v-model="model"
      >
    </span>
  </label>
</template>

<script setup lang="ts">
import { computed, type InputHTMLAttributes, ref } from 'vue';

const props = withDefaults(defineProps<{
  placeholder: string;
  label?: boolean;
  required?: boolean;
  type?: 'text' | 'password' | 'number';
  variant?: 'main' | 'control';
  inputAttrs?: InputHTMLAttributes;
}>(), {
  label: false,
  type: 'text',
  variant: 'main',
  inputAttrs: () => ({}),
});

const model = defineModel<string>({
  required: true,

  set: (value) => {
    if (props.required && !value) {
      inputRef.value.reportValidity();
    }
    return value;
  },
});

const inputRef = ref<HTMLInputElement>(null!);

const containerClasses = computed(() => `text-field__container--variant-${props.variant}`);
</script>

<style scoped>
@layer components {
  .text-field {
    display: block;
    max-width: 100%;
  }

  .text-field__label {
    display: block;
    margin-bottom: 2px;
    margin-left: 2px;
    font-size: var(--font-xs);
    color: color-mix(in srgb, var(--color-primary), transparent 40%);
  }

  .text-field__container {
    border: var(--divider);
    border-radius: var(--rounded-md);
    padding: 4px 8px;
    transition: border-color 0.15s ease-out;
    will-change: border-color;
    display: block;
    max-width: 100%;

    &:has(:focus:not(:user-invalid)) {
      border-color: var(--color-hover-divider);
    }

    &:has(:user-invalid) {
      border-color: var(--color-danger);
    }
  }

  .text-field__container--variant-main {
    background-color: var(--color-background-1);
  }

  .text-field__container--variant-control {
    background-color: var(--color-background-2);
  }

  .text-field__input {
    background-color: transparent;
    border: none;
    width: 100%;
    outline: none;
    font-size: calc(var(--font-md) - 1px);
    line-height: 20px;
  }
}
</style>
