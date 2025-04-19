<template>
  <slot
    name="activator"
    :open
    :activatorStyle="{ anchorName: anchorVar }"
  />

  <Teleport to="body">
    <Transition
      name="dropdown-"
      :duration="150"
    >
      <div
        ref="menuRef"
        role="menu"
        popover="manual"
        class="dropdown"
        :style="menuStyles"
        v-if="isOpened"
      >
        <slot />
      </div>
    </Transition>
  </Teleport>
</template>

<script setup lang="ts">
import { computed, nextTick, ref, type Slot } from 'vue';
import { newId, wait } from '@/helpers';

defineSlots<{
  activator: Slot<{
    open: () => void;
    activatorStyle: { anchorName: string };
  }>;

  default: Slot;
}>();

const menuRef = ref<HTMLElement | null>(null);

const anchorVar = `--dropdown-${newId()}`;
const isOpened = ref(false);

const menuStyles = computed(() => ({
  positionAnchor: anchorVar,
}));

async function open(): Promise<void> {
  isOpened.value = true;

  await nextTick();
  menuRef.value!.showPopover();

  await wait(50);
  window.addEventListener('click', () => close(), { once: true });
}

function close(): void {
  isOpened.value = false;
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
  }

  .dropdown--enter-active,
  .dropdown--leave-active {
    transition: opacity 0.15s ease-out;
    will-change: opacity;
  }

  .dropdown--enter-from,
  .dropdown--leave-to {
    opacity: 0;
  }

  .dropdown--enter-to,
  .dropdown--leave-from {
    opacity: 1;
  }
}
</style>
