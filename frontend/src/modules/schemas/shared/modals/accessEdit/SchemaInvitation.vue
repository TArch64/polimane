<template>
  <SchemaAccessRow
    :display-name="invitation.email"
    :uneven-access="invitation.isUnevenAccess"
    :access="invitation.access"
    @update:access="updateAccess"
    @delete="deleteInvitation"
  />
</template>

<script setup lang="ts">
import type { ISchemaUserInvitation } from '@/models';
import { useAsyncAction } from '@/composables';
import { AccessLevel } from '@/enums';
import { useSchemaUsersStore } from './schemaUsersStore';
import SchemaAccessRow from './SchemaAccessRow.vue';

const props = defineProps<{
  invitation: ISchemaUserInvitation;
}>();

const usersStore = useSchemaUsersStore();

const updateAccess = useAsyncAction(async (access: AccessLevel) => {
  await usersStore.updateInvitationAccess(props.invitation, access);
});

const deleteInvitation = useAsyncAction(async () => {
  await usersStore.deleteInvitation(props.invitation);
});
</script>
