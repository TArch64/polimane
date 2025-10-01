<template>
  <Button
    size="none"
    class="color-list__item"
    :class="classes"
    @click.stop="onClick"
    @dblclick.stop="onDblClick"
  >
    <input
      ref="pickerRef"
      type="color"
      class="color-list__item-picker"
      v-model="model"
    >
  </Button>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue';
import { Button } from '@/components/button';

const isActive = defineModel<boolean>('active', { required: true });

const model = defineModel<string>({
  required: true,

  set: (value) => {
    isActive.value = true;
    return value;
  },
});

const pickerRef = ref<HTMLInputElement>(null!);

const classes = computed(() => {
  const modifier = model.value ? 'value' : 'empty';

  return [
    { 'color-list__item--active': isActive.value },
    `color-list__item--${modifier}`,
  ];
});

function onClick(): void {
  if (!model.value) {
    pickerRef.value.click();
    return;
  }

  isActive.value = true;
}

function onDblClick(): void {
  if (model.value) pickerRef.value.click();
}
</script>

<style scoped>
@layer page {
  .color-list__item {
    padding: 0;
    border: var(--divider);
    transition: background-color 0.15s ease-out, border-color 0.15s ease-out;
    will-change: background-color, border-color;

    &:hover:not(.color-list__item--active) {
      border-color: var(--color-hover-divider);
    }
  }

  .color-list__item--active {
    outline: solid 1px v-bind("model");
    outline-offset: 1px;
  }

  .color-list__item--value {
    background-color: v-bind("model");
  }

  .color-list__item--empty {
    background-image: url("@/assets/emptyColor.svg");
    background-size: cover;
  }

  .color-list__item-picker {
    visibility: hidden;
    width: 100%;
    height: 100%;
  }
}
</style>
