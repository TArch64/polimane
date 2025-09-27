<template>
  <LabeledContent :label>
    <span
      class="color-picker__control"
      :class="controlClasses"
    >
      <input
        type="color"
        class="color-picker__input"
        v-model="model"
      />
    </span>
  </LabeledContent>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import LabeledContent from './LabeledContent.vue';

defineProps<{
  label: string;
}>();

const model = defineModel<string>({ required: true });

const controlClasses = computed(() => {
  const modifier = model.value ? 'value' : 'empty';
  return `color-picker__control--${modifier}`;
});
</script>

<style scoped>
@layer components {
  .color-picker__control {
    width: 24px;
    height: 24px;
    border: var(--divider);
    border-radius: var(--rounded-md);
    cursor: pointer;
  }

  .color-picker__control--value {
    background-color: v-bind("model");
  }

  .color-picker__control--empty {
    background-image: url("@/assets/emptyColor.svg");
    background-size: cover;
  }

  .color-picker__input {
    visibility: hidden;
    width: 100%;
    height: 100%;
  }
}
</style>
