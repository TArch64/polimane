import { defineComponent, h, type PropType } from 'vue';
import type { SchemaBeedCoordinate } from '@/models';
import { useEditorStore } from '../../stores';
import { CanvasBead } from './CanvasBead';
import type { BeadGridGenerator } from './useBeadsGrid';

export const CanvasBeadGrid = defineComponent({
  name: 'CanvasBeadGrid',

  props: {
    grid: {
      type: Object as PropType<BeadGridGenerator>,
      required: true,
    },
  },

  setup(props) {
    const editorStore = useEditorStore();
    const getColor = (pos: SchemaBeedCoordinate) => editorStore.schema.beads[pos] ?? null;

    return () => Array.from(props.grid, (item) => {
      return h(CanvasBead, {
        offset: item.offset,
        position: item.position,
        color: getColor(item.position),
      });
    });
  },
});
