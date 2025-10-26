import { type FunctionalComponent, h } from 'vue';
import { type BeadContentKind, BeadKind } from '@/enums';
import { type IconComponent, ToolsBeadBugleIcon, ToolsBeadCircleIcon } from '@/components/icon';

const map: Record<BeadContentKind, IconComponent> = {
  [BeadKind.CIRCLE]: ToolsBeadCircleIcon,
  [BeadKind.BUGLE]: ToolsBeadBugleIcon,
};

export interface IBeadIconProps {
  kind: BeadContentKind;
}

export const BeadIcon: FunctionalComponent<IBeadIconProps> = (props, { attrs }) => {
  return h(map[props.kind], attrs);
};
