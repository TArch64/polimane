<template>
  <HomeListEmpty>
    <template #description>
      <p>Поки що не створено жодної схеми для бісеру</p>
    </template>

    <template #actions>
        <Button
            variant="primary"
            :prepend-icon="PlusIcon"
            @click="createSchema({ overflowCount: 1 })"
        >
        Нова Схема
      </Button>
    </template>
  </HomeListEmpty>
</template>

<script setup lang="ts">
import {
  HomeListEmpty,
  SchemaCreateModal,
  SchemasLimitReachedModal,
} from '@/modules/home/components';
import { Button } from '@/components/button';
import { useModal } from '@/components/modal';
import { PlusIcon } from '@/components/icon';
import { useLimitedAction, useSchemasCreatedCounter } from '@/composables/subscription';

const createModal = useModal(SchemaCreateModal);

const createSchema = useLimitedAction({
  counter: useSchemasCreatedCounter(),
  modal: useModal(SchemasLimitReachedModal),
  onAction: () => void createModal.open(),
});
</script>
