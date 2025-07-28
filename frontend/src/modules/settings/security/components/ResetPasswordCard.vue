<template>
  <Card title="Пароль">
    <p>
      Для того щоб змінити пароль
      <Button variant="inline" @click="resetPassword">
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
import { useAsyncAction } from '@/composables';
import { useProgressBar } from '@/composables/useProgressBar';
import { useSecurityStore } from '../stores';

const router = useRouter();
const securityStore = useSecurityStore();

const resetPassword = useAsyncAction(async (): Promise<void> => {
  await securityStore.resetPassword();
  await router.push({ name: 'auth' });
});

useProgressBar(resetPassword);
</script>
