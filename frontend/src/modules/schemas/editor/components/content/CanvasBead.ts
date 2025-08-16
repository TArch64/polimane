import { type FunctionalComponent, h, resolveComponent } from 'vue';
import Konva from 'konva';
import { getThemeVar } from '@/composables';
import { SCREENSHOT_IGNORE } from '../../composables';
import { BEAD_SIZE, type BeadPosition } from './useBeadsGrid';

export interface ICanvasBeadProps {
  offset: BeadPosition;
  color: string | null;
}

export const CanvasBead: FunctionalComponent<ICanvasBeadProps> = (props) => {
  const KonvaRect = resolveComponent('KonvaRect');

  const backgroundConfig: Partial<Konva.RectConfig> = {
    x: props.offset[0],
    y: props.offset[1],
    width: BEAD_SIZE,
    height: BEAD_SIZE,
  };

  const contentConfig: Partial<Konva.RectConfig> = {
    name: props.color ? undefined : SCREENSHOT_IGNORE,
    x: props.offset[0] + 1,
    y: props.offset[1] + 1,
    width: BEAD_SIZE - 2,
    height: BEAD_SIZE - 2,
    cornerRadius: getThemeVar('--rounded-full'),
    fill: props.color ?? getThemeVar('--color-background-3'),
  };

  return [
    h(KonvaRect, { config: backgroundConfig }),
    h(KonvaRect, { config: contentConfig }),
  ];
};
