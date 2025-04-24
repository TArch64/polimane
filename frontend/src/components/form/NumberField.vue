<template>
  <TextField
    :placeholder
    :label
    :required
    :variant
    type="number"
    v-model="modelString"
  />
</template>

<script setup lang="ts">
import { computed } from 'vue';
import TextField from './TextField.vue';

withDefaults(defineProps<{
  placeholder: string;
  label?: boolean;
  required?: boolean;
  variant?: 'main' | 'control';
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
