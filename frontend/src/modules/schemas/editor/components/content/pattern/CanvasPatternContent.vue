<template>
  <CanvasStackV :gap="12" align="center">
    <CanvasRow
      v-for="row in pattern.content"
      :key="row.id"
      :row
    />
  </CanvasStackV>
</template>

<script setup lang="ts">
import type { ISchemaPattern } from '@/models';
import { CanvasStackV } from '../base';
import { CanvasRow } from '../row';

defineProps<{
  pattern: ISchemaPattern;
}>();
</script>
