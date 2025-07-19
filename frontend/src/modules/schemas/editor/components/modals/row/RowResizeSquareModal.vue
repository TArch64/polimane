<template>
  <Modal title="Змінити Ширину Рядка" save-button="Змінити" @save="save">
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
import { nextTick, reactive } from 'vue';
import { Modal, useActiveModal } from '@/components/modal';
import type { ISchemaRow } from '@/models';
import { NumberField } from '@/components/form';
import { useRowsStore } from '@/modules/schemas/editor/stores';
import { getObjectParent } from '@/modules/schemas/editor/models';
import { useCanvasStage } from '@/modules/schemas/editor/composables';

const props = defineProps<{
  row: ISchemaRow;
}>();

const stage = useCanvasStage();
const modal = useActiveModal();
const rowsStore = useRowsStore(() => getObjectParent(props.row));

const form = reactive({
  size: props.row.content.length,
});

function save() {
  rowsStore.resizeRow(props.row, form.size);

  modal.close(null, async () => {
    await nextTick();
    const rowNode = stage.value.findOne(`#${props.row.id}`);
    rowNode?.fire('layoutUpdate');
  });
}
</script>
