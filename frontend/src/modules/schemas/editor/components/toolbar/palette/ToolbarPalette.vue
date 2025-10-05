<template>
  <ColorItem
    active
    :color="store.activeColor"
    class="toolbar-palette"
    @click="open"
  />

  <Teleport to="body">
    <FadeTransition>
      <ColorPalette
        popover="auto"
        ref="paletteRef"
        class="toolbar-palette__floating"
        @toggle="isOpened = $event.newState === 'open'"
        v-if="isOpened"
      />
    </FadeTransition>
  </Teleport>
</template>

<script setup lang="ts">
import { usePaletteStore } from '@editor/stores';
import { useEventListener } from '@vueuse/core';
import { nextTick, ref } from 'vue';
import { useDomRef } from '@/composables';
import { FadeTransition } from '@/components/transition';
import ColorItem from './ColorItem.vue';
import ColorPalette from './ColorPalette.vue';

const store = usePaletteStore();

const isOpened = ref(false);
const paletteRef = useDomRef<HTMLElement>();

async function open(): Promise<void> {
  if (!isOpened.value) {
    isOpened.value = true;
    await nextTick();
    paletteRef.value.showPopover();
  }
}

useEventListener('keydown', (event) => {
  if (!event.metaKey) return;
  if (!event.code.startsWith('Digit')) return;

  event.preventDefault();
  const index = Number(event.code.replace('Digit', ''));
  store.activateTool(index === 0 ? 'eraser' : index - 1);
});
</script>

<style scoped>
@layer page {
  .toolbar-palette {
    width: 40px;
    height: 23px;
    border-radius: 4px;
    anchor-name: --toolbar-palette;
  }

  .toolbar-palette__floating {
    position-anchor: --toolbar-palette;
    position-area: right span-y-end;
    margin: -16px 0 0 16px;
  }
}
</style>
