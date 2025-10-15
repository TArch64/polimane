import { type FunctionalComponent, h } from 'vue';
import { BEAD_RADIUS } from '@editor/const';
import type { IBeadsGridItem } from '../../composables';

export interface ICanvasBeadProps {
  item: IBeadsGridItem;
}

export const CanvasBead: FunctionalComponent<ICanvasBeadProps> = (props) => (
  h('circle', {
    r: BEAD_RADIUS,
    coord: props.item.coord,
    fill: props.item.bead.color,
    cx: props.item.offset.x,
    cy: props.item.offset.y,
  })
);
