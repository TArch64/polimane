<template>
  <CanvasPatternEmpty :pattern />
</template>

<script setup lang="ts">
import type { ISchemaPattern } from '@/models';
import { provideObjectParent, useCanvasEntityObject } from '@/modules/schemas/editor/composables';
import { PatternObject } from './PatternObject';
import CanvasPatternEmpty from './CanvasPatternEmpty.vue';

const props = defineProps<{
  pattern: ISchemaPattern;
}>();

const object = useCanvasEntityObject(props.pattern, () => new PatternObject(props.pattern));

provideObjectParent(object.container);
</script>
