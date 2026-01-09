<template>
  <HomeListEmpty>
    <template #description>
      <p>Поки що не створено жодної схеми для бісеру</p>
    </template>

    <template #actions>
      <Button
        variant="primary"
        :prepend-icon="PlusIcon"
        :disabled="limit.isReached"
        @click="createModal.open()"
      >
        Нова Схема
      </Button>
    </template>
  </HomeListEmpty>
</template>

<script setup lang="ts">
import { HomeListEmpty, SchemaCreateModal } from '@/modules/home/components';
import { Button } from '@/components/button';
import { useModal } from '@/components/modal';
import { PlusIcon } from '@/components/icon';
import { useSchemasCreatedLimit } from '@/composables/subscription';

const limit = useSchemasCreatedLimit();
const createModal = useModal(SchemaCreateModal);
</script>
