<template>
  <div>
    <ColorItem
      title="Колір Бісеру"
      class="toolbar-palette"
      :color="store.activeColor"
      @click="open"
    />

    <Teleport to="body">
      <FadeTransition>
        <ColorPalette
          ref="paletteRef"
          class="toolbar-palette__floating"
          @close="isOpened = false"
          v-if="isOpened"
        />
      </FadeTransition>
    </Teleport>
  </div>
</template>

<script setup lang="ts">
import { useToolsStore } from '@editor/stores';
import { nextTick, ref } from 'vue';
import { type HotKeyDef, useHotKeys } from '@editor/composables';
import { useDomRef } from '@/composables';
import { FadeTransition } from '@/components/transition';
import ColorItem from './ColorItem.vue';
import ColorPalette from './ColorPalette.vue';

const store = useToolsStore();

const isOpened = ref(false);
const paletteRef = useDomRef<HTMLDialogElement>();

async function open(): Promise<void> {
  if (!isOpened.value) {
    isOpened.value = true;
    await nextTick();
    paletteRef.value.showModal();
  }
}

useHotKeys(
  store.palette.slice(0, 9).map((_, index): HotKeyDef => [
    `Meta_${index + 1}`,
    () => store.activateColor(store.palette[index]!),
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
