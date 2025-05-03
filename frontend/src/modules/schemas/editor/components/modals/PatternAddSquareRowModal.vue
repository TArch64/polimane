<template>
  <Modal title="Додати Рядок" save-button="Додати" @save="save">
    <NumberField
      required
      :min="1"
      placeholder="Кількість Елементів"
      v-model="size"
    />
  </Modal>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { Modal, useActiveModal } from '@/components/modal';
import { NumberField } from '@/components/form';
import type { ISchemaPattern } from '@/models';
import { useRowsStore } from '../../stores';

const props = defineProps<{
  pattern: ISchemaPattern;
}>();

const modal = useActiveModal();
const rowsStore = useRowsStore(() => props.pattern);

const size = ref(0);

function save() {
  rowsStore.addSquareRow({
    size: Number(size.value),
  });

  modal.close();
}
</script>
