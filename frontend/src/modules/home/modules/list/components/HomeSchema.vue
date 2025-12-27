<template>
  <HomeListSchema :schema :menuActions />
</template>

<script setup lang="ts">
import { HomeListSchema } from '@/modules/home/components';
import { useSchemaMenuActions } from '@/modules/home/composables';
import type { ListSchema } from '@/modules/home/stores';

const props = defineProps<{
  schema: ListSchema;
}>();

const menuActions = useSchemaMenuActions(() => props.schema);
</script>
