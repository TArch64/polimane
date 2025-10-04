import { type FunctionalComponent, h } from 'vue';
import type { SchemaBeadCoord } from '@/models';
import { BEAD_CENTER, BEAD_RADIUS, type BeadOffset } from '../../composables';

export interface ICanvasBeadProps {
  offset: BeadOffset;
  coord: SchemaBeadCoord;
  color: string | null;
}

export const CanvasBead: FunctionalComponent<ICanvasBeadProps> = (props) => {
  return h('circle', {
    r: BEAD_RADIUS,
    coord: props.coord,
    fill: props.color,
    cx: props.offset[0] + BEAD_CENTER,
    cy: props.offset[1] + BEAD_CENTER,
  });
};
