<template>

</template>

<script setup lang="ts">
import { toRef } from 'vue';
import { useInfinityScroll } from '@/composables';
import { useFolderSchemasStore } from '../stores';

const schemasStore = useFolderSchemasStore();

useInfinityScroll({
  load: schemasStore.loadNext,
  canLoadNext: toRef(schemasStore, 'canLoadNext'),
});
</script>
