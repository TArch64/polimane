<template>
  <Card class="welcome-login">
    <Form class="welcome-login__form" @submit="login.call">
      <TextField
        placeholder="Користувач"
        v-model="form.username"
      />

      <TextField
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
import { Button } from '@/components/button';
import { Form, TextField } from '@/components/form';
import { Card } from '@/components/card';
import { useAsyncAction } from '@/composables';

const form = reactive({
  username: '',
  password: '',
});

const login = useAsyncAction(async () => {
  return new Promise((resolve) => setTimeout(resolve, 1000));
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
