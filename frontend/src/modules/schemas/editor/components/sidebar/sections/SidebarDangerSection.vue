<template>
  <SidebarSection>
    <Button
      danger
      size="md"
      variant="secondary"
      class="sidebar-danger-section__button"
      :ref="deleteButtonRef.templateRef"
      @click="deleteSchema"
    >
      Видалити Схему
    </Button>
  </SidebarSection>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';
import { Button } from '@/components/button';
import { useConfirm } from '@/components/confirm';
import { useAsyncAction, useDomRef } from '@/composables';
import { useEditorStore } from '../../../stores';
import SidebarSection from './SidebarSection.vue';

const router = useRouter();
const editorStore = useEditorStore();
const deleteButtonRef = useDomRef<HTMLElement>();

const deleteConfirm = useConfirm({
  anchorEl: deleteButtonRef.ref,
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
  .sidebar-danger-section__button {
    width: 100%;
  }
}
</style>
