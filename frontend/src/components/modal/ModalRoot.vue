<template>
  <Teleport to="body" v-if="openedModal">
    <ModalRender :key="openedModal.id" :modal="openedModal as Modal" />
  </Teleport>
</template>

<script setup lang="ts">
import { defineComponent, h, nextTick, type PropType, ref, watch } from 'vue';
import { ModalPlugin } from './ModalPlugin';
import type { Modal } from './Modal';
import { provideActiveModal } from './useActiveModal';

const plugin = ModalPlugin.inject();
const openedModal = ref<Modal | null>(null);

watch(() => plugin.openedModal?.id, () => {
  document.startViewTransition(() => {
    openedModal.value = plugin.openedModal ?? null;
    return nextTick();
  });
});

const ModalRender = defineComponent({
  props: {
    modal: {
      type: Object as PropType<Modal>,
      required: true,
    },
  },
  setup(props) {
    provideActiveModal(props.modal);
    return () => h(props.modal.component, props.modal.props);
  },
});
</script>
