import { defineComponent, h, type PropType } from 'vue';
import { CanvasBead } from './CanvasBead';
import type { BeadGridGenerator } from './useBeadsGrid';

export const CanvasBeadGrid = defineComponent({
  name: 'CanvasBeadGrid',
  components: { CanvasBead },

  props: {
    grid: {
      type: Object as PropType<BeadGridGenerator>,
      required: true,
    },
  },

  setup(props) {
    return () => Array.from(props.grid, (item) => h(CanvasBead, item));
  },
});
