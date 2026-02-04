<template>
  <Button
    mobile-icon-only
    variant="primary"
    class="home-bar__create-schema"
    :prepend-icon="PlusIcon"
    @click="createSchema({ overflowCount: 1 })"
  >
    Нова Схема
  </Button>
</template>

<script setup lang="ts">
import { PlusIcon } from '@/components/icon';
import { Button } from '@/components/button';
import { useModal } from '@/components/modal';
import { SchemaCreateModal, SchemasLimitReachedModal } from '@/modules/home/components/modals';
import { useLimitedAction, useSchemasCreatedCounter } from '@/composables/subscription';

const createModal = useModal(SchemaCreateModal);

const createSchema = useLimitedAction({
  counter: useSchemasCreatedCounter(),
  modal: useModal(SchemasLimitReachedModal),
  onAction: () => void createModal.open(),
});
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
