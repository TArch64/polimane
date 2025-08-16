import { type FunctionalComponent, h, resolveComponent } from 'vue';
import Konva from 'konva';
import { getThemeVar } from '@/composables';
import type { SchemaBeedCoordinate } from '@/models';
import { SCREENSHOT_IGNORE } from '../../composables';
import { BEAD_SIZE, type BeadOffset } from './useBeadsGrid';

export interface ICanvasBeadProps {
  offset: BeadOffset;
  coord: SchemaBeedCoordinate;
  color: string | null;
}

const BASE_CONFIG: Partial<Konva.RectConfig> = {
  name: SCREENSHOT_IGNORE,
  width: BEAD_SIZE - 2,
  height: BEAD_SIZE - 2,
  cornerRadius: getThemeVar('--rounded-full'),
  fill: getThemeVar('--color-background-3'),
};

export const CanvasBead: FunctionalComponent<ICanvasBeadProps> = (props) => {
  const KonvaRect = resolveComponent('KonvaRect');

  const config: Partial<Konva.RectConfig> = {
    ...BASE_CONFIG,
    $position: props.coord,
    x: props.offset[0] + 1,
    y: props.offset[1] + 1,
  };

  if (props.color) {
    config.name = undefined;
    config.fill = props.color;
  }

  return h(KonvaRect, { config });
};
