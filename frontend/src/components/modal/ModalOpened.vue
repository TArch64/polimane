<template>
  <dialog ref="dialogRef" class="modal" @close="close">
    <header class="modal__header">
      <h1 class="modal__title">
        {{ title }}
      </h1>

      <Button icon variant="secondary" size="md" @click="close">
        <CloseIcon size="24" />
      </Button>
    </header>

    <Form @submit="$emit('save')">
      <main class="modal__body">
        <slot :close />
      </main>

      <footer class="modal__footer">
        <Button variant="primary" size="md" type="submit">
          {{ saveButton }}
        </Button>
      </footer>
    </Form>
  </dialog>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { Button } from '../button';
import { CloseIcon } from '../icon';
import { Form } from '../form';
import { onBackdropClick } from './onBackdropClick';
import type { ModalContentSlot } from './slots';

defineProps<{
  title: string;
  width: number;
  saveButton: string;
}>();

const emit = defineEmits<{
  closed: [];
  save: [];
}>();

defineSlots<{
  default: ModalContentSlot;
}>();

const dialogRef = ref<HTMLDialogElement>(null!);
const close = () => emit('closed');

onMounted(() => dialogRef.value.showModal());
onBackdropClick(dialogRef, close);
</script>

<style scoped>
@layer components {
  :global(body:has(.modal)) {
    overflow: hidden;
  }

  .modal {
    position: fixed;
    inset: 0;
    padding: 0;
    width: 100%;
    overflow-y: auto;
    max-width: v-bind("width + 'px'");
    max-height: 100%;
    background-color: var(--color-white);
    border: 1px solid var(--color-divider);
    border-radius: var(--rounded-md);

    &::backdrop {
      background-color: color-mix(in srgb, var(--color-black), transparent 30%);
    }
  }

  .modal__header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 8px 8px 4px 16px;
  }

  .modal__title {
    font-size: 18px;
    font-weight: 600;
  }

  .modal__body {
    padding: 8px 16px;
  }

  .modal__footer {
    padding: 8px 16px 12px;
    display: flex;
    justify-content: flex-end;
  }
}
</style>
