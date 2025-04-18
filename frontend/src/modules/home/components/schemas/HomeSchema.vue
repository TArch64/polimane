<template>
  <Card interactable :binding="cardBinding">
    {{ schema.name }}
  </Card>
</template>

<script setup lang="ts">
import { Card } from '@/components/card';
import { RouterLink } from '@/components/button';
import type { ISchema } from '@/models';
import { makeBinding } from '@/components/binding';

const props = defineProps<{
  schema: ISchema;
}>();

const cardBinding = makeBinding(RouterLink, () => ({
  to: {
    name: 'schema-editor',
    params: { schemaId: props.schema.id },
  },
}));
</script>
