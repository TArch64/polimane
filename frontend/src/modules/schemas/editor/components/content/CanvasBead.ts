import { type FunctionalComponent, h, type SVGAttributes } from 'vue';
import type { SchemaBeadCoord } from '@/models';
import { BEAD_SIZE, type BeadOffset } from '../../composables';

export interface ICanvasBeadProps {
  offset: BeadOffset;
  coord: SchemaBeadCoord;
  color: string | null;
  emptyColor: string;
}

const CENTER = BEAD_SIZE / 2;

const BASE_BEAD_CONFIG: Partial<SVGAttributes> = {
  r: CENTER - 1,
};

export const CanvasBead: FunctionalComponent<ICanvasBeadProps> = (props) => {
  return h('circle', {
    ...BASE_BEAD_CONFIG,
    coord: props.coord,
    cx: props.offset[0] + 1 + CENTER,
    cy: props.offset[1] + 1 + CENTER,
    fill: props.color ? props.color : props.emptyColor,
  });
};
