<template>
  <dialog
    ref="dialogRef"
    popover="manual"
    class="confirm"
    @keydown.esc="decline"
    v-popover-shift.defer
  >
    <p class="confirm__message">
      {{ message }}
    </p>

    <CheckboxField
      :label="model.additionalCondition"
      class="confirm__additional-condition"
      v-model="isAdditionalAccepted"
      v-if="model.additionalCondition"
    />

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
import { computed, type FunctionalComponent, h, nextTick, onMounted, ref, toValue } from 'vue';
import { onBackdropClick } from '@/composables';
import { vPopoverShift } from '@/directives';
import { CheckboxField } from '@/components/form';
import { Button, type ButtonVariant } from '../button';
import type { ConfirmModel } from './ConfirmModel';

const props = defineProps<{
  model: ConfirmModel;
}>();

interface IConfirmButtonProps {
  variant: ButtonVariant;
  button: string;
  danger?: boolean;
}

const isAdditionalAccepted = ref(false);

const ConfirmButton: FunctionalComponent<IConfirmButtonProps> = (props) => {
  return h(Button, {
    size: 'md',
    variant: props.variant,
    danger: props.danger,
  }, () => props.button);
};

const dialogRef = ref<HTMLDialogElement>(null!);

function decline() {
  props.model.complete({
    isAccepted: false,
  });
}

function accept() {
  props.model.complete({
    isAccepted: true,
    isSecondaryAccepted: isAdditionalAccepted.value,
  });
}

const message = computed(() => toValue(props.model.message));

const backgroundColor = computed(() => {
  return props.model.control ? 'var(--color-background-2)' : 'var(--color-background-1)';
});

onMounted(async () => {
  dialogRef.value.showPopover();
  await nextTick();
});

onBackdropClick(decline);
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
    border-radius: var(--rounded-lg);
    box-shadow: var(--box-shadow);
    view-transition-name: confirm;
  }

  .confirm__message:not(:has(+ .confirm__additional-condition)) {
    margin-bottom: 4px;
  }

  .confirm__additional-condition {
    margin-top: 12px;
    margin-bottom: 8px;
  }

  .confirm__footer {
    display: flex;
    justify-content: flex-end;
    gap: 8px;
  }
}
</style>
