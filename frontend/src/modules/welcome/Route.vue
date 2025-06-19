<template>
  <WelcomeLayout>
    <h1 class="welcome-header">
      Polimane
    </h1>

    <p class="welcome-description">
      Твій помічник у створенні схем для бісеру
    </p>

    <WelcomeLogin />
  </WelcomeLayout>
</template>

<script setup lang="ts">
import { WelcomeLayout, WelcomeLogin } from './components';
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
