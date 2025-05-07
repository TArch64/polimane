<template>
  <slot
    name="activator"
    :open
    :activatorStyle="{ anchorName: anchorVar }"
  />

  <Teleport to="body" v-if="isOpened">
    <div
      ref="menuRef"
      role="menu"
      popover="manual"
      class="dropdown"
      :style="menuStyles"
    >
      <slot />
    </div>
  </Teleport>
</template>

<script setup lang="ts">
import { computed, nextTick, ref, type Slot } from 'vue';
import { newId, waitClickComplete } from '@/helpers';
import { useRouteTransition } from '@/composables';

defineSlots<{
  activator: Slot<{
    open: () => void;
    activatorStyle: { anchorName: string };
  }>;

  default: Slot;
}>();

const routeTransition = useRouteTransition();
const menuRef = ref<HTMLElement | null>(null);

const anchorVar = `--dropdown-${newId()}`;
const isOpened = ref(false);

const menuStyles = computed(() => ({
  positionAnchor: anchorVar,
}));

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
    window.addEventListener('click', close, { once: true, capture: true });
  });
}

function close(): void {
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
    background-color: var(--color-background-2);
    border: var(--divider);
    border-radius: var(--rounded-md);
    box-shadow: var(--box-shadow);
    display: flex;
    flex-direction: column;
    gap: 4px;
    width: max-content;
    view-transition-name: dropdown;
  }
}
</style>
