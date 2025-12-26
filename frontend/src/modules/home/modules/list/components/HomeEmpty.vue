<template>
  <HomeListEmpty>
    <template #description>
      <p>Поки що не створено жодної схеми для бісеру</p>
    </template>

    <template #actions>
      <Button variant="primary" @click="createModal.open()">
        Нова Схема
      </Button>
    </template>
  </HomeListEmpty>
</template>

<script setup lang="ts">
import { HomeListEmpty, SchemaCreateModal } from '@/modules/home/components';
import { Button } from '@/components/button';
import { useModal } from '@/components/modal';

const createModal = useModal(SchemaCreateModal);
</script>
