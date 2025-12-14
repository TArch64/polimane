<template>
  <div class="auth">
    <h1>
      Привіт!
    </h1>

    <p class="auth__description">
      Щоб увійти у свій акаунт, натисни кнопку внизу, яка відкриє вікно авторизації
    </p>

    <Button variant="primary" @click="popup.open">
      Увійти
    </Button>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';
import { definePreload } from '@/router/define';
import { Button } from '@/components/button';
import { useAuthStore } from './stores';
import { useAuthPopup } from './composables';

defineOptions({
  beforeRouteEnter: definePreload<'auth'>(async () => {
    const store = useAuthStore();
    await store.load();
  }),
});

const router = useRouter();
const returnTo = (router.currentRoute.value.query['return-to'] as string) ?? '/';

const popup = useAuthPopup({
  onSuccess: () => router.push(returnTo),
  onDeletedUser: () => router.push({ name: 'auth-deleted-user' }),
});
</script>

<style scoped>
@layer page {
  .auth {
    height: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
    flex-direction: column;
    gap: 12px;
  }

  .auth__description {
    max-width: 400px;
    text-align: center;
    text-wrap: balance;
  }
}
</style>
