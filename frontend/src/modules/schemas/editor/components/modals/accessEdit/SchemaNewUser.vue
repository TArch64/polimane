<template>
  <TextField
    type="email"
    variant="control"
    placeholder="Додати нового користувача (пошта)"
    :disabled="addUser.isActive"
    v-model="email"
  >
    <template #append>
      <Button
        icon
        size="md"
        variant="secondary"
        @click.prevent.stop="addUser"
      >
        <PlusIcon />
      </Button>
    </template>
  </TextField>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useSchemaUsersStore } from '@editor/stores';
import { TextField } from '@/components/form';
import { Button } from '@/components/button';
import { PlusIcon } from '@/components/icon';
import { useAsyncAction } from '@/composables';

const usersStore = useSchemaUsersStore();
const email = ref('');

const addUser = useAsyncAction(async () => {
  const value = email.value.trim();

  if (value) {
    email.value = '';
    await usersStore.addUser(value);
  }
});
</script>
