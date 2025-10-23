<template>
  <Modal :footer="false" title="Редагувати Доступ">
    {{ usersStore.users }}
  </Modal>
</template>

<script setup lang="ts">
import { useSchemaUsersStore } from '@editor/stores';
import { Modal } from '@/components/modal';

const usersStore = useSchemaUsersStore();
</script>

<style scoped>
@layer page {

}
</style>
