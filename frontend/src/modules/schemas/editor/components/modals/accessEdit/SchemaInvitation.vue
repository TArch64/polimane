<template>
  <li class="schema-invitation">
    <p class="schema-invitation__name">
      {{ invitation.email }}
    </p>

    <div class="schema-invitation__actions">
      <SchemaAccessField
        :model-value="invitation.access"
        @update:model-value="updateAccess"
      />

      <Button
        icon
        danger
        size="md"
        variant="secondary"
        :loading="deleteInvitation.isActive"
        @click="deleteInvitation"
      >
        <TrashIcon />
      </Button>
    </div>
  </li>
</template>

<script setup lang="ts">
import { useSchemaUsersStore } from '@editor/stores';
import type { ISchemaUserInvitation } from '@/models';
import { Button } from '@/components/button';
import { TrashIcon } from '@/components/icon';
import { useAsyncAction } from '@/composables';
import { AccessLevel } from '@/enums';
import SchemaAccessField from './SchemaAccessField.vue';

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

<style scoped>
@layer page {
  .schema-invitation {
    display: flex;
    align-items: center;
    min-height: 32px;
    gap: 8px;
  }

  .schema-invitation__name {
    flex-basis: 0;
    flex-grow: 1;
    min-width: 0;
    white-space: nowrap;
    text-overflow: ellipsis;
    overflow: hidden;
  }

  .schema-invitation__actions {
    margin-left: auto;
    display: flex;
    align-items: center;
    gap: 12px;
  }
}
</style>
