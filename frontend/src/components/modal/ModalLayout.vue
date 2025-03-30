<template>
  <dialog class="modal" @close="$emit('close')">
    <header class="modal__header">
      <h1 class="modal__title">
        {{ title }}
      </h1>

      <Button icon variant="secondary" size="md" @click="$emit('close')">
        <CloseIcon size="24" />
      </Button>
    </header>

    <main class="modal__body">
      <slot />
    </main>
  </dialog>
</template>

<script setup lang="ts">
import type { Slot } from 'vue';
import { Button } from '@/components/button';
import { CloseIcon } from '@/components/icon';

defineProps<{
  title: string;
}>();

defineEmits<{
  close: [];
}>();

defineSlots<{
  default: Slot;
}>();
</script>

<style scoped>
.modal {
  background-color: var(--color-white);
  border: 1px solid var(--color-divider);
  border-radius: var(--rounded-md);
}

.modal__header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 4px 8px 4px 16px;
}

.modal__title {
  font-size: 18px;
  font-weight: 600;
}

.modal__body {
  padding: 4px 16px 8px;
}
</style>
