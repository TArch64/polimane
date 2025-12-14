<template>
  <foreignObject x="0" y="0" width="0" height="0">
    <Teleport to="body">
      <slot />
    </Teleport>
  </foreignObject>
</template>

<script setup lang="ts">
import type { Slot } from 'vue';

defineSlots<{
  default: Slot;
}>();
</script>
