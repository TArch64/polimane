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
    animation: button-loading-icon 3s linear infinite;
  }

  @keyframes button-loading-icon {
    from {
      rotate: 0deg
    }
    to {
      rotate: 360deg
    }
  }
}
</style>
