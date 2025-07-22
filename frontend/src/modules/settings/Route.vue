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

    <RouterView />
  </CommonLayout>
</template>

<script setup lang="ts">
import type { RouteLocationRaw } from 'vue-router';
import { CommonLayout } from '@/components/layout';
import { Button } from '@/components/button';
import { type IconComponent, PersonFillIcon, UnlockIcon } from '@/components/icon';

interface ISettingsRoute {
  title: string;
  to: RouteLocationRaw;
  icon: IconComponent;
}

const routes: ISettingsRoute[] = [
  {
    title: 'Профіль',
    to: { name: 'settings-profile' },
    icon: PersonFillIcon,
  },
  {
    title: 'Безпека',
    icon: UnlockIcon,
    to: { name: 'settings-security' },
  },
];
</script>
