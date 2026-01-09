<template>
  <Button
    mobile-icon-only
    variant="primary"
    class="home-bar__create-schema"
    :prepend-icon="PlusIcon"
    :disabled="limit.isReached"
    @click="createModal.open()"
  >
    Нова Схема
  </Button>
</template>

<script setup lang="ts">
import { PlusIcon } from '@/components/icon';
import { Button } from '@/components/button';
import { useModal } from '@/components/modal';
import { SchemaCreateModal } from '@/modules/home/components/modals';
import { useSchemasCreatedLimit } from '@/composables/subscription';

const limit = useSchemasCreatedLimit();
const createModal = useModal(SchemaCreateModal);
</script>

<style scoped>
@layer page {
  .home-bar__create-schema {
    view-transition-name: create-schema-button;
  }

  @media (min-width: 768px) {
    .home-bar__create-schema {
      margin-right: 12px;
    }
  }
}
</style>
