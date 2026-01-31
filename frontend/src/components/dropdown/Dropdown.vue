<template>
  <slot
    name="activator"
    :open="state.open"
    :isOpened="state.isOpened"
    :activatorStyle="{ anchorName }"
  />

  <Teleport to="body" v-if="state.isOpened || state.transition.isActive">
    <FadeTransition :state="state.transition">
      <DropdownMenu
        ref="menuRef"
        class="dropdown-menu"
        :style="menuStyles"
        @click="state.close"
        v-popover-shift
        v-if="state.isOpened"
      >
        <slot />
      </DropdownMenu>
    </FadeTransition>
  </Teleport>
</template>

<script setup lang="ts">
import { type Slot, useId } from 'vue';
import { useDomRef, usePopoverState } from '@/composables';
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

const menuRef = useDomRef<HTMLElement | null>();
const state = usePopoverState(menuRef);

const anchorName = `--dropdown-${useId()}`;
const menuStyles = { positionAnchor: anchorName };
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
