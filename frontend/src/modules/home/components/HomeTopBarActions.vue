<template>
  <Button
    variant="primary"
    class="home-top-bar__create-schema"
    :prepend-icon="PlusIcon"
    @click="createSchemaModal.open()"
    v-if="schemasStore.hasSchemas"
  >
    Нова Схема
  </Button>

  <Dropdown>
    <template #activator="{ activatorStyle, open }">
      <Button icon :style="activatorStyle" @click="open">
        <PersonFillIcon />
      </Button>
    </template>

    <DropdownAction
      title="Log out"
      :icon="LogOutIcon"
      @click="logout"
    />
  </Dropdown>

  <Button icon :to="settingsProfileRoute">
    <SettingsIcon />
  </Button>
</template>

<script setup lang="ts">
import type { RouteLocationRaw } from 'vue-router';
import { Button } from '@/components/button';
import { useModal } from '@/components/modal';
import { LogOutIcon, PersonFillIcon, PlusIcon, SettingsIcon } from '@/components/icon';
import { Dropdown, DropdownAction } from '@/components/dropdown';
import { useSessionStore } from '@/stores';
import { useAsyncAction, useProgressBar } from '@/composables';
import { useSchemasStore } from '../stores';
import { HomeCreateSchemaModal } from './schemas';

const schemasStore = useSchemasStore();
const sessionStore = useSessionStore();

const settingsProfileRoute: RouteLocationRaw = { name: 'settings-profile' };
const createSchemaModal = useModal(HomeCreateSchemaModal);

const logout = useAsyncAction(sessionStore.logout);
useProgressBar(logout);
</script>

<style scoped>
@layer page {
  .home-top-bar__create-schema {
    margin-right: 12px;
  }
}
</style>
