<template>
  <header class="common-layout-top-bar">
    <template v-if="isHomeRoute">
      <LogoIcon size="28" />

      <h1 class="common-layout-top-bar__title">
        Polimane
      </h1>
    </template>

    <div class="common-layout-top-bar__back-button-container" :class="backContainerClasses" v-else>
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

    <CommonLayoutActions v-model:has-actions="hasActions" v-show="hasActions">
      <slot />
    </CommonLayoutActions>
  </header>
</template>

<script setup lang="ts">
import { computed, ref, type Slot } from 'vue';
import { type RouteLocationRaw, useRoute } from 'vue-router';
import { ArrowBackIcon, LogoIcon } from '../icon';
import { Button } from '../button';
import CommonLayoutActions from './CommonLayoutActions.vue';

defineProps<{
  title: string;
}>();

defineSlots<{
  default: Slot;
}>();

const route = useRoute();

const hasActions = ref(false);

const homeRoute: RouteLocationRaw = { name: 'home' };
const isHomeRoute = computed(() => route.name === homeRoute.name);

const backContainerClasses = computed(() => ({
  'common-layout-top-bar__back-button-container--with-actions': hasActions.value,
}));
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

  .common-layout-top-bar__back-button-container--with-actions {
    min-width: 0;
    flex-basis: 0;
    flex-grow: 1;
    margin-right: 32px;
    max-width: 350px;
  }

  .common-layout-top-bar__back-button {
    font-size: 16px;
    font-weight: 450;
    max-width: 100%;
  }
}
</style>
