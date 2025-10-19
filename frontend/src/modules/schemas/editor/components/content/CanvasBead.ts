import { type FunctionalComponent, h } from 'vue';
import { BEAD_BUGLE_HITBOX_SIZE, BEAD_CIRCLE_RADIUS, BEAD_REF_HITBOX_SIZE } from '@editor/const';
import { getBeadSettings, type SchemaBead } from '@/models';
import { BeadKind } from '@/enums';
import type {
  IBeadsGridBugle,
  IBeadsGridCircle,
  IBeadsGridItem,
  IBeadsGridRef,
} from '../../composables';

export interface ICanvasBeadProps {
  item: IBeadsGridItem;
}

type BeadComponent = FunctionalComponent<ICanvasBeadProps>;

const CanvasBeadCircle: BeadComponent = (props) => {
  const { center } = props.item.precomputed as IBeadsGridCircle;
  const settings = getBeadSettings(props.item.bead as SchemaBead<BeadKind.CIRCLE>);

  return h('circle', {
    r: BEAD_CIRCLE_RADIUS,
    coord: props.item.coord,
    fill: settings.color,
    cx: center.x,
    cy: center.y,
  });
};

CanvasBeadCircle.displayName = 'CanvasBeadCircle';

const CanvasBeadBugle: BeadComponent = (props) => {
  const { shape } = props.item.precomputed as IBeadsGridBugle;
  const settings = getBeadSettings(props.item.bead as SchemaBead<BeadKind.BUGLE>);

  return [
    h('rect', {
      class: 'canvas-bead-bugle',
      x: shape.x,
      y: shape.y,
      width: shape.width,
      height: shape.height,
      fill: settings.color,
    }),
    h('rect', {
      x: shape.x,
      y: shape.y,
      width: BEAD_BUGLE_HITBOX_SIZE,
      height: BEAD_BUGLE_HITBOX_SIZE,
      fill: 'transparent',
      coord: props.item.coord,
    }),
  ];
};

CanvasBeadBugle.displayName = 'CanvasBeadBugle';

const CanvasBeadRef: BeadComponent = (props) => {
  const { hitbox } = props.item.precomputed as IBeadsGridRef;

  return h('rect', {
    class: 'canvas-bead-ref',
    x: hitbox.x,
    y: hitbox.y,
    width: BEAD_REF_HITBOX_SIZE,
    height: BEAD_REF_HITBOX_SIZE,
    fill: 'transparent',
    coord: props.item.coord,
  });
};

CanvasBeadRef.displayName = 'CanvasBeadRef';

const beadComponentMap: Record<BeadKind, BeadComponent> = {
  [BeadKind.CIRCLE]: CanvasBeadCircle,
  [BeadKind.BUGLE]: CanvasBeadBugle,
  [BeadKind.REF]: CanvasBeadRef,
};

export const CanvasBead: BeadComponent = (props) => (
  h(beadComponentMap[props.item.bead.kind], props)
);
