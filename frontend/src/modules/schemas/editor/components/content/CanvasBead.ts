import { type FunctionalComponent, h, resolveComponent } from 'vue';
import Konva from 'konva';
import { getThemeVar } from '@/composables';
import { BEAD_SIZE, type BeadPosition } from './useBeadsGrid';
import { GroupRenderer } from './base';

export interface ICanvasBeadProps {
  position: BeadPosition;
  offset: BeadPosition;
}

export const CanvasBead: FunctionalComponent<ICanvasBeadProps> = (props) => {
  const KonvaRect = resolveComponent('KonvaRect');

  const groupConfig: Partial<Konva.GroupConfig> = {
    x: props.offset[0],
    y: props.offset[1],
  };

  const backgroundConfig: Partial<Konva.RectConfig> = {
    width: BEAD_SIZE,
    height: BEAD_SIZE,
    // fill: getThemeVar('--color-background-1'),
  };

  const contentConfig: Partial<Konva.RectConfig> = {
    x: 1,
    y: 1,
    width: BEAD_SIZE - 2,
    height: BEAD_SIZE - 2,
    cornerRadius: getThemeVar('--rounded-full'),
    fill: getThemeVar('--color-background-3'),
  };

  return h(GroupRenderer, { config: groupConfig }, () => [
    h(KonvaRect, { config: backgroundConfig }),
    h(KonvaRect, { config: contentConfig }),
  ]);
};
