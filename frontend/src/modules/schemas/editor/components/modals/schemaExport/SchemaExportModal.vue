<template>
  <Modal
    title="Збереження як PDF"
    :width="ModalWidth.LG"
    :save-disabled="!hasColors"
    save-button="Зберегти"
    @save="save"
  >
    <SchemaExportCustomizer
      v-model="schema"
      v-if="hasColors"
    />

    <SchemaExportPreview ref="previewRef" :schema />
  </Modal>
</template>

<script setup lang="ts">
import { computed, ref, shallowRef, toRaw } from 'vue';
import type { ComponentExposed } from 'vue-component-type-helpers';
import { Modal, ModalWidth, useActiveModal } from '@/components/modal';
import { useAsyncAction } from '@/composables';
import { useEditorStore } from '../../../stores';
import SchemaExportCustomizer from './SchemaExportCustomizer.vue';
import SchemaExportPreview from './SchemaExportPreview.vue';
import { saveSchemaPdf } from './saveSchemaPdf';

const editorStore = useEditorStore();
const modal = useActiveModal();

const previewRef = ref<ComponentExposed<typeof SchemaExportPreview>>(null!);

const schema = shallowRef(toRaw(editorStore.schema));
const hasColors = computed(() => !!Object.keys(editorStore.schema.beads).length);

const save = useAsyncAction(async () => {
  await saveSchemaPdf(schema.value, previewRef.value.getSource());
  modal.close(null);
});
</script>
