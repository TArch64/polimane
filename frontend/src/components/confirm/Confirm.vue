<template>
  <dialog
    ref="dialogRef"
    popover="manual"
    class="confirm"
    @keydown.esc="decline"
    v-popover-shift
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
import { computed, type FunctionalComponent, h, nextTick, onMounted, ref } from 'vue';
import { onBackdropClick } from '@/composables';
import { vPopoverShift } from '@/directives';
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

ConfirmButton.displayName = 'ConfirmButton';

const dialogRef = ref<HTMLDialogElement>(null!);
const decline = () => props.model.complete(false);
const accept = () => props.model.complete(true);

const backgroundColor = computed(() => {
  return props.model.control ? 'var(--color-background-2)' : 'var(--color-background-1)';
});

onMounted(async () => {
  dialogRef.value.showPopover();
  await nextTick();
});

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
    background-color: v-bind("backgroundColor");
    border: var(--divider);
    border-radius: var(--rounded-md);
    box-shadow: var(--box-shadow);
    view-transition-name: confirm;
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
