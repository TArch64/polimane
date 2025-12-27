<template>
  <HomeListEmpty>
    <template #description>
      <p>Наразі немає видалених схем для бісеру</p>
      <p>Всі видалені схеми будуть зберігатися тут 30 днів</p>
    </template>
  </HomeListEmpty>
</template>

<script setup lang="ts">
import { HomeListEmpty } from '@/modules/home/components';
</script>
