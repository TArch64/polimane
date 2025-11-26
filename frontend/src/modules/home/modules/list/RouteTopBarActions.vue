<template>
  <HomeBarRouteActions>
    <Dropdown>
      <template #activator="{ activatorStyle, open }">
        <Button icon title="Профіль" :style="activatorStyle" @click="open">
          <PersonFillIcon />
        </Button>
      </template>

      <DropdownText class="home-top-bar__horizontal-divider">
        {{ displayName }}
      </DropdownText>

      <DropdownAction
        title="Log out"
        :icon="LogOutIcon"
        @click="logout"
      />
    </Dropdown>

    <Button icon title="Налаштування" :to="settingsProfileRoute">
      <SettingsIcon />
    </Button>
  </HomeBarRouteActions>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import type { RouteLocationRaw } from 'vue-router';
import { LogOutIcon, PersonFillIcon, SettingsIcon } from '@/components/icon';
import { Dropdown, DropdownAction, DropdownText } from '@/components/dropdown';
import { Button } from '@/components/button';
import { useSessionStore } from '@/stores';
import { useAsyncAction, useProgressBar } from '@/composables';
import { HomeBarRouteActions } from '@/modules/home/components';

const sessionStore = useSessionStore();

const displayName = computed(() => {
  const { firstName, lastName, email } = sessionStore.user;

  if (firstName || lastName) {
    return [firstName, lastName].filter(Boolean).join(' ');
  }

  return email;
});

const settingsProfileRoute: RouteLocationRaw = { name: 'settings-profile' };

const logout = useAsyncAction(sessionStore.logout);
useProgressBar(logout);
</script>

<style scoped>
@layer page {
  .home-top-bar__horizontal-divider {
    border-bottom: var(--divider);
  }
}
</style>
