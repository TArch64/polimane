<template>
  <Modal title="Додати Рядок" save-button="Додати" @save="save">
    <NumberField
      required
      placeholder="Кількість Елементів"
      v-model="size"
    />
  </Modal>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { Modal, useActiveModal } from '@/components/modal';
import { useRowsStore } from '@/modules/schemas/editor/stores';
import { NumberField } from '@/components/form';
import type { ISchemaPattern } from '@/models';

const props = defineProps<{
  pattern: ISchemaPattern;
}>();

const modal = useActiveModal();
const rowsStore = useRowsStore();

const size = ref(0);

function save() {
  rowsStore.addSquareRow(props.pattern, {
    size: Number(size.value),
  });

  modal.close();
}
</script>

<style scoped>
@layer page {

}
</style>
