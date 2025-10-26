<template>
  <LoaderIcon v-bind="props" class="spinner" />
</template>

<script setup lang="ts">
import { type IIconProps, LoaderIcon } from './icon';

const props = withDefaults(defineProps<IIconProps>(), {
  size: undefined,
});
</script>

<style scoped>
@layer components {
  .spinner {
    animation: spinner 3s linear infinite;
  }

  @keyframes spinner {
    from {
      rotate: 0deg
    }
    to {
      rotate: 360deg
    }
  }
}
</style>
