<template>
  <BottomBarMetric
    label="Бісер"
    :value="formattedValue"
    :max-value="formattedMaxValue"
  />
</template>

<script setup lang="ts">
import { useEditorStore } from '@editor/stores';
import { computed } from 'vue';
import { useNumberFormatter } from '@/composables';
import { useSchemaBeadsLimit } from '@/composables/subscription';
import BottomBarMetric from './BottomBarMetric.vue';

const editorStore = useEditorStore();

const limit = useSchemaBeadsLimit(() => editorStore.schema);

const formattedCurrent = useNumberFormatter(() => limit.current);
const formattedMax = useNumberFormatter(() => limit.max);

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
</script>
