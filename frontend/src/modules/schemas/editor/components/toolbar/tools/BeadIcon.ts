import { type FunctionalComponent, h } from 'vue';
import { BeadKind } from '@/enums';
import { type IconComponent, ToolsBeadBugleIcon, ToolsBeadCircleIcon } from '@/components/icon';

const map: Record<BeadKind, IconComponent> = {
  [BeadKind.CIRCLE]: ToolsBeadCircleIcon,
  [BeadKind.BUGLE]: ToolsBeadBugleIcon,
};

export interface IBeadIconProps {
  kind: BeadKind;
}

export const BeadIcon: FunctionalComponent<IBeadIconProps> = (props, { attrs }) => {
  return h(map[props.kind], attrs);
};
