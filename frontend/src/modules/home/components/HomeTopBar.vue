<template>
  <header class="home-top-bar">
    <LogoIcon size="28" />

    <h1 class="home-top-bar__title">
      Polimane
    </h1>

    <div class="home-top-bar__actions">
      <Dropdown>
        <template #activator="{ activatorStyle, open }">
          <Button icon :style="activatorStyle" @click="open">
            <PersonFillIcon />
          </Button>
        </template>

        <DropdownAction
          title="Log out"
          :icon="LogOutIcon"
          @click="sessionStore.logout"
        />
      </Dropdown>

      <Button
        variant="primary"
        @click="createSchemaModal.open()"
        v-if="schemasStore.hasSchemas"
      >
        <PlusIcon class="home-top-bar__create-icon" />
        Нова Схема
      </Button>
    </div>
  </header>
</template>

<script setup lang="ts">
import { Button } from '@/components/button';
import { useModal } from '@/components/modal';
import { LogoIcon, LogOutIcon, PersonFillIcon, PlusIcon } from '@/components/icon';
import { Dropdown, DropdownAction } from '@/components/dropdown';
import { useSessionStore } from '@/stores';
import { useSchemasStore } from '../stores';
import { HomeCreateSchemaModal } from './schemas';

const schemasStore = useSchemasStore();
const sessionStore = useSessionStore();

const createSchemaModal = useModal(HomeCreateSchemaModal);
</script>

<style scoped>
@layer page {
  .home-top-bar {
    position: sticky;
    top: 0;
    display: flex;
    align-items: center;
    padding: 12px 16px;
    min-height: 60px;
    background-color: var(--color-background-1);
    border-bottom: var(--divider);
  }

  .home-top-bar__title {
    margin-left: 8px;
    font-size: 20px;
    font-weight: 500;
  }

  .home-top-bar__actions {
    margin-left: auto;
    margin-right: 2px;
    display: flex;
    gap: 12px;
    align-items: center;
  }

  .home-top-bar__create-icon {
    margin-right: 4px;
  }
}
</style>
