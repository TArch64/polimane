import { type FunctionalComponent, h, type SVGAttributes } from 'vue';
import { BEAD_SIZE, type BeadOffset } from '../../composables';

export interface IPreviewBeadProps {
  offset: BeadOffset;
  color: string;
}

const BASE_BEAD_ATTRS: SVGAttributes = {
  width: BEAD_SIZE - 2,
  height: BEAD_SIZE - 2,
  rx: BEAD_SIZE,
  ry: BEAD_SIZE,
};

export const PreviewBead: FunctionalComponent<IPreviewBeadProps> = (props) => h('rect', {
  ...BASE_BEAD_ATTRS,
  x: props.offset[0] + 1,
  y: props.offset[1] + 1,
  fill: props.color,
});
