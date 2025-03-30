<template>
  <header class="home-top-bar">
    <h1 class="home-top-bar__title">
      Polimane
    </h1>

    <div class="home-top-bar__actions">
      <HomeSettingsModal v-slot="{ open }">
        <Button icon size="md" variant="secondary" @click="open">
          <SettingsIcon />
        </Button>
      </HomeSettingsModal>

      <Button size="md" variant="primary">
        Додати Схему
      </Button>
    </div>
  </header>
</template>

<script setup lang="ts">
import { Button } from '@/components/button';
import { SettingsIcon } from '@/components/icon';
import HomeSettingsModal from './HomeSettingsModal.vue';
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
    border-bottom: 1px solid var(--color-divider);
  }

  .home-top-bar__title {
    font-size: 20px;
    font-weight: 500;
  }

  .home-top-bar__actions {
    margin-left: auto;
    margin-right: 2px;
    display: flex;
    gap: 8px;
    align-items: center;
  }
}
</style>
