<template>
  <main class="editor-empty">
    <p>
      Додати першу схему
    </p>

    <EditorAddSchemaModal v-slot="{ open }">
      <Button variant="primary" size="md" @click="open">
        Додати
      </Button>
    </EditorAddSchemaModal>
  </main>
</template>

<script setup lang="ts">

import { Button } from '@/components/button';
import EditorAddSchemaModal from './EditorAddSchemaModal.vue';
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
