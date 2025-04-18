<template>
  <div class="schemas-empty">
    <p class="schemas-empty__description">
      Поки що не створено жодної схеми для бісеру
    </p>

    <Button variant="primary" @click="createSchemaModal.open()">
      Створити Схему
    </Button>
  </div>
</template>

<script setup lang="ts">
import { Button } from '@/components/button';
import { useModal } from '@/components/modal';
import HomeCreateSchemaModal from './HomeCreateSchemaModal.vue';

const createSchemaModal = useModal(HomeCreateSchemaModal);
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
