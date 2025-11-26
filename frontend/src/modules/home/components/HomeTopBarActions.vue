<template>
  <Button
    mobile-icon-only
    variant="primary"
    class="home-top-bar__create-schema"
    :prepend-icon="PlusIcon"
    @click="createModal.open()"
  >
    Нова Схема
  </Button>

  <div
    id="homeTopBarRouteActions"
    class="home-top-bar__route-actions"
  />
</template>

<script setup lang="ts">
import { Button } from '@/components/button';
import { useModal } from '@/components/modal';
import { PlusIcon } from '@/components/icon';
import { SchemaCreateModal } from './modals';

const createModal = useModal(SchemaCreateModal);
</script>

<style scoped>
@layer page {
  .home-top-bar__create-schema + .home-top-bar__route-actions:has(:first-child)::before {
    content: "";
    display: block;
  }

  .home-top-bar__route-actions {
    display: contents;
  }

  @media (max-width: 768px) {
    .home-top-bar__route-actions::before {
      display: none !important;
    }
  }
}
</style>
