<template>
  <img
    :src
    :alt
    draggable="false"
    decoding="async"
    loading="lazy"
    class="home-list-screenshot"
    v-if="src"
  >

  <div class="home-list-screenshot" v-else />
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { buildCdnUrl } from '@/helpers/buildCdnUrl';

const props = defineProps<{
  alt: string;
  path: string | null;
  backgroundColor: string;
}>();

const src = computed(() => buildCdnUrl(props.path));
</script>

<style scoped>
@layer page {
  .home-list-screenshot {
    display: block;
    aspect-ratio: 16 / 9;
    object-fit: contain;
    object-position: center;
    border-bottom: var(--divider);
    margin: calc(0px - var(--card-padding-top)) calc(0px - var(--card-padding-right)) 8px calc(0px - var(--card-padding-left));
    width: calc(100% + var(--card-padding-left) + var(--card-padding-right));
    background-color: v-bind("backgroundColor");
  }
}
</style>
