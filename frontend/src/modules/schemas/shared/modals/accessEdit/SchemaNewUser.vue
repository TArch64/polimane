<template>
  <TextField
    type="email"
    ref="emailRef"
    variant="control"
    placeholder="Додати нового користувача (пошта)"
    :disabled="addUser.isActive"
    @keyup.enter.prevent="addUser"
    v-model="email"
  >
    <template #append>
      <Button
        icon
        size="md"
        variant="secondary"
        @click="addUser"
      >
        <PlusIcon />
      </Button>
    </template>
  </TextField>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import type { ComponentExposed } from 'vue-component-type-helpers';
import { TextField } from '@/components/form';
import { Button } from '@/components/button';
import { PlusIcon } from '@/components/icon';
import { HttpError, HttpErrorReason, useAsyncAction } from '@/composables';
import { useSchemaUsersStore } from './schemaUsersStore';

const usersStore = useSchemaUsersStore();

const email = ref('');
const emailRef = ref<ComponentExposed<typeof TextField>>(null!);

const addUser = useAsyncAction(async () => {
  try {
    const value = email.value.trim();

    if (value) {
      email.value = '';
      await usersStore.addUser(value);
    }
  } catch (error) {
    if (HttpError.isReason(error, HttpErrorReason.LIMIT_REACHED_SCHEMAS_CREATED)) {
      emailRef.value.setError('Користувач якого ви намагаєтесь додати не може мати доступ до більшої кількості схем');
    }
  }
});
</script>
