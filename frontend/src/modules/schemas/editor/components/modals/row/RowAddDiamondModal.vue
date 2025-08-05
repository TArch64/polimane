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
      class="add-row__field"
      v-model="form.size"
    />

    <NumberField
      label
      required
      :min="1"
      :max="5"
      variant="control"
      placeholder="Ширина Сторони"
      v-model="form.sideSize"
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

const modal = useActiveModal<boolean>();
const rowsStore = useRowsStore(() => props.pattern);

const lastForm = useLocalStorage('schema-editor-row-add-diamond-form', {
  rows: 1,
  size: 50,
  sideSize: 1,
});

const form = reactive({ ...lastForm.value });

function save() {
  lastForm.value = { ...form };

  rowsStore.addDiamondRow({
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
