<template />

<script setup lang="ts">
import { onObjectEvent, useCanvasObject } from '@/modules/schemas/editor/composables';
import type { ISchemaPattern } from '@/models';
import { PatternEmptyObject } from './PatternEmptyObject';

const props = defineProps<{
  pattern: ISchemaPattern;
}>();

const emit = defineEmits<{
  'add-row': [];
}>();

const objectId = `${props.pattern.id}-empty`;
const object = useCanvasObject(objectId, () => new PatternEmptyObject());

onObjectEvent(object.parent!, 'mousedown', () => emit('add-row'));
</script>
