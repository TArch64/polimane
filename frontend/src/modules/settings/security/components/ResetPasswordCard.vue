<template>
  <Card title="Пароль">
    <p>
      Для того щоб змінити пароль

      <Button
        variant="inline"
        :style="deleteConfirm.anchorStyle"
        @click="resetPasswordIntent"
      >
        натисніть тут.
      </Button>

      Ця дія примусово закінче дію всіх активних сесій після чого ви зможете увійти в систему зі
      встановленням нового пароля
    </p>
  </Card>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';
import { Card } from '@/components/card';
import { Button } from '@/components/button';
import { useAsyncAction, useProgressBar } from '@/composables';
import { useConfirm } from '@/components/confirm';
import { useSecurityStore } from '../stores';

const router = useRouter();
const securityStore = useSecurityStore();

const deleteConfirm = useConfirm({
  message: 'Ви впевнені, що хочете скинути пароль?',
  acceptButton: 'Так',
  danger: true,
  control: false,
});

const resetPassword = useAsyncAction(async (): Promise<void> => {
  await securityStore.resetPassword();
  await router.push({ name: 'auth' });
});

async function resetPasswordIntent() {
  if (await deleteConfirm.ask()) {
    await resetPassword();
  }
}

useProgressBar(resetPassword);
</script>
