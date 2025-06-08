<template>
  <header class="home-top-bar">
    <h1 class="home-top-bar__title">
      Polimane
    </h1>

    <div class="home-top-bar__actions">
      <Button
        variant="primary"
        @click="createSchemaModal.open()"
        v-if="schemasStore.hasSchemas"
      >
        <PlusIcon class="home-top-bar__create-icon" />
        Нова Схема
      </Button>
    </div>
  </header>
</template>

<script setup lang="ts">
import { Button } from '@/components/button';
import { useModal } from '@/components/modal';
import { PlusIcon } from '@/components/icon';
import { useSchemasStore } from '../stores';
import { HomeCreateSchemaModal } from './schemas';

const schemasStore = useSchemasStore();
const createSchemaModal = useModal(HomeCreateSchemaModal);
</script>

<style scoped>
@layer page {
  .home-top-bar {
    position: sticky;
    top: 0;
    display: flex;
    align-items: center;
    padding: 12px 16px;
    min-height: 60px;
    background-color: var(--color-background-1);
    border-bottom: var(--divider);
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

  .home-top-bar__create-icon {
    margin-right: 4px;
  }
}
</style>
