<template>
  <form @submit.prevent="$emit('submit')">
    <slot />
  </form>
</template>

<script setup lang="ts">
import type { Slot } from 'vue';

defineEmits<{
  submit: [];
}>();

defineSlots<{
  default: Slot;
}>();
</script>
