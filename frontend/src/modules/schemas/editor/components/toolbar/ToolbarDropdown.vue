<template>
  <div ref="containerRef" class="toolbar-dropdown">
    <slot :open name="activator" />

    <Teleport to="body" :disabled="!isOpened && !transitionState.isActive">
      <FadeTransition appear :state="transitionState">
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
import { onBackdropClick, useDomRef, useTransitionState } from '@/composables';
import { FadeTransition } from '@/components/transition';
import { Card } from '@/components/card';
import { useToolbarRef } from './toolbarRef';

defineSlots<{
  default: Slot<{ close: () => void }>;
  activator: Slot<{ open: () => Promise<void> }>;
}>();

const positionAnchor = `--toolbar-dropdown-${useId()}`;

const toolbarRef = useToolbarRef();
const containerRef = useDomRef<HTMLElement>();
const dropdownRef = useDomRef<HTMLDialogElement | null>();

const isOpened = ref(false);
const transitionState = useTransitionState();

const offsetTop = ref(0);
const offsetLeft = ref(0);

async function open(): Promise<void> {
  if (isOpened.value) return;

  isOpened.value = true;

  await nextTick();
  dropdownRef.value!.showModal();

  const toolbarRect = toolbarRef.value.getBoundingClientRect();
  const containerRect = containerRef.value.getBoundingClientRect();
  const dropdownRect = dropdownRef.value!.getBoundingClientRect();
  offsetTop.value = -Math.min(dropdownRect.height / 10, 16);
  offsetLeft.value = (toolbarRect.right - containerRect.right) + 4;
}

function close(): void {
  isOpened.value = false;
}

onBackdropClick(close);

const dropdownStyles = computed(() => ({
  positionAnchor,
  '--toolbar-dropdown-offset-top': `${offsetTop.value}px`,
  '--toolbar-dropdown-offset-left': `${offsetLeft.value}px`,
}));

defineExpose({
  open,
  close,
});
</script>

<style scoped>
@property --toolbar-dropdown-offset-top {
  syntax: "<length>";
  inherits: false;
  initial-value: 0;
}

@property --toolbar-dropdown-offset-left {
  syntax: "<length>";
  inherits: false;
  initial-value: 0;
}

@layer page {
  .toolbar-dropdown {
    anchor-name: v-bind("positionAnchor");
  }

  .toolbar-dropdown__floating {
    position-area: right span-y-end;
    margin: var(--toolbar-dropdown-offset-top) 0 0 var(--toolbar-dropdown-offset-left);
    padding: 8px 6px;

    &::backdrop {
      background: transparent;
    }
  }
}
</style>
