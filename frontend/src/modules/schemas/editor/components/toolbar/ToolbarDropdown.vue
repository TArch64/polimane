<template>
  <div class="toolbar-dropdown">
    <slot :open name="activator" />

    <Teleport to="body">
      <FadeTransition>
        <Card
          as="dialog"
          ref="dropdownRef"
          class="toolbar-dropdown__floating"
          :style="dropdownStyles"
          v-if="isOpened"
        >
          <slot :close />
        </Card>
      </FadeTransition>
    </Teleport>
  </div>
</template>

<script setup lang="ts">
import { computed, nextTick, ref, type Slot, useId } from 'vue';
import { FadeTransition } from '@/components/transition';
import { onBackdropClick, useDomRef } from '@/composables';
import { Card } from '@/components/card';

const props = defineProps<{
  offsetTop: number;
}>();

defineSlots<{
  default: Slot<{
    close: () => void;
  }>;

  activator: Slot<{
    open: () => Promise<void>;
  }>;
}>();

const positionAnchor = `--toolbar-dropdown-${useId()}`;
const dropdownRef = useDomRef<HTMLDialogElement | null>();

const dropdownStyles = computed(() => ({
  positionAnchor,
  '--dropdown-offset-top': `${props.offsetTop}px`,
}));

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
    anchor-name: v-bind("positionAnchor");
  }

  .toolbar-dropdown__floating {
    position-area: right span-y-end;
    margin: var(--dropdown-offset-top) 0 0 8px;
    padding: 8px 6px;

    &::backdrop {
      background: transparent;
    }
  }
}
</style>
