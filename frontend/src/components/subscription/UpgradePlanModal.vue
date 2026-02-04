<template>
  <Modal
    title="Оновити Підписку"
    :footer="false"
    :width="ModalWidth.LG"
  >
    <SubscriptionPlans
      embedded
      @upgraded="modal.close(true)"
    />
  </Modal>
</template>

<script setup lang="ts">
import { Modal, ModalWidth, useActiveModal } from '@/components/modal';
import { usePlansStore } from '@/stores';
import SubscriptionPlans from './SubscriptionPlans.vue';

defineOptions({
  async beforeModalOpen(): Promise<void> {
    await usePlansStore().load();
  },
});

const modal = useActiveModal<boolean>();
</script>
