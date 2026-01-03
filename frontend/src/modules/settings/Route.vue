<template>
  <CommonLayout title="Налаштування">
    <template #submenu>
      <Button
        v-for="route of routes"
        :key="route.title"
        :to="route.to"
        :prepend-icon="route.icon"
      >
        {{ route.title }}
      </Button>
    </template>

    <div class="settings">
      <RouterView />
    </div>
  </CommonLayout>
</template>

<script setup lang="ts">
import type { RouteLocationRaw } from 'vue-router';
import { CommonLayout } from '@/components/layout';
import { Button } from '@/components/button';
import { CreditCartIcon, type IconComponent, PersonIcon, UnlockIcon } from '@/components/icon';

interface ISettingsRoute {
  title: string;
  to: RouteLocationRaw;
  icon: IconComponent;
}

const routes: ISettingsRoute[] = [
  {
    title: 'Профіль',
    to: { name: 'settings-profile' },
    icon: PersonIcon,
  },
  {
    title: 'Безпека',
    icon: UnlockIcon,
    to: { name: 'settings-security' },
  },
  {
    title: 'Підписка',
    icon: CreditCartIcon,
    to: { name: 'settings-subscription' },
  },
];
</script>

<style scoped>
@layer page {
  .settings {
    display: flex;
    flex-direction: column;
    gap: 16px;
  }
}
</style>
