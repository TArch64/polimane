<template>
  <ColorPicker
    label="Змінити Колір Фону"
    class="export__background-color"
    :model-value="schema.backgroundColor"
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
import { useDebounceFn } from '@vueuse/core';
import { ColorPicker } from '@/components/form';
import { getBeadSettings, type ISchema, isRefBead, type SchemaContentBead } from '@/models';
import type { ISchemaColorModel } from './colorsModel';

const schema = defineModel<ISchema>('schema', { required: true });
const colors = defineModel<ISchemaColorModel[]>('colors', { required: true });

const updateBackgroundColor = useDebounceFn((color: string) => {
  schema.value = { ...schema.value, backgroundColor: color };
}, 0);

const updateColor = useDebounceFn((index: number, color: string) => {
  const colorModel = colors.value[index]!;
  const beads = structuredClone(schema.value.beads);

  for (const bead of Object.values(beads)) {
    if (isRefBead(bead)) {
      continue;
    }

    const settings = getBeadSettings(bead as SchemaContentBead);
    settings.color = color;
  }

  colorModel.current = color;
  schema.value = { ...schema.value, beads };
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
