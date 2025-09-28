<template>
  <Modal
    title="Збереження як PDF"
    :width="ModalWidth.LG"
    :save-disabled="!colors.length"
    save-button="Зберегти"
    @save="save"
  >
    <template v-if="colors.length">
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

    <PreviewSchema
      ref="previewRef"
      class="export__preview"
      :schema
    />
  </Modal>
</template>

<script setup lang="ts">
import { reactive, shallowRef, toRaw } from 'vue';
import { useDebounceFn } from '@vueuse/core';
import type { jsPDF } from 'jspdf';
import { Modal, ModalWidth, useActiveModal } from '@/components/modal';
import { useAsyncAction, useDomRef } from '@/composables';
import { ColorPicker } from '@/components/form';
import type { SchemaBeadCoord } from '@/models';
import { PreviewSchema } from '../preview';
import { useEditorStore } from '../../stores';

const previewRef = useDomRef<SVGSVGElement>();

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

const updateBackgroundColor = useDebounceFn((color: string) => {
  schema.value = { ...schema.value, backgroundColor: color };
}, 0);

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
}, 0);

async function insertSchema(pdf: jsPDF): Promise<void> {
  const pageWidth = pdf.internal.pageSize.getWidth();
  const pageHeight = pdf.internal.pageSize.getHeight();

  pdf.setFillColor(schema.value.backgroundColor);
  pdf.rect(0, 0, pageWidth, pageHeight, 'F');

  const margin = 10;
  const availableWidth = pageWidth - (margin * 2);
  const availableHeight = pageHeight - (margin * 2);

  const svgEl = previewRef.value.cloneNode(true) as SVGSVGElement;
  const viewBox = svgEl.getAttribute('viewBox')!.split(' ');
  const svgWidth = parseFloat(viewBox[2]!);
  const svgHeight = parseFloat(viewBox[3]!);
  const aspectRatio = svgWidth / svgHeight;

  let width = availableWidth;
  let height = availableWidth / aspectRatio;

  if (height > availableHeight) {
    height = availableHeight;
    width = availableHeight * aspectRatio;
  }

  const x = margin + (availableWidth - width) / 2;
  const y = margin + (availableHeight - height) / 2;

  await pdf.svg(svgEl, {
    x: x,
    y: y,
    width: width,
    height: height,
  });
}

async function savePDF(): Promise<void> {
  const { default: JSPDF } = await import('jspdf');
  await import('svg2pdf.js');

  const pdf = new JSPDF({
    orientation: 'landscape',
    unit: 'mm',
    format: 'a4',
  });

  await insertSchema(pdf);
  await pdf.save(`${schema.value.name}.pdf`, { returnPromise: true });
}

const save = useAsyncAction(async () => {
  await savePDF();
  modal.close(null);
});
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

  .export__preview {
    width: 100%;
    height: auto;
    border-radius: var(--rounded-md);
  }
}
</style>
