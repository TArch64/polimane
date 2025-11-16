<template>
  <p>
    {{ folderId }}
  </p>
</template>

<script setup lang="ts">
import { useHomeStore } from '@/modules/home/stores';

const props = defineProps<{
  folderId: string;
}>();

const homeStore = useHomeStore();

homeStore.setRouteConfig({
  title: props.folderId,
});
</script>

<style scoped>

</style>
