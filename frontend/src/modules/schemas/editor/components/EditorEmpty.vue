<template>
  <main class="editor-empty">
    <p>
      Додати першу схему
    </p>

    <Button variant="primary" size="md">
      Додати
    </Button>
  </main>
</template>

<script setup lang="ts">

import { Button } from '@/components/button';
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
