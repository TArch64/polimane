<template>
  <p class="sortable-drag-preview">
    {{ title }}
  </p>
</template>

<script lang="ts" setup>
defineProps<{
  title: string;
}>();
</script>

<style scoped>
@layer components {
  .sortable-drag-preview {
    padding: 4px 16px;
    background-color: var(--color-white);
    box-shadow: var(--box-shadow);
    border-radius: var(--rounded-md);
  }
}
</style>
