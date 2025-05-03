<template>
  <header class="editor-sidebar__header">
    <Button class="editor-sidebar__header-back" :to="{ name: 'home' }">
      <ArrowBackIcon />
      Едітор
    </Button>

    <Dropdown>
      <template #activator="{ open, activatorStyle }">
        <Button
          icon
          :style="mergeAnchorName(activatorStyle, deleteConfirm.anchorStyle)"
          @click="open"
        >
          <MoreHorizontalIcon />
        </Button>
      </template>

      <DropdownAction
        danger
        title="Видалити Схему"
        :icon="TrashIcon"
        @click="deleteSchema"
      />
    </Dropdown>
  </header>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';
import { Button } from '@/components/button';
import { ArrowBackIcon, MoreHorizontalIcon, TrashIcon } from '@/components/icon';
import { useEditorStore } from '@/modules/schemas/editor/stores';
import { useAsyncAction } from '@/composables';
import { useConfirm } from '@/components/confirm';
import { Dropdown, DropdownAction } from '@/components/dropdown';
import { mergeAnchorName } from '@/helpers';

const router = useRouter();
const editorStore = useEditorStore();

const deleteConfirm = useConfirm({
  danger: true,
  message: 'Ви впевнені, що хочете видалити цю схему?',
  acceptButton: 'Видалити',
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
    position: sticky;
    z-index: 10;
    top: 0;
  }

  .editor-sidebar__header-back {
    gap: 8px;
  }
}
</style>
