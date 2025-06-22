<template>
  <slot
    name="activator"
    :open
    :activatorStyle="{ anchorName }"
  />

  <Teleport to="body" v-if="isOpened">
    <DropdownMenu class="dropdown" ref="menuRef" :style="menuStyles">
      <slot />
    </DropdownMenu>
  </Teleport>
</template>

<script setup lang="ts">
import { computed, nextTick, ref, type Slot } from 'vue';
import { newId, waitClickComplete } from '@/helpers';
import { useDomRef, useRouteTransition } from '@/composables';
import DropdownMenu from './DropdownMenu.vue';

defineSlots<{
  activator: Slot<{
    open: () => void;
    activatorStyle: { anchorName: string };
  }>;

  default: Slot;
}>();

const routeTransition = useRouteTransition();
const menuRef = useDomRef<HTMLElement | null>();

const anchorName = `--dropdown-${newId()}`;
const isOpened = ref(false);

const menuStyles = computed(() => ({
  positionAnchor: anchorName,
}));

let closeController: AbortController | null = null;

function open() {
  if (isOpened.value) {
    return;
  }

  routeTransition.start(async () => {
    isOpened.value = true;
    await nextTick();
    menuRef.value!.showPopover();
  });

  waitClickComplete().then(() => {
    closeController = new AbortController();
    window.addEventListener('click', close, { signal: closeController.signal });
    window.addEventListener('contextmenu', close, { signal: closeController.signal });
  });
}

function close(): void {
  closeController?.abort();

  routeTransition.start(async () => {
    isOpened.value = false;
    await nextTick();
  });
}
</script>

<style scoped>
@layer components {
  .dropdown {
    position-area: bottom center;
    margin-top: 4px;
  }
}
</style>
