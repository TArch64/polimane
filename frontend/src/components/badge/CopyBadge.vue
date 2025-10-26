<template>
  <Badge interactable :binding class="copy-badge">
    <span class="copy-badge__text">
      {{ text }}
    </span>

    <CheckmarkIcon
      size="inline"
      class="copy-badge__icon"
      v-if="copy.isActive"
    />
  </Badge>
</template>

<script setup lang="ts">
import { useAsyncAction } from '@/composables';
import { wait } from '@/helpers';
import { makeBinding } from '../binding';
import { CheckmarkIcon } from '../icon';
import Badge from './Badge.vue';

const props = defineProps<{
  text: string;
}>();

const copy = useAsyncAction(async () => {
  await navigator.clipboard.writeText(props.text);
  return wait(5000);
});

const binding = makeBinding('button', () => ({
  type: 'button',
  disabled: copy.isActive,
  onClick: copy,
}));
</script>

<style scoped>
@layer components {
  .copy-badge {
    display: flex;
    align-items: center;
    line-height: 1;
    max-width: 100%;
    transition: width 0.1s ease-out;
  }

  .copy-badge__text {
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
    max-width: 100%;

    &:has(+ .copy-badge__icon) {
      max-width: calc(100% - 1em - 3px);
    }
  }

  .copy-badge__icon {
    margin-left: 3px;
  }
}
</style>
