import { type FunctionalComponent, h } from 'vue';
import { BEAD_CIRCLE_RADIUS } from '@editor/const';
import { getBeadSettings, type SchemaBead } from '@/models';
import { type BeadContentKind, BeadKind } from '@/enums';
import type { IBeadsGridBugle, IBeadsGridCircle, IBeadsGridItem } from '../../composables';

export interface ICanvasBeadProps {
  item: IBeadsGridItem;
}

type BeadComponent = FunctionalComponent<ICanvasBeadProps>;

export const CanvasBeadCircle: BeadComponent = (props) => {
  const precomputed = props.item.precomputed as IBeadsGridCircle;
  const settings = getBeadSettings(props.item.bead as SchemaBead<BeadKind.CIRCLE>);

  return h('circle', {
    r: BEAD_CIRCLE_RADIUS,
    coord: props.item.coord,
    fill: settings.color,
    cx: precomputed.center.x,
    cy: precomputed.center.y,
  });
};

CanvasBeadCircle.displayName = 'CanvasBeadCircle';

export const CanvasBeadBugle: BeadComponent = (props) => {
  const precomputed = props.item.precomputed as IBeadsGridBugle;
  const settings = getBeadSettings(props.item.bead as SchemaBead<BeadKind.BUGLE>);

  return h('rect', {
    class: 'canvas-bead-bugle',
    x: precomputed.x,
    y: precomputed.y,
    width: precomputed.width,
    height: precomputed.height,
    coord: props.item.coord,
    fill: settings.color,
  });
};

CanvasBeadBugle.displayName = 'CanvasBeadBugle';

const beadComponentMap: Record<BeadContentKind, BeadComponent> = {
  [BeadKind.CIRCLE]: CanvasBeadCircle,
  [BeadKind.BUGLE]: CanvasBeadBugle,
};

export const CanvasBead: BeadComponent = (props) => (
  h(beadComponentMap[props.item.bead.kind], props)
);
