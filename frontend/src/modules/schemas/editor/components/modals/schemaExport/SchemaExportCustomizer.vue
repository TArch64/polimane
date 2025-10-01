<template>
  <ColorPicker
    label="Змінити Колір Фону"
    class="export__background-color"
    :model-value="model.backgroundColor"
    @update:model-value="updateBackgroundColor"
  />

  <p>
    Змінити кольори бісеринок:
  </p>

  <div class="export__colors-editor">
    <ColorPicker
      v-for="(model, index) of colors"
      :key="model.initial"
      :model-value="model.current"
      @update:model-value="updateColor(index, $event)"
    />
  </div>
</template>

<script setup lang="ts">
import { reactive } from 'vue';
import { useDebounceFn } from '@vueuse/core';
import { ColorPicker } from '@/components/form';
import type { ISchema, SchemaBeadCoord } from '@/models';
import { collectUniqColors } from './collectUniqColors';

const model = defineModel<ISchema>({ required: true });

interface IColorModel {
  initial: string;
  current: string;
}

function buildColorsModel(): IColorModel[] {
  return collectUniqColors(model.value).map((color) => ({
    initial: color,
    current: color,
  }));
}

const colors = reactive(buildColorsModel());

const updateBackgroundColor = useDebounceFn((color: string) => {
  model.value = { ...model.value, backgroundColor: color };
}, 0);

const updateColor = useDebounceFn((index: number, color: string) => {
  const colorModel = colors[index]!;
  const updatingSchema = structuredClone(model.value);

  for (const coord_ in updatingSchema.beads) {
    const coord = coord_ as SchemaBeadCoord;
    if (updatingSchema.beads[coord] === colorModel.current) {
      updatingSchema.beads[coord] = color;
    }
  }

  colorModel.current = color;
  model.value = updatingSchema;
}, 0);
</script>

<style scoped>
@layer page {
  .export__background-color {
    margin-bottom: 20px;
  }

  .export__colors-editor {
    display: flex;
    flex-wrap: wrap;
    gap: 8px;
    margin-top: 8px;
    margin-bottom: 16px;
  }
}
</style>
