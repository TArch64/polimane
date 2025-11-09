<template>
  <SchemaAccessRow
    :displayName
    :editable="!isCurrentUser"
    :uneven-access="user.isUnevenAccess"
    :access="user.access"
    @update:access="updateAccess"
    @delete="deleteUser"
  />
</template>

<script setup lang="ts">
import { computed } from 'vue';
import type { ISchemaUser } from '@/models';
import { useSessionStore } from '@/stores';
import { useAsyncAction } from '@/composables';
import { AccessLevel } from '@/enums';
import { useSchemaUsersStore } from './schemaUsersStore';
import SchemaAccessRow from './SchemaAccessRow.vue';

const props = defineProps<{
  user: ISchemaUser;
}>();

const sessionStore = useSessionStore();
const usersStore = useSchemaUsersStore();
const isCurrentUser = computed(() => props.user.id === sessionStore.user.id);

const displayName = computed(() => {
  const { firstName, lastName, email } = props.user;

  if (firstName || lastName) {
    return [firstName, lastName].filter(Boolean).join(' ');
  }

  return email;
});

const updateAccess = useAsyncAction(async (access: AccessLevel) => {
  await usersStore.updateUserAccess(props.user, access);
});

const deleteUser = useAsyncAction(async () => {
  await usersStore.deleteUser(props.user);
});
</script>

<style scoped>
@layer page {
  .access-user {
    display: flex;
    align-items: center;
    min-height: 32px;
    gap: 8px;
  }

  .access-user__name {
    flex-basis: 0;
    flex-grow: 1;
    min-width: 0;
    white-space: nowrap;
    text-overflow: ellipsis;
    overflow: hidden;
  }

  .access-user__name-label {
    font-weight: 400;
    color: var(--color-text-3)
  }

  .access-user__actions {
    margin-left: auto;
    display: flex;
    align-items: center;
    gap: 12px;
  }
}
</style>
