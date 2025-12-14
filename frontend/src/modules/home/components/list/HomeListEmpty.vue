<template>
  <div class="home-list-empty">
    <p class="home-list-empty__description">
      {{ description }}
    </p>

    <div class="home-list-empty__actions">
      <slot />
    </div>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  description: string;
}>();

</script>

<style scoped>
@layer page {
  .home-list-empty {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    flex-grow: 1;
  }

  .home-list-empty__description {
    margin-bottom: 12px;
  }

  .home-list-empty__actions {
    display: flex;
    gap: 8px;
    align-items: center;
  }
}
</style>
