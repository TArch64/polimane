<template>
  <ColorItem
    v-bind="attrs"
    :color="store.activeColor"
    class="toolbar-palette"
    @click="open"
  />

  <Teleport to="body">
    <FadeTransition>
      <ColorPalette
        popover="manual"
        ref="paletteRef"
        class="toolbar-palette__floating"
        @close="isOpened = false"
        v-if="isOpened"
      />
    </FadeTransition>
  </Teleport>
</template>

<script setup lang="ts">
import { useToolsStore } from '@editor/stores';
import { nextTick, ref, useAttrs } from 'vue';
import { type HotKeyDef, useHotKeys } from '@editor/composables';
import { useDomRef } from '@/composables';
import { FadeTransition } from '@/components/transition';
import ColorItem from './ColorItem.vue';
import ColorPalette from './ColorPalette.vue';

const attrs = useAttrs();
const store = useToolsStore();

const isOpened = ref(false);
const paletteRef = useDomRef<HTMLElement>();

async function open(): Promise<void> {
  if (!isOpened.value) {
    isOpened.value = true;
    await nextTick();
    paletteRef.value.showPopover();
  }
}

useHotKeys(
  store.palette.slice(0, 9).map((color, index): HotKeyDef => [
    `Meta_${index + 1}`,
    () => store.activateColor(color),
  ]),
);
</script>

<style scoped>
@layer page {
  .toolbar-palette {
    anchor-name: --toolbar-palette;
  }

  .toolbar-palette__floating {
    position-anchor: --toolbar-palette;
    position-area: right span-y-end;
    margin: -16px 0 0 8px;
  }
}
</style>
