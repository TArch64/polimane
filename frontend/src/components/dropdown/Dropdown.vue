<template>
  <slot
    name="activator"
    :open
    :isOpened
    :activatorStyle="{ anchorName }"
  />

  <Teleport to="body" :disabled="!isOpened && !transitionState.isActive">
    <FadeTransition :state="transitionState">
      <DropdownMenu
        ref="menuRef"
        class="dropdown-menu"
        :style="menuStyles"
        @click="close"
        v-popover-shift
        v-if="isOpened"
      >
        <slot />
      </DropdownMenu>
    </FadeTransition>
  </Teleport>
</template>

<script setup lang="ts">
import { computed, nextTick, ref, type Slot } from 'vue';
import { newId, waitClickComplete } from '@/helpers';
import { useDomRef, useTransitionState } from '@/composables';
import { vPopoverShift } from '@/directives';
import { FadeTransition } from '@/components/transition';
import DropdownMenu from './DropdownMenu.vue';

defineSlots<{
  activator: Slot<{
    open: () => void;
    isOpened: boolean;
    activatorStyle: { anchorName: string };
  }>;

  default: Slot;
}>();

const transitionState = useTransitionState();
const menuRef = useDomRef<HTMLElement | null>();

const anchorName = `--dropdown-${newId()}`;
const isOpened = ref(false);

const menuStyles = computed(() => ({
  positionAnchor: anchorName,
}));

let closeController: AbortController | null = null;

function close(): void {
  closeController?.abort();
  isOpened.value = false;
}

function closeEvent(event: Event): void {
  if (menuRef.value?.contains(event.target as Node)) {
    return;
  }

  close();
}

async function open() {
  if (isOpened.value) {
    return;
  }

  isOpened.value = true;
  await nextTick();

  menuRef.value!.showPopover();
  await waitClickComplete();

  closeController = new AbortController();

  window.addEventListener('mousedown', closeEvent, {
    signal: closeController.signal,
    capture: true,
  });

  window.addEventListener('contextmenu', closeEvent, {
    signal: closeController.signal,
    capture: true,
  });
}
</script>

<style scoped>
@layer components {
  .dropdown-menu {
    position-area: bottom center;
    margin-top: 4px;
    min-width: 150px;
  }
}
</style>
