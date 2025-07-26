<template>
  <TextField
    :placeholder
    :label
    :required
    :variant
    :inputAttrs
    :type="fieldType"
    v-model="model"
    v-model:custom-error="customErrorModel"
  >
    <template #append>
      <Button icon size="sm" @click="isVisible = !isVisible">
        <ToggleIcon size="16" />
      </Button>
    </template>
  </TextField>
</template>

<script setup lang="ts">
import { computed, type InputHTMLAttributes, ref } from 'vue';
import { Button } from '../button';
import { EyeIcon, EyeOffIcon } from '../icon';
import TextField from './TextField.vue';

defineProps<{
  placeholder: string;
  label?: boolean;
  required?: boolean;
  variant?: 'main' | 'control';
  inputAttrs?: InputHTMLAttributes;
}>();

const model = defineModel<string>({ required: true });

const customErrorModel = defineModel<string>('customError', {
  required: false,
  default: '',
});

const isVisible = ref(false);
const fieldType = computed(() => isVisible.value ? 'text' : 'password');
const ToggleIcon = computed(() => isVisible.value ? EyeIcon : EyeOffIcon);
</script>
