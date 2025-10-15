import { type FunctionalComponent, h } from 'vue';
import { BEAD_RADIUS } from '@editor/const';
import type { IBeadsGridItem } from '../../composables';

export interface ICanvasBeadProps {
  bead: IBeadsGridItem;
}

export const CanvasBead: FunctionalComponent<ICanvasBeadProps> = (props) => (
  h('circle', {
    r: BEAD_RADIUS,
    coord: props.bead.coord,
    fill: props.bead.color,
    cx: props.bead.offset.x,
    cy: props.bead.offset.y,
  })
);
