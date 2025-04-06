<template>
  <div class="schemas-empty">
    <p class="schemas-empty__description">
      Поки що не створено жодної схеми для бісеру
    </p>

    <HomeCreateSchemaModal v-slot="{ open }">
      <Button variant="primary" size="lg" @click="open">
        Створити Схему
      </Button>
    </HomeCreateSchemaModal>
  </div>
</template>

<script setup lang="ts">
import { Button } from '@/components/button';
import HomeCreateSchemaModal from './HomeCreateSchemaModal.vue';
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
