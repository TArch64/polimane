<template>
  <Modal :footer="false" title="Редагувати Доступ">
    <SchemaNewUser class="access-edit__new-user" />

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

    <div class="access-edit__invitations" v-if="usersStore.invitations.length">
      <p class="access-edit__description">
        Користувачі, яким надіслані запрошення
      </p>

      <ul class="access-edit__list">
        <SchemaInvitation
          v-for="invitation of usersStore.invitations"
          :key="invitation.email"
          :invitation
        />
      </ul>
    </div>

    <div class="access-edit__footer" />
  </Modal>
</template>

<script setup lang="ts">
import { Modal } from '@/components/modal';
import { useSchemaUsersStore } from './schemaUsersStore';
import SchemaNewUser from './SchemaNewUser.vue';
import SchemaUser from './SchemaUser.vue';
import SchemaInvitation from './SchemaInvitation.vue';

const usersStore = useSchemaUsersStore();
</script>

<style scoped>
@layer page {
  .access-edit__new-user {
    margin-bottom: 16px;
  }

  .access-edit__description {
    margin-bottom: 16px;
    color: var(--color-text-3);
    font-size: var(--font-sm);
  }

  .access-edit__list {
    display: flex;
    flex-direction: column;
    gap: 8px;
    padding: 0;
    margin: 0;
    list-style-type: none;
  }

  .access-edit__invitations {
    margin-top: 16px;
  }

  .access-edit__footer {
    padding-bottom: 8px;
  }
}
</style>
