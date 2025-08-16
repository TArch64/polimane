import { computed, defineComponent, h, type PropType } from 'vue';
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
    const getColor = (coord: SchemaBeedCoordinate) => editorStore.schema.beads[coord] ?? null;
    const beads = computed(() => Array.from(props.grid));

    return () => beads.value.map((item) => {
      return h(CanvasBead, {
        offset: item.offset,
        coord: item.coord,
        color: getColor(item.coord),
      });
    });
  },
});
