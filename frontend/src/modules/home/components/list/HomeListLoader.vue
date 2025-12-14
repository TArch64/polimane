<template>
  <div class="home-list-loader" v-visible="visible">
    <Spinner />
  </div>
</template>

<script setup lang="ts">
import Spinner from '@/components/Spinner.vue';
import { vVisible } from '@/directives';

defineProps<{
  visible: boolean;
}>();
</script>

<style scoped>
@layer page {
  .home-list-loader {
    display: flex;
    justify-content: center;
    padding: 20px 12px;
  }
}
</style>
