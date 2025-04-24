<template>
  <main class="editor-empty">
    <Button class="editor-empty__back" :to="{ name: 'home' }">
      <ArrowBackIcon />
      Редактор
    </Button>

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
import { ArrowBackIcon } from '@/components/icon';
import { EditorAddPatternModal } from './modals';

const addPatternModal = useModal(EditorAddPatternModal);
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

  .editor-empty__back {
    position: fixed;
    top: 4px;
    left: 8px;
    gap: 8px;
  }
}
</style>
