import { type AsyncComponentLoader, defineAsyncComponent, type FunctionalComponent, h } from 'vue';

export interface IIconProps {
  size?: number | string | 'inline';
}

export type IconComponent = FunctionalComponent<IIconProps>;

const makeIcon = (loader: AsyncComponentLoader): IconComponent => {
  const asyncComponent = defineAsyncComponent(loader);

  function makeSize(size?: number | string) {
    if (!size) return '20px';
    if (size === 'inline') return '1em';
    return `${size}px`;
  }

  return (props, ctx) => {
    const size = makeSize(props.size);

    return h('span', {
      ...ctx.attrs,

      style: {
        display: props.size === 'inline' ? 'inline-flex' : 'flex',
        width: size,
        height: size,
      },
    }, [
      h(asyncComponent, {
        xmlns: 'http://www.w3.org/2000/svg',
        width: '100%',
        height: '100%',
      }),
    ]);
  };
};

export const SettingsIcon = makeIcon(() => import('~icons/eva/settings-outline'));
export const CloseIcon = makeIcon(() => import('~icons/eva/close-outline'));
export const ArrowBackIcon = makeIcon(() => import('~icons/eva/arrow-back-outline'));
export const TrashIcon = makeIcon(() => import('~icons/eva/trash-outline'));
export const PlusIcon = makeIcon(() => import('~icons/eva/plus-outline'));
export const EditIcon = makeIcon(() => import('~icons/eva/edit-outline'));
export const MoreHorizontalIcon = makeIcon(() => import('~icons/eva/more-horizontal-outline'));
export const CopyIcon = makeIcon(() => import('~icons/eva/copy-outline'));
export const CornerUpLeftIcon = makeIcon(() => import('~icons/eva/corner-up-left-outline'));
export const CornerUpRightIcon = makeIcon(() => import('~icons/eva/corner-up-right-outline'));
export const LogOutIcon = makeIcon(() => import('~icons/eva/log-out-outline'));
export const SaveIcon = makeIcon(() => import('~icons/eva/save-outline'));
export const CheckmarkCircleIcon = makeIcon(() => import('~icons/eva/checkmark-circle-outline'));
export const CheckmarkIcon = makeIcon(() => import('~icons/eva/checkmark-outline'));
export const LoaderIcon = makeIcon(() => import('~icons/eva/loader-outline'));
export const UnlockIcon = makeIcon(() => import('~icons/eva/unlock-outline'));
export const RepeatIcon = makeIcon(() => import('~icons/eva/repeat-outline'));
export const FileTextIcon = makeIcon(() => import('~icons/eva/file-text-outline'));
export const DropletOffIcon = makeIcon(() => import('~icons/eva/droplet-off-outline'));
export const PersonIcon = makeIcon(() => import('~icons/eva/person-outline'));
export const PeopleIcon = makeIcon(() => import('~icons/eva/people-outline'));
export const MoveIcon = makeIcon(() => import('~icons/eva/move-outline'));
export const MaximizeIcon = makeIcon(() => import('~icons/eva/maximize-outline'));

export const PersonFillIcon = makeIcon(() => import('~icons/eva/person-fill'));

export const LogoIcon = makeIcon(() => import('~icons/custom/logo'));
export const ToolsBeadCircleIcon = makeIcon(() => import('~icons/custom/tools-bead-circle'));
export const ToolsBeadBugleIcon = makeIcon(() => import('~icons/custom/tools-bead-bugle'));
export const ToolsSelectionIcon = makeIcon(() => import('~icons/custom/tools-selection'));
