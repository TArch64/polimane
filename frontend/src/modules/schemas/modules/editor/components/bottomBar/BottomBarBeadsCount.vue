<template>
  <BottomBarMetric
    label="Бісер"
    :class="classes"
    :value="formattedValue"
    :max-value="formattedMaxValue"
  >
    <Transition name="bottom-bar-beads-count__value-" mode="out-in" :duration="100">
      <span :key="fillStatus">
        {{ formattedValue }}
      </span>
    </Transition>
  </BottomBarMetric>
</template>

<script setup lang="ts">
import { useEditorStore } from '@editor/stores';
import { computed } from 'vue';
import { useNumberFormatter } from '@/composables';
import { useSchemaBeadsCounter } from '@/composables/subscription';
import BottomBarMetric from './BottomBarMetric.vue';

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
</script>

<style scoped>
@layer page {
  .bottom-bar-beads-count--limit-reach-soon {
    --metric-value-color: var(--color-warning);
  }

  .bottom-bar-beads-count--limit-reached {
    --metric-value-color: var(--color-danger);
  }

  .bottom-bar-beads-count__value--enter-from,
  .bottom-bar-beads-count__value--leave-to {
    scale: 1.05;
  }

  .bottom-bar-beads-count__value--enter-active,
  .bottom-bar-beads-count__value--leave-active {
    transition: scale 0.1s ease-out;
    will-change: scale;
  }
}
</style>
