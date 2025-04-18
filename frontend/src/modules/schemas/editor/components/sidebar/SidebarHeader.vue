<template>
  <header class="editor-sidebar__header">
    <Button
      class="editor-sidebar__header-back"
      variant="secondary"
      size="md"
      :to="{ name: 'home' }"
    >
      <ArrowBackIcon />
      Редактор
    </Button>

    <Button
      icon
      danger
      size="md"
      variant="secondary"
      :ref="deleteConfirm.anchorRef"
      @click="deleteSchema"
    >
      <TrashIcon />
    </Button>
  </header>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';
import { Button } from '@/components/button';
import { ArrowBackIcon, TrashIcon } from '@/components/icon';
import { useEditorStore } from '@/modules/schemas/editor/stores';
import { useAsyncAction } from '@/composables';
import { useConfirm } from '@/components/confirm';

const router = useRouter();
const editorStore = useEditorStore();

const deleteConfirm = useConfirm({
  message: 'Ви впевнені, що хочете видалити цю схему?',

  acceptButton: {
    text: 'Видалити',
    danger: true,
  },
});

const deleteSchema = useAsyncAction(async () => {
  if (await deleteConfirm.ask()) {
    await editorStore.deleteSchema();
    await router.push({ name: 'home' });
  }
});
</script>

<style scoped>
@layer page {
  .editor-sidebar__header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 6px 8px;
    background-color: var(--color-background-1);
    border-bottom: 1px solid var(--color-divider);
  }

  .editor-sidebar__header-back {
    gap: 8px;
  }
}
</style>
