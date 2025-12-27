<template>
  <HomeListSchema
    disabled
    :schema
    :menuActions
  />
</template>

<script setup lang="ts">
import { HomeListSchema } from '@/modules/home/components';
import type { ListSchema } from '@/modules/home/stores';
import type { MaybeContextMenuAction } from '@/components/contextMenu';

defineProps<{
  schema: ListSchema;
}>();

const menuActions: MaybeContextMenuAction[] = [];
</script>
