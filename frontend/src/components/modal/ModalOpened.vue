<template>
  <ModalLayout
    :title
    class="dialog"
    ref="dialogRef"
    @close="close"
  >
    <slot :close />
  </ModalLayout>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';
import type { ComponentExposed } from 'vue-component-type-helpers';
import { onBackdropClick } from './onBackdropClick';
import type { ModalContentSlot } from './slots';
import ModalLayout from './ModalLayout.vue';

defineProps<{
  title: string;
  width: number;
}>();

const emit = defineEmits<{
  closed: [];
}>();

defineSlots<{
  default: ModalContentSlot;
}>();

const dialogRef = ref<ComponentExposed<typeof ModalLayout>>(null!);
const dialogEl = computed((): HTMLDialogElement => dialogRef.value.$el);
const close = () => emit('closed');

onMounted(() => dialogEl.value.showModal());
onBackdropClick(dialogEl, close);
</script>

<style scoped>
@layer components {
  :global(body:has(.dialog)) {
    overflow: hidden;
  }

  .dialog {
    position: fixed;
    inset: 0;
    padding: 0;
    width: 100%;
    overflow-y: auto;
    max-width: v-bind("width + 'px'");
    max-height: 100%;
    border: none;
    background-color: transparent;

    &::backdrop {
      background-color: color-mix(in srgb, var(--color-black), transparent 30%);
    }
  }
}
</style>
