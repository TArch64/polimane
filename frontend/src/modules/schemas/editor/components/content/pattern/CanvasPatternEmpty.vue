<template />

<script setup lang="ts">
import { useCanvasObject } from '@/modules/schemas/editor/composables';
import type { ISchemaPattern } from '@/models';
import { PatternEmptyObject } from './PatternEmptyObject';

const props = defineProps<{
  pattern: ISchemaPattern;
}>();

const objectId = `${props.pattern.id}-empty`;
useCanvasObject(objectId, () => new PatternEmptyObject());
</script>
