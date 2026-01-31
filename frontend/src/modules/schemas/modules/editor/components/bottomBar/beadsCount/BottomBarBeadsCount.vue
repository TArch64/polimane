<template>
  <div>
    <Popover
      ref="popoverRef"
      position-area="top center"
    >
      <template #activator="{ open, ref }">
        <BottomBarMetric
          label="Бісер"
          :ref
          :interactive="counter.isReached"
          :class="classes"
          :value="formattedValue"
          :max-value="formattedMaxValue"
          :disabled="!counter.isReached"
          @click="open"
        />
      </template>

      <BeadsLimitReachedPopover />
    </Popover>
  </div>
</template>

<script setup lang="ts">
import { useEditorStore } from '@editor/stores';
import { computed, onMounted, ref, watch } from 'vue';
import type { ComponentExposed } from 'vue-component-type-helpers';
import { useNumberFormatter } from '@/composables';
import { useSchemaBeadsCounter } from '@/composables/subscription';
import { Popover } from '@/components/popover';
import BottomBarMetric from '../BottomBarMetric.vue';
import BeadsLimitReachedPopover from './BeadsLimitReachedPopover.vue';

const popoverRef = ref<ComponentExposed<typeof Popover>>(null!);

const editorStore = useEditorStore();

const counter = useSchemaBeadsCounter(() => editorStore.schema);

const formattedCurrent = useNumberFormatter(() => counter.current);
const formattedMax = useNumberFormatter(() => counter.max);

const formattedValue = computed(() => {
  const current = formattedCurrent.value || '0';
  const max = formattedMax.value;
  return max ? `${current} / ${max}` : current;
});

const formattedMaxValue = computed(() => {
  const max = formattedMax.value;
  return max ? `${max} / ${max}` : '00 000';
});

const fillStatus = computed(() => {
  if (counter.isReached) {
    return 'limit-reached';
  }
  if (counter.willOverlow(300)) {
    return 'limit-reach-soon';
  }
  return 'normal';
});

const classes = computed(() => {
  return `bottom-bar-beads-count--${fillStatus.value}`;
});

watch(fillStatus, (newStatus) => {
  if (newStatus === 'limit-reached') {
    popoverRef.value.open();
  }
});

onMounted(() => {
  if (counter.isReached) {
    popoverRef.value.open();
  }
});
</script>

<style scoped>
@layer page {
  .bottom-bar-beads-count--limit-reach-soon {
    --metric-value-color: var(--color-warning);
  }

  .bottom-bar-beads-count--limit-reached {
    --metric-value-color: var(--color-danger);
  }
}
</style>
