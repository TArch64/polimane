<template>
  <FormCard
    title="Електронна Пошта"
    :has-changes="form.hasChanges"
    @reset="form.reset"
  >
    <TextField
      label
      required
      placeholder="Електронна Пошта"
      v-model="form.data.email"
    />
  </FormCard>
</template>

<script setup lang="ts">
import { FormCard, TextField, useFormData } from '@/components/form';
import { useProfileStore } from '../stores';

const profileStore = useProfileStore();

const form = useFormData({
  email: profileStore.user.email,
});
</script>
