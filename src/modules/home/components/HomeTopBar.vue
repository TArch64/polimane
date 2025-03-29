<template>
  <header class="home-top-bar">
    <h1 class="home-top-bar__title">
      Polimane
    </h1>

    <Button size="md" variant="primary" class="home-top-bar__add">
      Додати Схему
    </Button>
  </header>
</template>

<script setup lang="ts">
import { Button } from '@/components/button';
</script>

<style scoped>
@layer page {
  .home-top-bar {
    position: sticky;
    top: 0;
    display: flex;
    align-items: center;
    padding: 12px 16px;
    background-color: var(--color-white);
    border-bottom: 1px solid color-mix(in srgb, var(--color-primary), transparent 90%);
  }

  .home-top-bar__title {
    font-size: 20px;
    font-weight: 500;
  }

  .home-top-bar__add {
    margin-left: auto;
    margin-right: 2px;
  }
}
</style>
