<template>
  <Teleport to="body" v-if="openedModal">
    <ModalRender :key="openedModal.id" :modal="openedModal as Modal" />
  </Teleport>
</template>

<script setup lang="ts">
import { defineComponent, h, nextTick, type PropType, ref, watch } from 'vue';
import { useRouteTransition } from '@/composables';
import { ModalPlugin } from './ModalPlugin';
import type { Modal } from './Modal';
import { provideActiveModal } from './useActiveModal';

const plugin = ModalPlugin.inject();
const routeTransition = useRouteTransition();
const openedModal = ref<Modal | null>(null);

watch(() => plugin.openedModal?.id, () => {
  routeTransition.start(() => {
    const promises = [];

    if (openedModal.value?.onClose) {
      promises.push(openedModal.value.onClose());
    }

    openedModal.value = plugin.openedModal ?? null;
    return Promise.all([nextTick(), ...promises]);
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
