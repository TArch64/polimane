<template>
  <ResetPasswordCard />
</template>

<script setup lang="ts">
import { definePreload } from '@/router/define';
import { ResetPasswordCard } from './components';
import { useAuthFactorsStore } from './stores';

defineOptions({
  beforeRouteEnter: definePreload<'settings-security'>(async () => {
    const authFactorsStore = useAuthFactorsStore();
    await authFactorsStore.load();
  }),
});
</script>
