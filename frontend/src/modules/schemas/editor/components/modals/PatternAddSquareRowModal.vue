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
import { Modal, useActiveModal } from '@/components/modal';
import { NumberField } from '@/components/form';
import type { ISchemaPattern } from '@/models';
import { useRowsStore } from '../../stores';

const props = defineProps<{
  pattern: ISchemaPattern;
}>();

const modal = useActiveModal();
const rowsStore = useRowsStore(() => props.pattern);

const form = reactive({
  rows: 1,
  size: 1,
});

function save() {
  rowsStore.addSquareRow(form);
  modal.close();
}
</script>

<style scoped>
@layer page {
  .add-row__row-count {
    margin-bottom: 16px;
  }
}
</style>
