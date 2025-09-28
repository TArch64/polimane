<template>
  <Modal
    title="Збереження як PDF"
    :width="ModalWidth.LG"
    :save-disabled="!colors.length"
    save-button="Зберегти"
    @save="save"
  >
    <div class="export__colors-editor" v-if="colors.length">
      <ColorPicker
        v-for="(model, index) of colors"
        :key="model.initial"
        :label="`Колір ${index + 1}`"
        :model-value="model.current"
        @update:model-value="updateColor(index, $event)"
      />
    </div>

    <PreviewSchema class="export__preview" :schema />
  </Modal>
</template>

<script setup lang="ts">
import { reactive, shallowRef, toRaw } from 'vue';
import { useDebounceFn } from '@vueuse/core';
import { Modal, ModalWidth, useActiveModal } from '@/components/modal';
import { useAsyncAction } from '@/composables';
import { ColorPicker } from '@/components/form';
import type { SchemaBeadCoord } from '@/models';
import { PreviewSchema } from '../preview';
import { useEditorStore } from '../../stores';

const editorStore = useEditorStore();
const modal = useActiveModal();

interface IColorModel {
  initial: string;
  current: string;
}

function buildColorsModel(): IColorModel[] {
  const list = Object.values(editorStore.schema.beads);
  const colors = Array.from(new Set(list));

  return colors.map((color) => ({
    initial: color,
    current: color,
  }));
}

const colors = reactive(buildColorsModel());

const schema = shallowRef(toRaw(editorStore.schema));

const updateColor = useDebounceFn((index: number, color: string) => {
  const model = colors[index]!;
  const updatingSchema = structuredClone(schema.value);

  for (const coord_ in updatingSchema.beads) {
    const coord = coord_ as SchemaBeadCoord;
    if (updatingSchema.beads[coord] === model.current) {
      updatingSchema.beads[coord] = color;
    }
  }

  model.current = color;
  schema.value = updatingSchema;
}, 100);

const save = useAsyncAction(async () => {
  modal.close(null);
});
</script>

<style scoped>
@layer page {
  .export__colors-editor {
    display: flex;
    flex-direction: column;
    gap: 8px;
    margin-bottom: 16px;
  }

  .export__preview {
    width: 100%;
    height: auto;
    border-radius: var(--rounded-md);
  }
}
</style>
