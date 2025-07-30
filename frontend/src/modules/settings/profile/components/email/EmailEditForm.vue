<template>
  <FormCard
    title="Електронна Пошта"
    :has-changes="form.hasChanges"
    :loading="save.isActive"
    @reset="form.reset"
    @submit="save"
  >
    <TextField
      label
      required
      type="email"
      variant="control"
      placeholder="Електронна Пошта"
      v-model="form.data.email"
    />
  </FormCard>
</template>

<script setup lang="ts">
import { FormCard, TextField, useFormData } from '@/components/form';
import { useAsyncAction } from '@/composables';
import { useProfileStore } from '../../stores';

const profileStore = useProfileStore();

const form = useFormData(() => ({
  email: profileStore.user.email,
}));

const save = useAsyncAction(async () => {
  await profileStore.update(form.data);
  form.reset();
});
</script>
