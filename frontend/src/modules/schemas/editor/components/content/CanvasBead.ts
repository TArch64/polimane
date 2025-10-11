import { type FunctionalComponent, h } from 'vue';
import { BEAD_CENTER, BEAD_RADIUS, type IBeadsGridItem } from '../../composables';

export interface ICanvasBeadProps {
  bead: IBeadsGridItem;
}

export const CanvasBead: FunctionalComponent<ICanvasBeadProps> = (props) => {
  return h('circle', {
    r: BEAD_RADIUS,
    coord: props.bead.coord,
    fill: props.bead.color,
    cx: props.bead.offset[0] + BEAD_CENTER,
    cy: props.bead.offset[1] + BEAD_CENTER,
  });
};
