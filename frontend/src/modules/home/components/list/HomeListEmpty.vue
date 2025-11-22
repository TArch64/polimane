<template>
  <div class="home-list-empty">
    <p class="home-list-empty__description">
      {{ description }}
    </p>

    <Button variant="primary" @click="createModal.open()">
      Додати Схему
    </Button>
  </div>
</template>

<script setup lang="ts">
import { Button } from '@/components/button';
import { useModal } from '@/components/modal';
import { SchemaCreateModal } from '../modals';

defineProps<{
  description: string;
}>();

const createModal = useModal(SchemaCreateModal);
</script>

<style scoped>
@layer page {
  .home-list-empty {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    flex-grow: 1;
  }

  .home-list-empty__description {
    margin-bottom: 12px;
  }
}
</style>
