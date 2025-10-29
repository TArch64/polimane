<template>
  <label class="form-field">
    <span class="form-field__label" v-if="label">
      {{ placeholder }}
    </span>

    <span class="form-field__container" :class="containerClasses">
      <slot />

      <span class="form-field__append" @click.stop.prevent v-if="slots.append">
        <slot name="append" />
      </span>
    </span>
  </label>
</template>

<script setup lang="ts">
import { computed, type Slot } from 'vue';

const props = withDefaults(defineProps<{
  placeholder: string;
  label?: boolean;
  variant?: 'main' | 'control';
}>(), {
  label: false,
  variant: 'main',
});

const slots = defineSlots<{
  default: Slot;
  append?: Slot;
}>();

const containerClasses = computed(() => [
  `form-field__container--variant-${props.variant}`,
]);
</script>

<style scoped>
@layer components {
  .form-field {
    display: block;
    max-width: 100%;
  }

  .form-field__label {
    display: block;
    margin-bottom: 2px;
    margin-left: 2px;
    font-size: var(--font-xs);
    color: color-mix(in srgb, var(--color-black), transparent 40%);
  }

  .form-field__container {
    border: var(--divider);
    border-radius: var(--rounded-md);
    padding: 4px 8px;
    transition: border-color 0.15s ease-out;
    will-change: border-color;
    display: flex;
    align-items: center;
    max-width: 100%;

    &:has(:focus:not(:user-invalid)) {
      border-color: var(--color-hover-divider);
    }

    &:has(:user-invalid) {
      border-color: var(--color-danger);
    }
  }

  .form-field__container--variant-main {
    background-color: var(--color-background-1);
  }

  .form-field__container--variant-control {
    background-color: var(--color-background-2);
  }

  .form-field__append {
    flex-shrink: 0;

    &:has(.button--icon) {
      display: flex;
    }

    :deep(.button--icon) {
      margin: -3px -6px -3px 0;
      padding: 3px;
      min-height: 0;
    }
  }
}
</style>
