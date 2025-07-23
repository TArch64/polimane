<template>
  <FormCard
    title="Загальна інформація"
    :has-changes="form.hasChanges"
    @reset="form.reset"
  >
    <TextField
      required
      placeholder="Імʼя"
      v-model="form.data.firstName"
    />

    <TextField
      required
      placeholder="Прізвище"
      v-model="form.data.lastName"
    />

    <TextField
      required
      placeholder="Нікнейм"
      v-model="form.data.username"
    />
  </FormCard>
</template>

<script setup lang="ts">
import { FormCard, TextField, useFormData } from '@/components/form';
import { useSessionStore } from '@/stores';

const sessionStore = useSessionStore();

const form = useFormData(sessionStore.user);
</script>
