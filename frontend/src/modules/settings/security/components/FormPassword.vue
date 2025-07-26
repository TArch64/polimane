<template>
  <FormCard
    title="Пароль"
    :has-changes="formData.hasChanges"
  >
    <div class="row">
      <PasswordField
        label
        required
        class="row__column"
        variant="control"
        placeholder="Новий Пароль"
        v-model="formData.data.password"
      />

      <PasswordField
        label
        required
        class="row__column"
        variant="control"
        placeholder="Підтвердження"
        v-model="formData.data.passwordConfirm"
        v-model:custom-error="confirmFieldError"
      />
    </div>
  </FormCard>
</template>

<script setup lang="ts">
import { ref, toRef, watch } from 'vue';
import { FormCard, PasswordField, useFormData } from '@/components/form';

const formData = useFormData({
  password: '',
  passwordConfirm: '',
});

const confirmFieldError = ref('');

watch([
  toRef(formData.data, 'password'),
  toRef(formData.data, 'passwordConfirm'),
], ([password, passwordConfirm]) => {
  confirmFieldError.value = password === passwordConfirm ? '' : 'Паролі не співпадають';
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
