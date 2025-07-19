<template>
  <Modal title="Змінити Ширину Рядка" save-button="Додати" @save="save">
    <NumberField
      label
      required
      :min="row.content.length"
      placeholder="Кількість Бісеринок"
      v-model="form.size"
    />
  </Modal>
</template>

<script setup lang="ts">
import { reactive } from 'vue';
import { Modal, useActiveModal } from '@/components/modal';
import type { ISchemaRow } from '@/models';
import { NumberField } from '@/components/form';
import { useRowsStore } from '@/modules/schemas/editor/stores';
import { getObjectParent } from '@/modules/schemas/editor/models';

const props = defineProps<{
  row: ISchemaRow;
}>();

const modal = useActiveModal();
const rowsStore = useRowsStore(() => getObjectParent(props.row));

const form = reactive({
  size: props.row.content.length,
});

function save() {
  rowsStore.resizeRow(props.row, form.size);
  modal.close(null);
}
</script>
