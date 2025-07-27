<template>
  <div class="auth-factors-empty">
    <Button variant="primary" @click="addModal.open()">
      Увімкнути
    </Button>
  </div>
</template>

<script setup lang="ts">
import { Button } from '@/components/button';
import { useModal } from '@/components/modal';
import AuthFactorAddModal from './AuthFactorAddModal.vue';

const addModal = useModal(AuthFactorAddModal);
</script>

<style scoped>
@layer page {
  .auth-factors-empty {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    padding: 40px;
  }
}
</style>
