import { defineComponent, h, type PropType } from 'vue';
import type { Modal } from './Modal';
import { provideActiveModal } from './useActiveModal';

export const ModalRender = defineComponent({
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
