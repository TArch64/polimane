<template>
  <LabeledContent :label gap="12">
    <input
      type="checkbox"
      class="checkbox__input tap-animation"
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
  .checkbox__input {
    accent-color: var(--color-primary);
    width: 16px;
    height: 16px;
    cursor: pointer;
    --tap-scale: 0.9;
  }
}
</style>
