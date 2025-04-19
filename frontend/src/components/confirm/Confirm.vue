<template>
  <dialog
    ref="dialogRef"
    popover="manual"
    class="confirm"
    @keydown.esc="decline"
  >
    <p class="confirm__message">
      {{ model.message }}
    </p>

    <footer class="confirm__footer">
      <ConfirmButton
        variant="secondary"
        :button="model.declineButton"
        @click="decline"
      />

      <ConfirmButton
        variant="primary"
        :button="model.acceptButton"
        :danger="model.danger"
        @click="accept"
      />
    </footer>
  </dialog>
</template>

<script setup lang="ts">
import { type FunctionalComponent, h, onMounted, ref } from 'vue';
import { onBackdropClick } from '@/composables';
import { Button, type ButtonVariant } from '../button';
import type { Confirm } from './Confirm';

const props = defineProps<{
  model: Confirm;
}>();

interface IConfirmButtonProps {
  variant: ButtonVariant;
  button: string;
  danger?: boolean;
}

const ConfirmButton: FunctionalComponent<IConfirmButtonProps> = (props) => {
  return h(Button, {
    size: 'md',
    variant: props.variant,
    danger: props.danger,
  }, () => props.button);
};

const dialogRef = ref<HTMLDialogElement>(null!);
const decline = () => props.model.complete(false);
const accept = () => props.model.complete(true);

onMounted(() => dialogRef.value.showPopover());
onBackdropClick(dialogRef, decline);
</script>

<style scoped>
@layer components {
  .confirm {
    position-anchor: v-bind("model.anchorVar");
    position-area: bottom center;
    margin-top: 4px;
    width: 300px;
    padding: 12px;
    overflow-y: auto;
    background-color: var(--color-background-2);
    border: var(--divider);
    border-radius: var(--rounded-md);
    box-shadow: var(--box-shadow);
    view-transition-name: modal;
  }

  .confirm__message {
    margin-bottom: 4px;
  }

  .confirm__footer {
    display: flex;
    justify-content: flex-end;
    gap: 8px;
  }
}
</style>
