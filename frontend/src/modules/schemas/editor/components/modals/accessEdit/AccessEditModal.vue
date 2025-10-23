<template>
  <Modal :footer="false" title="Редагувати Доступ">
    <p class="access-edit__description">
      Користувачі, які мають доступ до цієї схеми
    </p>

    <ul class="access-edit__list">
      <SchemaUser
        v-for="user of usersStore.users"
        :key="user.id"
        :user
      />
    </ul>
  </Modal>
</template>

<script setup lang="ts">
import { useSchemaUsersStore } from '@editor/stores';
import { Modal } from '@/components/modal';
import SchemaUser from './SchemaUser.vue';

const usersStore = useSchemaUsersStore();
</script>

<style scoped>
@layer page {
  .access-edit__description {
    margin-bottom: 16px;
    color: var(--color-text-2);
  }

  .access-edit__list {
    display: flex;
    flex-direction: column;
    gap: 8px;
    padding: 0;
    padding-bottom: 8px;
    margin: 0;
    list-style-type: none;
  }
}
</style>
