<template>
  <Card as="dialog" ref="dialogRef" class="modal" @close="close">
    <div class="modal__inner">
      <header class="modal__header">
        <h1 class="modal__title">
          {{ title }}
        </h1>

          <Button icon title="Закрити" class="modal__close" @click="close">
          <CloseIcon size="24" />
        </Button>
      </header>

      <Form @submit="$emit('save')">
        <main class="modal__body">
          <slot :close />
        </main>

        <footer class="modal__footer" v-if="footer">
          <Button
            variant="primary"
            type="submit"
            :loading
            :disabled="saveDisabled"
          >
            {{ saveButton }}
          </Button>
        </footer>
      </Form>
    </div>
  </Card>
</template>

<script setup lang="ts">
import { onMounted, type Slot } from 'vue';
import { onBackdropClick, useDomRef } from '@/composables';
import { Card } from '../card';
import { CloseIcon } from '../icon';
import { Form } from '../form';
import { Button } from '../button';
import { ModalWidth } from './ModalWidth';
import { useActiveModal } from './useActiveModal';

withDefaults(defineProps<{
  title: string;
  width?: number;
  saveButton?: string;
  saveDisabled?: boolean;
  loading?: boolean;
  footer?: boolean;
}>(), {
  width: ModalWidth.MD,
  saveButton: 'Зберегти',
  saveDisabled: false,
  footer: true,
});

defineEmits<{
  save: [];
}>();

defineSlots<{
  default: Slot;
}>();

const modal = useActiveModal();
const close = () => modal.close(null);

const dialogRef = useDomRef<HTMLDialogElement>();

onMounted(() => dialogRef.value.showModal());
onBackdropClick(close);
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
    width: calc(100% - 32px);
    overflow-y: auto;
    max-width: v-bind("width + 'px'");
    max-height: 100%;
    view-transition-name: modal;

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

  .modal__close.modal__close {
    padding: 4px;
  }

  .modal__title {
    font-size: 18px;
    font-weight: 550;
  }

  .modal__body {
    padding: 8px 16px;

    &:not(:has(+ .modal__footer)) {
      padding-bottom: 12px;
    }
  }

  .modal__footer {
    padding: 8px 16px 12px;
    display: flex;
    justify-content: flex-end;
  }
}
</style>
