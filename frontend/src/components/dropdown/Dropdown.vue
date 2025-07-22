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

const offsetX = ref(0);
let closeController: AbortController | null = null;

function getOffsetX(): number {
  const menuRect = menuRef.value!.getBoundingClientRect();
  const offset = window.innerWidth - menuRect.right - 8;
  return offset > 0 ? 0 : offset;
}

function open() {
  if (isOpened.value) {
    return;
  }

  routeTransition.start(async () => {
    isOpened.value = true;
    await nextTick();

    menuRef.value!.showPopover();
    offsetX.value = getOffsetX();
    await nextTick();
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
    offsetX.value = 0;
    await nextTick();
  });
}
</script>

<style scoped>
@layer components {
  .dropdown {
    position-area: bottom center;
    translate: v-bind("offsetX + 'px'");
    margin-top: 4px;
    min-width: 150px;
  }
}
</style>
