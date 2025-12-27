<template>
  <header class="common-layout-top-bar">
    <template v-if="isHomeRoute">
      <LogoIcon size="28" />

      <h1 class="common-layout-top-bar__title">
        Polimane
      </h1>
    </template>

    <div class="common-layout-top-bar__back-button-container" v-else>
      <Button
        truncate
        :to="homeRoute"
        :prepend-icon="ArrowBackIcon"
        title="Назад"
        class="common-layout-top-bar__back-button"
      >
        {{ title }}
      </Button>
    </div>

    <div class="common-layout-top-bar__actions">
      <slot />
    </div>
  </header>
</template>

<script setup lang="ts">
import { computed, type Slot } from 'vue';
import { type RouteLocationRaw, useRoute } from 'vue-router';
import { ArrowBackIcon, LogoIcon } from '../icon';
import { Button } from '../button';

defineProps<{
  title: string;
}>();

defineSlots<{
  default: Slot;
}>();

const route = useRoute();

const homeRoute: RouteLocationRaw = { name: 'home' };
const isHomeRoute = computed(() => route.name === homeRoute.name);
</script>

<style scoped>
@layer components {
  .common-layout-top-bar {
    position: sticky;
    top: 0;
    display: flex;
    align-items: center;
    padding: 12px 16px;
    min-height: 60px;
    background-color: var(--color-background-1);
    border-bottom: var(--divider);
    z-index: 10;
  }

  .common-layout-top-bar__title {
    margin-left: 8px;
    font-size: 20px;
    font-weight: 500;
  }

  .common-layout-top-bar__back-button-container {
    margin-right: 32px;
    max-width: 350px;
  }

  .common-layout-top-bar__back-button {
    font-size: 16px;
    font-weight: 450;
    max-width: 100%;
  }

  .common-layout-top-bar__actions {
    margin-left: auto;
    margin-right: 2px;
    display: flex;
    gap: 8px;
    align-items: center;
  }
}
</style>
