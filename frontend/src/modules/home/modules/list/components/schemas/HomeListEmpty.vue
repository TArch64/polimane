<template>
  <div class="schemas-empty">
    <p class="schemas-empty__description">
      Поки що не створено жодної схеми для бісеру
    </p>

    <Button variant="primary" @click="createModal.open()">
      Додати Схему
    </Button>
  </div>
</template>

<script setup lang="ts">
import { Button } from '@/components/button';
import { useModal } from '@/components/modal';
import { SchemaCreateModal } from '@/modules/home/components/modals';

const createModal = useModal(SchemaCreateModal);
</script>

<style scoped>
@layer page {
  .schemas-empty {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    flex-grow: 1;
  }

  .schemas-empty__description {
    margin-bottom: 12px;
  }
}
</style>
