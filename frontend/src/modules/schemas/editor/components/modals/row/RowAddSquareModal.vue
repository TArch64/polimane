<template>
  <Modal title="Додати Рядок" save-button="Додати" @save="save">
    <NumberField
      label
      required
      :min="1"
      placeholder="Кількість Рядків"
      class="add-row__row-count"
      v-model="form.rows"
    />

    <NumberField
      label
      required
      :min="1"
      placeholder="Кількість Бісеринок"
      v-model="form.size"
    />
  </Modal>
</template>

<script setup lang="ts">
import { reactive } from 'vue';
import { useLocalStorage } from '@vueuse/core';
import type { ISchemaPattern } from '@/models';
import { Modal, useActiveModal } from '@/components/modal';
import { NumberField } from '@/components/form';
import { useRowsStore } from '@/modules/schemas/editor/stores';

const props = withDefaults(defineProps<{
  pattern: ISchemaPattern;
  toIndex?: number;
}>(), {
  toIndex: -1,
});

const modal = useActiveModal();
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

  modal.close(null);
}
</script>

<style scoped>
@layer page {
  .add-row__row-count {
    margin-bottom: 16px;
  }
}
</style>
