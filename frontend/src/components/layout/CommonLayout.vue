<template>
  <FadeTransition>
    <CommonLayoutSelectionBar
      :selected
      :selectedTitle
      data-cursor-selection-ignore
      @clear-selection="$emit('clear-selection')"
      v-if="selected"
    />
  </FadeTransition>

  <CommonLayoutTopBar
    :title
    data-cursor-selection-ignore
    v-visible="!selected"
  >
    <slot name="top-bar-actions" />
  </CommonLayoutTopBar>

  <div class="common-layout__row common-layout__content" v-if="slots.submenu">
    <CommonLayoutSubmenu class="common-layout__submenu">
      <slot name="submenu" />
    </CommonLayoutSubmenu>

    <main class="common-layout__main common-layout__main--aside-menu">
      <slot />
    </main>
  </div>

  <main class="common-layout__main common-layout__content" v-else>
    <slot />
  </main>

  <div
    ref="selectionOverlayRef"
    class="common-layout__selection-overlay"
    data-cursor-selection-include
    v-if="selected"
  />
</template>

<script setup lang="ts">
import { ref, type Slot, toRef } from 'vue';
import { usePageClass } from '@/composables';
import { type MaybeContextMenuAction, useContextMenu } from '@/components/contextMenu';
import { vVisible } from '@/directives';
import { FadeTransition } from '../transition';
import CommonLayoutTopBar from './CommonLayoutTopBar.vue';
import CommonLayoutSelectionBar from './CommonLayoutSelectionBar.vue';
import CommonLayoutSubmenu from './CommonLayoutSubmenu.vue';

const props = withDefaults(defineProps<{
  title?: string;
  selected?: number;
  selectedTitle?: string;
  selectedActions?: MaybeContextMenuAction[];
}>(), {
  title: '',
  selected: 0,
  selectedTitle: '',
  selectedActions: () => [],
});

defineEmits<{
  'clear-selection': [];
}>();

const slots = defineSlots<{
  'default': Slot;
  'submenu'?: Slot;
  'top-bar-actions'?: Slot;
}>();

const selectionOverlayRef = ref<HTMLElement | null>(null);

usePageClass('app--common-layout');

useContextMenu({
  el: selectionOverlayRef,
  control: false,
  title: toRef(props, 'selectedTitle'),
  actions: toRef(props, 'selectedActions'),
});
</script>

<style scoped>
@layer components {
  :global(.app--common-layout) {
    background-color: var(--color-background-2);
  }

  .common-layout__content {
    margin: 0 auto;
    width: 100%;
    max-width: 1024px;
    padding-top: 24px;
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

  .common-layout__main {
    flex-grow: 1;
    display: flex;
    flex-direction: column;
  }

  .common-layout__main--aside-menu {
    padding: 20px;
    min-width: 0;
  }

  .common-layout__selection-overlay {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
  }

  @media (max-width: 768px) {
    .common-layout__row {
      display: block;
      padding: 16px;
    }

    .common-layout__submenu {
      position: static;
      width: 100%;
      margin: 0 0 32px;
    }

    .common-layout__main--aside-menu {
      padding: 0;
    }
  }
}
</style>
