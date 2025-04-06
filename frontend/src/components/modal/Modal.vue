<template>
  <slot name="activator" :open />

  <Teleport to="body" v-if="isOpened">
    <ModalOpened
      :title
      :width
      :saveButton
      @closed="close"
      @save="$emit('save')"
      v-slot="ctx"
    >
      <slot v-bind="ctx" />
    </ModalOpened>
  </Teleport>
</template>

<script setup lang="ts">
import { nextTick, ref } from 'vue';
import { ModalWidth } from './ModalWidth';
import ModalOpened from './ModalOpened.vue';
import type { ModalActivatorSlot, ModalContentSlot } from './slots';

withDefaults(defineProps<{
  title: string;
  width?: number;
  saveButton?: string;
}>(), {
  width: ModalWidth.MD,
  saveButton: 'Зберегти',
});

defineEmits<{
  save: [];
}>();

defineSlots<{
  activator: ModalActivatorSlot;
  default: ModalContentSlot;
}>();

const isOpened = ref(false);

function toggle(toOpen: boolean): void {
  document.startViewTransition(() => {
    isOpened.value = toOpen;
    return nextTick();
  });
}

const open = () => toggle(true);
const close = () => toggle(false);

defineExpose({
  close,
  open,
});
</script>
