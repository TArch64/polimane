<template>
  <FormCard
    submit-persistent
    :cancelable="false"
    title="Електронна Пошта"
    submit-text="Підтвердити"
    :has-changes="form.hasChanges"
    :loading="verify.isActive"
    @reset="form.reset"
    @submit="verify"
  >
    <TextField
      label
      required
      variant="control"
      ref="codeFieldRef"
      placeholder="Код Верифікації"
      v-model="form.data.code"
    >
      <template #append>
        <Button icon size="sm" :disabled="retryVerification.isActive" @click="retryVerification">
          <RepeatIcon size="16" />
        </Button>
      </template>
    </TextField>

    <p class="profile-settings-verify-email__note">
      Очікуєтьтся верифікація пошти
      <b>{{ profileStore.user.email }}</b>

      <Button
        variant="inline"
        class="profile-settings-verify-email__change"
        @click="changeVerifyingEmail"
      >
        Обрати іншу
      </Button>
    </p>
  </FormCard>
</template>

<script setup lang="ts">
import { nextTick, ref } from 'vue';
import type { ComponentExposed } from 'vue-component-type-helpers';
import { FormCard, TextField, useFormData } from '@/components/form';
import { RepeatIcon } from '@/components/icon';
import { HttpError, HttpErrorReason, useAsyncAction, useRouteTransition } from '@/composables';
import { Button } from '@/components/button';
import { wait } from '@/helpers';
import { useProfileStore } from '../../stores';

const routeTransition = useRouteTransition();
const profileStore = useProfileStore();

const codeFieldRef = ref<ComponentExposed<typeof TextField>>(null!);

const form = useFormData({
  code: '',
});

function changeVerifyingEmail() {
  routeTransition.start(() => {
    profileStore.activateChangeVerifyingEmail();
    return nextTick();
  });
}

const retryVerification = useAsyncAction(async () => {
  await profileStore.retryEmailVerification();
  return wait(30_000);
});

const verify = useAsyncAction(async () => {
  try {
    await profileStore.verifyEmail(form.data.code);
    form.reset();
  } catch (error) {
    if (HttpError.isError(error) && error.reason === HttpErrorReason.CODE_EXPIRED) {
      codeFieldRef.value.setError('Час життя коду верифікації закінчився');
    }
  }
});
</script>

<style>
@layer page {
  .profile-settings-verify-email__note {
    color: var(--color-text-2);
    font-size: 14px;
    padding-left: 2px;
  }

  .profile-settings-verify-email__change {
    margin-left: 4px;
  }
}
</style>
