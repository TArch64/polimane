<template>
  <div class="auth">
    <Button variant="primary" @click="popup.open">
      Log In
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
});
</script>

<style scoped>
@layer page {
  .auth {
    height: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 12px;
  }
}
</style>
