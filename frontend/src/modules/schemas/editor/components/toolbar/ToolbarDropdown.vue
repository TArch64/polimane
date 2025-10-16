<template>
  <div class="toolbar-dropdown">
    <slot :open name="activator" />

    <Teleport to="body">
      <FadeTransition>
        <Card
          as="dialog"
          ref="dropdownRef"
          class="toolbar-dropdown__floating"
          v-if="isOpened"
        >
          <slot :close />
        </Card>
      </FadeTransition>
    </Teleport>
  </div>
</template>

<script setup lang="ts">
import { nextTick, ref, type Slot } from 'vue';
import { FadeTransition } from '@/components/transition';
import { onBackdropClick, useDomRef } from '@/composables';
import { Card } from '@/components/card';

defineSlots<{
  default: Slot<{
    close: () => void;
  }>;

  activator: Slot<{
    open: () => Promise<void>;
  }>;
}>();

const dropdownRef = useDomRef<HTMLDialogElement | null>(null);
const isOpened = ref(false);

async function open(): Promise<void> {
  if (!isOpened.value) {
    isOpened.value = true;
    await nextTick();
    dropdownRef.value!.showModal();
  }
}

function close(): void {
  isOpened.value = false;
}

onBackdropClick(dropdownRef, close);
</script>

<style scoped>
@layer page {
  .toolbar-dropdown {
    anchor-name: --toolbar-dropdown;
  }

  .toolbar-dropdown__floating {
    position-anchor: --toolbar-dropdown;
    position-area: right span-y-end;
    margin: -16px 0 0 8px;
    padding: 8px 6px;

    &::backdrop {
      background: transparent;
    }
  }
}
</style>
