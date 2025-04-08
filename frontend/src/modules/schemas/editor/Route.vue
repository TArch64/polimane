<template>
  editor
</template>

<script lang="ts" setup>
import { useRoute } from 'vue-router';

const props = defineProps<{
  schemaId: number;
}>();

console.log(props);

const route = useRoute<'schema-editor'>();
console.log(route.params);
</script>
