<template>
  <Modal
    title="Збереження як PDF"
    :width="ModalWidth.LG"
    :save-disabled="!hasColors"
    save-button="Зберегти"
    @save="save"
  >
    <SchemaExportCustomizer
      v-model:schema="schema"
      v-model:colors="colors"
      v-if="hasColors"
    />

    <SchemaExportPreview
      :schema
      :colors
      ref="previewRef"
    />
  </Modal>
</template>

<script setup lang="ts">
import { computed, ref, shallowRef, toRaw } from 'vue';
import type { ComponentExposed } from 'vue-component-type-helpers';
import { useEditorStore } from '@editor/stores';
import { Modal, ModalWidth, useActiveModal } from '@/components/modal';
import { useAsyncAction } from '@/composables';
import SchemaExportCustomizer from './SchemaExportCustomizer.vue';
import SchemaExportPreview from './SchemaExportPreview.vue';
import { saveSchemaPdf } from './saveSchemaPdf';
import { buildColorsModel } from './colorsModel';

const editorStore = useEditorStore();
const modal = useActiveModal();

const previewRef = ref<ComponentExposed<typeof SchemaExportPreview>>(null!);

const schema = shallowRef(toRaw(editorStore.schema));

const colors = ref(buildColorsModel(schema.value));
const hasColors = computed(() => !!colors.value.length);

const save = useAsyncAction(async () => {
  await saveSchemaPdf(schema.value, previewRef.value.getSource(), colors.value);
  modal.close(null);
});
</script>
