<template>
  <Teleport :to="topRef">
    <div
      inert
      ref="indicatorRef"
      popover="manual"
      class="sortable__indicator"
      :class="classes"
    />
  </Teleport>
</template>

<script setup lang="ts">
import { computed, inject, onMounted, ref } from 'vue';
import type { Edge } from '@atlaskit/pragmatic-drag-and-drop-hitbox/types';
import { TOKEN_TOP_EL } from '@/InjectionToken';
import type { DragDirection } from './DragDirection';

const props = defineProps<{
  edge: Edge;
  elRef: HTMLElement;
  gap: number;
  direction: DragDirection;
  sortableAnchorVar: string;
}>();

const topRef = inject(TOKEN_TOP_EL)!;
const indicatorRef = ref<HTMLElement>(null!);

onMounted(() => indicatorRef.value.showPopover());

const classes = computed(() => [
  `sortable__indicator--${props.direction}`,
  `sortable__indicator--${props.edge}`,
]);
</script>

<style scoped>
@layer components {
  .sortable__indicator {
    position-area: v-bind("edge") center;
    position-anchor: v-bind("sortableAnchorVar");
    margin: 0;
    padding: 0;
  }

  .sortable__indicator--vertical {
    border-top: 2px solid var(--color-primary);
    border-bottom: none;
    width: 100%;
  }

  .sortable__indicator--top {
    margin-bottom: -1px;
  }

  .sortable__indicator--bottom {
    margin-top: -1px;
  }

  .sortable__indicator--horizontal {
    border-left: 2px solid var(--color-primary);
    border-right: none;
  }

  .sortable__indicator--left {
    margin-right: -1px;
  }

  .sortable__indicator--right {
    margin-left: -1px;
  }
}
</style>
