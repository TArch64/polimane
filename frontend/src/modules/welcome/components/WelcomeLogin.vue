<template>
  <Card class="welcome-login">
    <Form class="welcome-login__form" @submit="login.call">
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
        size="lg"
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
import { useRouter } from 'vue-router';
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

const login = useAsyncAction(async () => {
  await sessionStore.login(form);
  document.startViewTransition(() => router.push({ name: 'home' }));
});
</script>

<style scoped>
@layer page {
  .welcome-login {
    width: 100%;
  }

  .welcome-login__form {
    width: 100%;
    display: flex;
    flex-direction: column;
    gap: 12px;
  }

  .welcome-login__submit {
    margin-top: 8px;
  }
}
</style>
