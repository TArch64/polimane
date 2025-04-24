<template />

<script setup lang="ts">
import type { ISchemaPattern } from '@/models';
import { useModal } from '@/components/modal';
import { onObjectClick, useCanvasObject } from '@/modules/schemas/editor/composables';
import { getPatternAddRowModal } from '../../modals';
import { PatternEmptyObject } from './PatternEmptyObject';

const props = defineProps<{
  pattern: ISchemaPattern;
}>();

const addModal = useModal(getPatternAddRowModal(props.pattern));

const objectId = `${props.pattern.id}-empty`;
const object = useCanvasObject(objectId, () => new PatternEmptyObject());

onObjectClick(object.parent!, () => {
  addModal.open({ pattern: props.pattern });
});
</script>
