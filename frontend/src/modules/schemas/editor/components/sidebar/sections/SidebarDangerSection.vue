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

    <!--    <div id="mypopover" popover anchor="mypopover" class="popover">-->
    <!--      Popover content-->
    <!--    </div>-->
  </SidebarSection>
</template>

<script setup lang="ts">
import { Button } from '@/components/button';
import { useConfirm } from '@/components/confirm';
import { useDomRef } from '@/composables';
import SidebarSection from './SidebarSection.vue';

const deleteButtonRef = useDomRef<HTMLElement>();

const deleteConfirm = useConfirm({
  anchorEl: deleteButtonRef.ref,
  message: 'Ви впевнені, що хочете видалити цю схему?',

  acceptButton: {
    text: 'Видалити',
    danger: true,
  },
});

async function deleteSchema(): Promise<void> {
  if (!await deleteConfirm.ask()) {
    return;
  }

  console.log('Schema deleted');
}
</script>

<style scoped>
@layer page {
  .sidebar-danger-section__button {
    width: 100%;
  }
}
</style>
