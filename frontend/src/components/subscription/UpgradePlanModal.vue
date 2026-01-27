<template>
  <Modal
    title="Оновити Підписку"
    :footer="false"
    :width="ModalWidth.LG"
  >
    <SubscriptionPlans
      embedded
      class="upgrade-plan__plans"
      @upgraded="modal.close(null)"
    />
  </Modal>
</template>

<script setup lang="ts">
import { Modal, ModalWidth, useActiveModal } from '@/components/modal';
import SubscriptionPlans from './SubscriptionPlans.vue';

const modal = useActiveModal();
</script>

<style scoped>
@layer components {
  .upgrade-plan__plans {
    margin-bottom: 4px;
  }
}
</style>
