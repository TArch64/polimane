import { defineComponent, h, type PropType } from 'vue';
import type { ModalModel } from './ModalModel';
import { provideActiveModal } from './useActiveModal';

export const ModalRender = defineComponent({
  name: 'ModalRender',

  props: {
    modal: {
      type: Object as PropType<ModalModel>,
      required: true,
    },
  },

  setup(props) {
    provideActiveModal(props.modal);
    return () => h(props.modal.component, props.modal.props);
  },
});
