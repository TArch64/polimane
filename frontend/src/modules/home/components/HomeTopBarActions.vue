<template>
  <Button
    variant="primary"
    class="home-top-bar__create-schema"
    @click="createSchemaModal.open()"
    v-if="schemasStore.hasSchemas"
  >
    <PlusIcon class="home-top-bar__create-icon" />
    Нова Схема
  </Button>

  <Dropdown>
    <template #activator="{ activatorStyle, open }">
      <Button icon :style="activatorStyle" @click="open">
        <PersonFillIcon />
      </Button>
    </template>

    <DropdownAction
      title="Profile"
      :icon="SettingsIcon"
      :to="settingsProfileRoute"
    />

    <DropdownAction
      title="Log out"
      :icon="LogOutIcon"
      @click="sessionStore.logout"
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
import { useSchemasStore } from '../stores';
import { HomeCreateSchemaModal } from './schemas';

const schemasStore = useSchemasStore();
const sessionStore = useSessionStore();

const settingsProfileRoute: RouteLocationRaw = { name: 'settings-profile' };
const createSchemaModal = useModal(HomeCreateSchemaModal);
</script>

<style scoped>
@layer page {
  .home-top-bar__create-schema {
    margin-right: 12px;
  }

  .home-top-bar__create-icon {
    margin-right: 4px;
  }
}
</style>
