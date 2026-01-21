<template>
  <TextField
    :placeholder
    :label
    :required
    :variant
    :input-attrs="{ min, max }"
    type="number"
    v-model="modelString"
  />
</template>

<script setup lang="ts">
import { computed } from 'vue';
import type { ComponentVariant } from '@/types';
import TextField from './TextField.vue';

withDefaults(defineProps<{
  placeholder: string;
  label?: boolean;
  required?: boolean;
  variant?: ComponentVariant;
  min?: number;
  max?: number;
}>(), {
  label: false,
  variant: 'main',
});

const model = defineModel<number>({ required: true });

const modelString = computed({
  get: () => model.value.toString(),

  set: (value) => {
    const numberValue = Number(value);
    if (!isNaN(numberValue)) {
      model.value = numberValue;
    } else {
      model.value = 0;
    }
  },
});
</script>
