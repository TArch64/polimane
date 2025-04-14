<template>
  <main class="editor-empty">
    <p>
      Додайте перший паттерн до вашої схеми
    </p>

    <EditorAddPatternModal v-slot="{ open }">
      <Button variant="primary" size="md" @click="open">
        Додати
      </Button>
    </EditorAddPatternModal>
  </main>
</template>

<script setup lang="ts">

import { Button } from '@/components/button';
import EditorAddPatternModal from './EditorAddPatternModal.vue';
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
