<template>
  <SchemaAccessRow
    :displayName
    :editable="!isCurrentUser"
    :access="user.access"
    :even-access="user.isEvenAccess"
    :all-access="user.isAllAccess"
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
