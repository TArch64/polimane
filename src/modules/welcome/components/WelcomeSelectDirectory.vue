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
import { Button } from '@/components/button';
import { useStorageStore } from '@/stores';
import { useRouter } from 'vue-router';

const router = useRouter();
const storageStore = useStorageStore();

async function selectDirectory(): Promise<void> {
  await storageStore.selectDirectory();
  document.startViewTransition(() => router.push({ name: 'home' }));
}
</script>
