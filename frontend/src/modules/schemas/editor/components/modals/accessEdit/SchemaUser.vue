<template>
  <li class="access-user">
    <p class="access-user__name">
      {{ displayName }}

      <span v-if="isCurrentUser" class="access-user__name-label">
        (Ви)
      </span>
    </p>

    <div class="access-user__actions">
      <Button icon danger size="md" variant="secondary">
        <TrashIcon />
      </Button>
    </div>
  </li>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import type { ISchemaUser } from '@/models';
import { useSessionStore } from '@/stores';
import { Button } from '@/components/button';
import { TrashIcon } from '@/components/icon';

const props = defineProps<{
  user: ISchemaUser;
}>();

const sessionStore = useSessionStore();
const isCurrentUser = computed(() => props.user.id === sessionStore.user.id);

const displayName = computed(() => {
  const { firstName, lastName, email } = props.user;

  if (firstName || lastName) {
    return [firstName, lastName].filter(Boolean).join(' ');
  }

  return email;
});
</script>

<style scoped>
@layer page {
  .access-user {
    display: flex;
    align-items: center;
  }

  .access-user__name {
    font-weight: 500;
  }

  .access-user__name-label {
    font-weight: 400;
    color: var(--color-text-3)
  }

  .access-user__actions {
    margin-left: auto;
  }
}
</style>
