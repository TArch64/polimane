<template>
  <WelcomeLayout>
    <h1 class="welcome-header">
      Polimane
    </h1>

    <p class="welcome-description">
      Твій помічник у створенні схем для бісеру
    </p>

    <WelcomeSelectDirectory />
  </WelcomeLayout>
</template>

<script setup lang="ts">
import { WelcomeLayout, WelcomeSelectDirectory } from './components';
</script>

<style scoped>
@layer page {
  .welcome-header {
    margin-bottom: 4px;
  }

  .welcome-description {
    margin-bottom: 24px;
  }
}
</style>
