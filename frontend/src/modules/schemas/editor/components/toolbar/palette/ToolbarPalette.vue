<template>
  <ColorItem
    :active="isActive"
    :color="store.palette[lastActiveColorId]!"
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
        v-model="activeColor"
        v-if="isOpened"
      />
    </FadeTransition>
  </Teleport>
</template>

<script setup lang="ts">
import { useToolsStore } from '@editor/stores';
import { computed, nextTick, ref } from 'vue';
import { type HotKeyDef, useHotKeys } from '@editor/composables';
import { useDomRef } from '@/composables';
import { FadeTransition } from '@/components/transition';
import ColorItem from './ColorItem.vue';
import ColorPalette from './ColorPalette.vue';

const store = useToolsStore();

const lastActiveColorId = ref(store.activeColorId);
const isActive = computed(() => store.activeColorId !== -1);

const activeColor = computed({
  get: () => lastActiveColorId.value,

  set: (index: number) => {
    store.activateTool(index);
    lastActiveColorId.value = index;
  },
});

const isOpened = ref(false);
const paletteRef = useDomRef<HTMLElement>();

async function open(): Promise<void> {
  if (!isOpened.value) {
    isOpened.value = true;
    store.activateTool(lastActiveColorId.value);
    await nextTick();
    paletteRef.value.showPopover();
  }
}

useHotKeys(
  store.palette.slice(0, 9).map((_, index): HotKeyDef => [
    `Meta_${index + 1}`,
    () => activeColor.value = index,
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
    margin: -16px 0 0 16px;
  }
}
</style>
