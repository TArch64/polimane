<template>
  <FormCard
    title="Загальна інформація"
    :has-changes="form.hasChanges"
    @reset="form.reset"
    @submit="save"
  >
    <div class="row">
      <TextField
        label
        required
        placeholder="Імʼя"
        class="row__column"
        v-model="form.data.firstName"
      />

      <TextField
        label
        required
        placeholder="Прізвище"
        class="row__column"
        v-model="form.data.lastName"
      />
    </div>
  </FormCard>
</template>

<script setup lang="ts">
import { FormCard, TextField, useFormData } from '@/components/form';
import { useAsyncAction } from '@/composables';
import { useProfileStore } from '../stores';

const profileStore = useProfileStore();

const form = useFormData(() => profileStore.user);

const save = useAsyncAction(async () => {
  await profileStore.update(form.data);
  form.reset();
});
</script>

<style scoped>
@layer page {
  .row {
    display: flex;
    gap: 8px;
  }

  .row__column {
    flex: 1;
  }
}
</style>
