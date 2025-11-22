<template>
  <HomeListEmpty description="Поки що не створено жодної схеми для бісеру">
    <Button variant="primary" @click="createModal.open()">
      Нова Схема
    </Button>
  </HomeListEmpty>
</template>

<script setup lang="ts">
import { HomeListEmpty, SchemaCreateModal } from '@/modules/home/components';
import { Button } from '@/components/button';
import { useModal } from '@/components/modal';

const createModal = useModal(SchemaCreateModal);
</script>
