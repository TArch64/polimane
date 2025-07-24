<template>
  <CommonLayoutTopBar :title>
    <slot name="top-bar-actions" />
  </CommonLayoutTopBar>

  <div class="common-layout__row common_layout__content" v-if="slots.submenu">
    <CommonLayoutSubmenu class="common-layout__submenu">
      <slot name="submenu" />
    </CommonLayoutSubmenu>

    <main class="common-layout__main common-layout__main--aside-menu">
      <slot />
    </main>
  </div>

  <main class="common-layout__main common_layout__content" v-else>
    <slot />
  </main>
</template>

<script setup lang="ts">
import type { Slot } from 'vue';
import { usePageClass } from '@/composables';
import CommonLayoutTopBar from './CommonLayoutTopBar.vue';
import CommonLayoutSubmenu from './CommonLayoutSubmenu.vue';

defineProps<{
  title?: string;
}>();

const slots = defineSlots<{
  'default': Slot;
  'submenu'?: Slot;
  'top-bar-actions'?: Slot;
}>();

usePageClass('app--common-layout');
</script>

<style scoped>
@layer components {
  :global(.app--common-layout) {
    background-color: var(--color-background-2);
  }

  .common_layout__content {
    margin: 0 auto;
    width: 100%;
    max-width: 1024px;
  }

  .common-layout__row {
    display: flex;
    align-items: flex-start;
  }

  .common-layout__submenu {
    flex-shrink: 0;
    width: 250px;
    margin: 20px;
    position: sticky;
    top: 20px;
  }

  .common-layout__main--aside-menu {
    flex-grow: 1;
    padding: 20px;
  }
}
</style>
