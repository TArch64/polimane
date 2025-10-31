<template>
  <Button
    icon
    size="md"
    variant="secondary"
    title="Видалити обрані схеми"
    :style="deleteConfirm.anchorStyle"
    @click="deleteIntent"
  >
    <TrashIcon />
  </Button>
</template>

<script setup lang="ts">
import { Button } from '@/components/button';
import { TrashIcon } from '@/components/icon';
import { useConfirm } from '@/components/confirm';
import { useSchemasStore } from '@/modules/home/stores';
import { useAsyncAction } from '@/composables';

const schemasStore = useSchemasStore();

const deleteConfirm = useConfirm({
  danger: true,
  control: false,
  message: `Ви впевнені, що хочете видалити ${schemasStore.selected.size} схеми?`,
  acceptButton: 'Видалити',
});

const deleteSchemas = useAsyncAction(async () => {
  await schemasStore.deleteMany(schemasStore.selected);
  schemasStore.clearSelection();
});

async function deleteIntent(): Promise<void> {
  if (await deleteConfirm.ask()) {
    await deleteSchemas();
  }
}
</script>
