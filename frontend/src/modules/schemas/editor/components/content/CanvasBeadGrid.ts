import { defineComponent, h, type PropType } from 'vue';
import { useEditorStore } from '../../stores';
import { CanvasBead } from './CanvasBead';
import type { BeadGridGenerator, BeadPosition } from './useBeadsGrid';

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
    const editorStore = useEditorStore();
    const getColor = (pos: BeadPosition) => editorStore.schema.beads[`${pos[0]}:${pos[1]}`] ?? null;

    return () => Array.from(props.grid, (item) => {
      return h(CanvasBead, {
        offset: item.offset,
        color: getColor(item.position),
      });
    });
  },
});
