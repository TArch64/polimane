<template>
  <main class="editor-empty">
    <p>
      Додайте перший паттерн до вашої схеми
    </p>

    <Button variant="primary" @click="addPatternModal.open()">
      Додати
    </Button>
  </main>
</template>

<script setup lang="ts">
import { Button } from '@/components/button';
import { useModal } from '@/components/modal';
import { PatternAddModal } from './modals';

const addPatternModal = useModal(PatternAddModal);
</script>

<style scoped>
@layer page {
  .editor-empty {
    display: flex;
    align-items: center;
    justify-content: center;
    flex-direction: column;
    gap: 12px;
  }
}
</style>
