<template>
  <Teleport to="body" v-if="openedModal">
    <ModalRender :key="openedModal.id" :modal="openedModal as Modal" />
  </Teleport>
</template>

<script setup lang="ts">
import { defineComponent, h, nextTick, type PropType, type Ref, ref, watch } from 'vue';
import { useRouteTransition } from '@/composables';
import { ModalPlugin } from './ModalPlugin';
import type { Modal } from './Modal';
import { provideActiveModal } from './useActiveModal';

const plugin = ModalPlugin.inject();
const routeTransition = useRouteTransition();
const openedModal: Ref<Modal | null> = ref(null);

watch(() => plugin.openedModal?.id, () => {
  routeTransition.start(async () => {
    const previousModal = openedModal.value;
    openedModal.value = plugin.openedModal;
    await nextTick();
    await previousModal?.completeClose();
  });
});

const ModalRender = defineComponent({
  name: 'ModalRender',

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
