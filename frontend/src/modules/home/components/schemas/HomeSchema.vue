<template>
  <Card interactable :binding="cardBinding">
    {{ schema.name }}
  </Card>
</template>

<script setup lang="ts">
import { RouterLink } from 'vue-router';
import { Card } from '@/components/card';
import type { ISchema } from '@/models';
import { makeBinding } from '@/components/binding';

defineProps<{
  schema: ISchema;
}>();

const cardBinding = makeBinding(RouterLink, () => ({
  to: { name: 'home' },
  viewTransition: true,
}));
</script>

<style scoped>

</style>
