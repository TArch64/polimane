<template />

<script setup lang="ts">
import { onObjectClick, useCanvasObject } from '@/modules/schemas/editor/composables';
import type { ISchemaPattern } from '@/models';
import { useModal } from '@/components/modal';
import EditorAddSquareRowModal
  from '@/modules/schemas/editor/components/EditorAddSquareRowModal.vue';
import { PatternEmptyObject } from './PatternEmptyObject';

const props = defineProps<{
  pattern: ISchemaPattern;
}>();

const addModal = useModal(EditorAddSquareRowModal);

const objectId = `${props.pattern.id}-empty`;
const object = useCanvasObject(objectId, () => new PatternEmptyObject());

onObjectClick(object.parent!, () => {
  addModal.open({ pattern: props.pattern });
});
</script>
