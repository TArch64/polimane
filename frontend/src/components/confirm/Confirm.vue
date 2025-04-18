<template>
  <dialog
    ref="dialogRef"
    popover="manual"
    class="popover"
    @keydown.esc="decline"
  >
    <p class="popover__message">
      {{ model.message }}
    </p>

    <footer class="popover__footer">
      <ConfirmButton variant="secondary" :button="model.declineButton" @click="decline" />
      <ConfirmButton variant="primary" :button="model.acceptButton" @click="accept" />
    </footer>
  </dialog>
</template>

<script setup lang="ts">
import { type FunctionalComponent, h, onMounted, onUnmounted, ref } from 'vue';
import { onBackdropClick } from '@/composables';
import { Button, type ButtonVariant } from '../button';
import type { Confirm, IConfirmButton } from './Confirm';

const props = defineProps<{
  model: Confirm;
}>();

interface IConfirmButtonProps {
  variant: ButtonVariant;
  button: IConfirmButton;
}

const ConfirmButton: FunctionalComponent<IConfirmButtonProps> = (props) => {
  return h(Button, {
    size: 'md',
    variant: props.variant,
    danger: props.button.danger,
  }, () => props.button.text);
};

const dialogRef = ref<HTMLDialogElement>(null!);
const decline = () => props.model.complete(false);
const accept = () => props.model.complete(true);

onMounted(() => {
  props.model.anchorEl?.style.setProperty('anchor-name', props.model.anchorVar);
  dialogRef.value.showPopover();
});

onUnmounted(() => {
  props.model.anchorEl?.style.removeProperty('anchor-name');
});

onBackdropClick(dialogRef, decline);
</script>

<style scoped>
@layer components {
  .popover {
    view-transition-name: confirm;
    position-anchor: v-bind("model.anchorVar");
    position-area: bottom center;
    margin-top: 12px;
    width: 300px;
    padding: 12px;
    overflow-y: auto;
    background-color: var(--color-background-2);
    border: var(--divider);
    border-radius: var(--rounded-md);
    box-shadow: 0 1px 2px 0 rgba(0, 0, 0, 0.05);
  }

  .popover__message {
    margin-bottom: 4px;
  }

  .popover__footer {
    display: flex;
    justify-content: flex-end;
    gap: 8px;
  }
}
</style>
