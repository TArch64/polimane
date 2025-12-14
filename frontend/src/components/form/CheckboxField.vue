<template>
  <LabeledContent :label gap="12">
    <input
      type="checkbox"
      class="checkbox"
      v-model="model"
    >
  </LabeledContent>
</template>

<script setup lang="ts">
import LabeledContent from './LabeledContent.vue';

defineProps<{
  label: string;
}>();

const model = defineModel<boolean>({ required: true });
</script>

<style scoped>
@layer components {
  .checkbox {
    accent-color: var(--color-primary);
    width: 16px;
    height: 16px;
    cursor: pointer;
  }
}
</style>
