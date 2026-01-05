<template>
  <LabeledContent :label v-if="label">
    <Picker />
  </LabeledContent>

  <Picker v-else />
</template>

<script setup lang="ts">
import { computed, type FunctionalComponent, h, withDirectives } from 'vue';
import { vTappable } from '@/directives';
import LabeledContent from './LabeledContent.vue';

withDefaults(defineProps<{
  label?: string;
}>(), {
  label: '',
});

const model = defineModel<string>({ required: true });

const controlClasses = computed(() => {
  const modifier = model.value ? 'value' : 'empty';
  return `color-picker__control--${modifier}`;
});

const Picker: FunctionalComponent = () => withDirectives(
  h('span', {
    class: ['color-picker__control', controlClasses.value],
  }, [
    h('input', {
      type: 'color',
      class: 'color-picker__input',
      value: model.value,

      onInput: (event: Event) => {
        model.value = (event.target as HTMLInputElement).value;
      },
    }),
  ]),
  [[vTappable]],
);
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

  :deep(.color-picker__input) {
    opacity: 0;
    width: 100%;
    height: 100%;
    cursor: inherit;
  }
}
</style>
