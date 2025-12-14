<template>
  <Teleport to="#homeTopBarRouteActions">
    <slot />
  </Teleport>
</template>

<script setup lang="ts">
import type { Slot } from 'vue';

defineSlots<{
  default: Slot;
}>();
</script>
