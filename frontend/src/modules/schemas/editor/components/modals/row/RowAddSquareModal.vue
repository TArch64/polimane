<template>
  <Modal title="Додати Рядок" save-button="Додати" @save="save">
    <NumberField
      label
      required
      :min="1"
      variant="control"
      placeholder="Кількість Рядків"
      class="add-row__field"
      v-model="form.rows"
    />

    <NumberField
      label
      required
      :min="1"
      variant="control"
      placeholder="Ширина Рядка"
      v-model="form.size"
    />
  </Modal>
</template>

<script setup lang="ts">
import { reactive } from 'vue';
import { useLocalStorage } from '@vueuse/core';
import type { SchemaPattern } from '@/models';
import { Modal, useActiveModal } from '@/components/modal';
import { NumberField } from '@/components/form';
import { useRowsStore } from '@/modules/schemas/editor/stores';

const props = withDefaults(defineProps<{
  pattern: SchemaPattern;
  toIndex?: number;
}>(), {
  toIndex: -1,
});

const modal = useActiveModal<boolean>();
const rowsStore = useRowsStore(() => props.pattern);

const lastForm = useLocalStorage('schema-editor-row-add-square-form', {
  rows: 4,
  size: 50,
});

const form = reactive({ ...lastForm.value });

function save() {
  lastForm.value = { ...form };

  rowsStore.addSquareRow({
    ...form,
    toIndex: props.toIndex,
  });

  modal.close(true);
}
</script>

<style scoped>
@layer page {
  .add-row__field {
    margin-bottom: 16px;
  }
}
</style>
