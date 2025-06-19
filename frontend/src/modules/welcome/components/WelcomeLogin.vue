<template>
  <Card variant="control" class="welcome-login">
    <Form class="welcome-login__form" @submit="login">
      <TextField
        required
        placeholder="Користувач"
        v-model="form.username"
      />

      <TextField
        required
        type="password"
        placeholder="Пароль"
        v-model="form.password"
      />

      <Button
        type="submit"
        variant="primary"
        class="welcome-login__submit"
        :disabled="login.isActive"
      >
        Увійти
      </Button>
    </Form>
  </Card>
</template>

<script setup lang="ts">
import { reactive } from 'vue';
import { type RouteLocationRaw, useRouter } from 'vue-router';
import { Button } from '@/components/button';
import { Form, TextField } from '@/components/form';
import { Card } from '@/components/card';
import { useAsyncAction } from '@/composables';
import { type ILoginInput, useSessionStore } from '@/stores';

const router = useRouter();
const sessionStore = useSessionStore();

const form = reactive<ILoginInput>({
  username: '',
  password: '',
});

function getSuccessRedirect(): RouteLocationRaw {
  return router.currentRoute.value.query['return-to'] as string ?? { name: 'home' };
}

const login = useAsyncAction(async () => {
  await sessionStore.login({
    username: form.username.trim(),
    password: form.password.trim(),
  });

  document.startViewTransition(() => router.push(getSuccessRedirect()));
});
</script>

<style scoped>
@layer page {
  .welcome-login {
    width: 100%;
    padding-left: 8px;
    padding-right: 8px;
  }

  .welcome-login__form {
    width: 100%;
    display: flex;
    flex-direction: column;
    gap: 8px;
  }

  .welcome-login__submit {
    margin-top: 8px;
  }
}
</style>
