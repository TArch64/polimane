<template>
  <Button
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

const model = defineModel<string>({ required: true });

const pickerRef = ref<HTMLInputElement>(null!);

const classes = computed(() => {
  const modifier = model.value ? 'value' : 'empty';
  return `color-list__item--${modifier}`;
});

function onClick(): void {
  if (!model.value) {
    pickerRef.value.click();
    return;
  }
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
  }

  .color-list__item--value {
    background-color: v-bind("model");
  }

  .color-list__item--empty {
    background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='100%25' height='100%25'%3E%3Crect width='100%25' height='100%25' fill='%23f8f8f8'/%3E%3Cline x1='0' y1='0' x2='100%25' y2='100%25' stroke='%23dddddd' stroke-width='1.5'/%3E%3C/svg%3E");
    background-size: cover;
  }

  .color-list__item-picker {
    visibility: hidden;
    width: 100%;
    height: 100%;
  }
}
</style>
