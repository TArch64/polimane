import { type FunctionalComponent, h, resolveComponent } from 'vue';
import Konva from 'konva';
import { getThemeVar } from '@/composables';
import type { SchemaBeadCoord } from '@/models';
import { BEAD_SIZE, type BeadOffset } from '../../composables';

export interface ICanvasBeadProps {
  offset: BeadOffset;
  coord: SchemaBeadCoord;
  color: string | null;
}

const BASE_BACKGROUND_CONFIG: Partial<Konva.RectConfig> = {
  width: BEAD_SIZE,
  height: BEAD_SIZE,
  fill: getThemeVar('--color-background-2'),
};

const BASE_BEAD_CONFIG: Partial<Konva.RectConfig> = {
  width: BEAD_SIZE - 2,
  height: BEAD_SIZE - 2,
  cornerRadius: getThemeVar('--rounded-full'),
  fill: getThemeVar('--color-background-3'),
};

export const CanvasBead: FunctionalComponent<ICanvasBeadProps> = (props) => {
  const KonvaRect = resolveComponent('KonvaRect');

  const backgroundConfig: Partial<Konva.RectConfig> = {
    ...BASE_BACKGROUND_CONFIG,
    $position: props.coord,
    x: props.offset[0],
    y: props.offset[1],
  };

  const beadConfig: Partial<Konva.RectConfig> = {
    ...BASE_BEAD_CONFIG,
    $position: props.coord,
    x: props.offset[0] + 1,
    y: props.offset[1] + 1,
    fill: props.color ? props.color : BASE_BEAD_CONFIG.fill!,
  };

  return [
    h(KonvaRect, { config: backgroundConfig }),
    h(KonvaRect, { config: beadConfig }),
  ];
};
