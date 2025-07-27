<template>
  <ResetPasswordCard />
  <AuthFactorsCard />
</template>

<script setup lang="ts">
import { definePreload } from '@/router/define';
import { AuthFactorsCard, ResetPasswordCard } from './components';
import { useAuthFactorsStore } from './stores';

defineOptions({
  beforeRouteEnter: definePreload<'settings-security'>(async () => {
    const authFactorsStore = useAuthFactorsStore();
    await authFactorsStore.load();
  }),
});
</script>
