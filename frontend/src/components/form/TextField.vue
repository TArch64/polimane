<template>
  <FormField :label :placeholder :variant>
    <input
      :type
      :placeholder
      :required
      :disabled
      ref="inputRef"
      class="text-field__input"
      @blur="onBlur"
      v-bind="inputAttrs"
      v-model="model"
    >

    <template #append v-if="slots.append">
      <slot name="append" />
    </template>
  </FormField>
</template>

<script setup lang="ts">
import { type InputHTMLAttributes, ref, type Slot } from 'vue';
import FormField from './FormField.vue';

withDefaults(defineProps<{
  placeholder: string;
  label?: boolean;
  required?: boolean;
  disabled?: boolean;
  type?: 'text' | 'password' | 'number' | 'email';
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
    if (isDirty.value && inputRef.value) {
      inputRef.value.validity.customError
        ? inputRef.value.setCustomValidity('')
        : inputRef.value.reportValidity();
    }
    return value;
  },
});

const slots = defineSlots<{
  append?: Slot;
}>();

const inputRef = ref<HTMLInputElement | null>(null);
const isDirty = ref(false);

function onBlur() {
  isDirty.value = true;
  inputRef.value?.reportValidity();
}

function focus() {
  inputRef.value?.focus();
}

function blur() {
  inputRef.value?.blur();
}

function setError(message: string) {
  if (inputRef.value) {
    inputRef.value.focus();
    inputRef.value.setCustomValidity(message);
    setTimeout(() => inputRef.value?.reportValidity());
  }
}

defineExpose({
  focus,
  blur,
  setError,
});
</script>

<style scoped>
@layer components {
  .text-field__input {
    background-color: transparent;
    border: none;
    flex-grow: 1;
    outline: none;
    font-size: calc(var(--font-md) - 1px);
    line-height: 20px;
  }
}
</style>
