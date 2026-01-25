<template>
  <Card variant="control">
    {{ plan }}
  </Card>
</template>

<script setup lang="ts">
import { Card } from '@/components/card';
import type { ISubscriptionPlan } from '@/models';

defineProps<{
  plan: ISubscriptionPlan;
}>();
</script>

<style scoped>
@layer components {

}
</style>
