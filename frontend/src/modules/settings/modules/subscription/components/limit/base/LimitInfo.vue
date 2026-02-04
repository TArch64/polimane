<template>
  <div class="subscription-limit__info">
    <h3 class="subscription-limit__title">
      {{ title }}
    </h3>

    <p class="subscription-limit__value">
      {{ value }}
    </p>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  title: string;
  value: string;
}>();
</script>

<style scoped>
@layer page {
  .subscription-limit__info {
    display: flex;
    align-items: center;
    justify-content: space-between;
  }

  .subscription-limit__title {
    font-size: 16px;
    font-weight: 400;
    text-align: center;
    text-wrap: balance;
  }

  .subscription-limit__value {
    font-size: 14px;
    text-align: center;
  }
}
</style>
