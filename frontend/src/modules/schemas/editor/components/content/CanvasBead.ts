import { type FunctionalComponent, h } from 'vue';
import Konva from 'konva';
import { getThemeVar } from '@/composables';
import { SCREENSHOT_IGNORE } from '../../composables';
import { BEAD_SIZE, type BeadPosition } from './useBeadsGrid';
import { GroupRenderer } from './base';
import { KonvaRect } from './konva';

export interface ICanvasBeadProps {
  offset: BeadPosition;
  color: string | null;
}

export const CanvasBead: FunctionalComponent<ICanvasBeadProps> = (props) => {
  const groupConfig: Partial<Konva.GroupConfig> = {
    name: props.color ? undefined : SCREENSHOT_IGNORE,
    x: props.offset[0],
    y: props.offset[1],
  };

  const backgroundConfig: Partial<Konva.RectConfig> = {
    width: BEAD_SIZE,
    height: BEAD_SIZE,
  };

  const contentConfig: Partial<Konva.RectConfig> = {
    x: 1,
    y: 1,
    width: BEAD_SIZE - 2,
    height: BEAD_SIZE - 2,
    cornerRadius: getThemeVar('--rounded-full'),
    fill: props.color ?? getThemeVar('--color-background-3'),
  };

  return h(GroupRenderer, { config: groupConfig }, () => [
    h(KonvaRect, { config: backgroundConfig }),
    h(KonvaRect, { config: contentConfig }),
  ]);
};
