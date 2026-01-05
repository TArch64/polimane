<template>
  <LabeledContent :label gap="12">
    <input
      type="checkbox"
      class="checkbox__input"
      v-model="model"
      v-tappable
    >
  </LabeledContent>
</template>

<script setup lang="ts">
import { vTappable } from '@/directives';
import LabeledContent from './LabeledContent.vue';

defineProps<{
  label: string;
}>();

const model = defineModel<boolean>({ required: true });
</script>

<style scoped>
@layer components {
  .checkbox__input {
    accent-color: var(--color-primary);
    width: 16px;
    height: 16px;
    cursor: pointer;
  }
}
</style>
