<template>
  <BottomBarMetric
    label="Бісер"
    :class="classes"
    :value="formattedValue"
    :max-value="formattedMaxValue"
  />
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
  return formattedMax.value
    ? `${formattedCurrent.value} / ${formattedMax.value}`
    : formattedCurrent.value;
});

const formattedMaxValue = computed(() => {
  return formattedMax.value
    ? `${formattedMax.value} / ${formattedMax.value}`
    : '00 000';
});

const classes = computed(() => {
  if (counter.isReached) {
    return 'bottom-bar-beads-count--limit-reached';
  }
  if (counter.willOverlow(300)) {
    return 'bottom-bar-beads-count--limit-reach-soon';
  }
  return null;
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
