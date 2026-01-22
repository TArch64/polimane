<template>
  <HomeListEmpty>
    <template #description>
      <p>Поки що не створено жодної схеми для бісеру</p>
    </template>

    <template #actions>
        <Button
            variant="primary"
            :prepend-icon="PlusIcon"
            :disabled="schemasCreatedCounter.isReached"
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
import { useSchemasCreatedCounter } from '@/composables/subscription';

const schemasCreatedCounter = useSchemasCreatedCounter();
const createModal = useModal(SchemaCreateModal);
</script>
