<template />

<script setup lang="ts">
import type { ISchemaPattern } from '@/models';
import { useCanvasObject } from '@/modules/schemas/editor/composables';
import { PatternObject } from './PatternObject';

const props = defineProps<{
  pattern: ISchemaPattern;
}>();

useCanvasObject(props.pattern, () => new PatternObject(props.pattern));
</script>
