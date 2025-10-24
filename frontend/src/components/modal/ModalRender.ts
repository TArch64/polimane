import { defineComponent, h, type PropType } from 'vue';
import { provideTopElement, useDomRef } from '@/composables';
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
    const topEl = useDomRef<HTMLElement>();

    provideActiveModal(props.modal);
    provideTopElement(topEl);

    return () => h(props.modal.component, {
      ...props.modal.props,
      ref: topEl,
    });
  },
});
