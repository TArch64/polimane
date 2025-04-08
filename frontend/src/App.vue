<template>
  <RouterView />
</template>

<script setup lang="ts">
import { watch } from 'vue';
import { RouterView, useRoute } from 'vue-router';

const route = useRoute();

watch(() => route.name, (name, oldName) => {
  document.scrollingElement!.classList.add(`app--${name}`);
  if (oldName) document.scrollingElement!.classList.remove(`app--${oldName}`);
});
</script>
