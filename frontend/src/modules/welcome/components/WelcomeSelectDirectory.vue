<template>
  <Button
    size="lg"
    variant="primary"
    @click="selectDirectory"
  >
    Обрати Директорію
  </Button>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';
import { Button } from '@/components/button';
import { useStorageStore } from '@/stores';

const router = useRouter();
const storageStore = useStorageStore();

async function selectDirectory(): Promise<void> {
  await storageStore.selectDirectory();
  document.startViewTransition(() => router.push({ name: 'home' }));
}
</script>
