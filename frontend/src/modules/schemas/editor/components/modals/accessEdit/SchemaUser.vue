<template>
  <li class="access-user">
    <p class="access-user__name">
      {{ displayName }}

      <span v-if="isCurrentUser" class="access-user__name-label">
        (Ви)
      </span>
    </p>

    <div class="access-user__actions" v-if="!isCurrentUser">
      <SelectField
        variant="control"
        :model-value="user.access"
        @update:model-value="updateAccess"
      >
        <option
          v-for="option of accessOptions"
          :key="option.level"
          :value="option.level"
        >
          {{ option.title }}
        </option>
      </SelectField>

      <Button
        icon
        danger
        size="md"
        variant="secondary"
        :loading="deleteUser.isActive"
        :style="deleteConfirm.anchorStyle"
        @click="deleteUserIntent"
      >
        <TrashIcon />
      </Button>
    </div>
  </li>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { useSchemaUsersStore } from '@editor/stores';
import type { ISchemaUser } from '@/models';
import { useSessionStore } from '@/stores';
import { Button } from '@/components/button';
import { TrashIcon } from '@/components/icon';
import { SelectField } from '@/components/form';
import { useAsyncAction } from '@/composables';
import { useConfirm } from '@/components/confirm';
import { AccessLevel, AccessLeveList, getAccessLevelTitle } from '@/enums';

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

const deleteConfirm = useConfirm({
  danger: true,
  message: 'Ви впевнені, що хочете заборонити доступ до схему цьому користувачу?',
  acceptButton: 'Заборонити',
});

const deleteUser = useAsyncAction(async () => {
  await usersStore.deleteUser(props.user);
});

async function deleteUserIntent(): Promise<void> {
  if (await deleteConfirm.ask()) await deleteUser();
}

const accessOptions = AccessLeveList.map((level) => ({
  level,
  title: getAccessLevelTitle(level),
}));

const updateAccess = useAsyncAction(async (access: AccessLevel) => {
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
    font-weight: 500;
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
