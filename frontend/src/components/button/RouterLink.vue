<template>
  <a :href @click="navigate">
    <slot />
  </a>
</template>

<script setup lang="ts">
import { type RouteLocationRaw, useLink } from 'vue-router';
import { type Slot, toRef } from 'vue';

const props = defineProps<{
  to: RouteLocationRaw;
}>();

defineSlots<{
  default: Slot;
}>();

const { href, navigate } = useLink({
  viewTransition: true,
  to: toRef(props, 'to'),
});
</script>
