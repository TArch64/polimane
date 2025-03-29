<template>
  <WelcomeLayout>
    <h1 class="welcome-header">Polimane</h1>
    <p class="welcome-description">Твій помічник у створенні схем для бісеру</p>

    <Button
      size="lg"
      variant="primary"
    >
      Обрати Директорію
    </Button>
  </WelcomeLayout>
</template>

<script setup lang="ts">
import { WelcomeLayout } from './components';
import { Button } from '../../components/button';
</script>

<style scoped>
.welcome-header {
  margin-bottom: 8px;
}

.welcome-description {
  margin-bottom: 16px;
}
</style>
