<template>
  <div class="auth-factors-empty">
    <Button
      variant="primary"
      :disabled="createAuthFactor.isActive"
      @click="createAuthFactor"
    >
      Увімкнути
    </Button>
  </div>
</template>

<script setup lang="ts">
import { Button } from '@/components/button';
import { useCreateAuthFactor } from '../../composables';

const createAuthFactor = useCreateAuthFactor();
</script>

<style scoped>
@layer page {
  .auth-factors-empty {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    padding: 40px;
  }
}
</style>
