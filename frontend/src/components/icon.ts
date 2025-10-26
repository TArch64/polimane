import { type AsyncComponentLoader, defineAsyncComponent, type FunctionalComponent, h } from 'vue';

export interface IIconProps {
  size?: number | string | 'inline';
}

export type IconComponent = FunctionalComponent<IIconProps>;

const makeIcon = (name: string, loader: AsyncComponentLoader): IconComponent => {
  const asyncComponent = defineAsyncComponent(loader);

  function makeSize(size?: number | string) {
    if (!size) return '20px';
    if (size === 'inline') return '1em';
    return `${size}px`;
  }

  const Component: IconComponent = (props, ctx) => {
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

  Component.displayName = name;
  return Component;
};

export const SettingsIcon = makeIcon('SettingsIcon', () => import('~icons/eva/settings-outline'));
export const CloseIcon = makeIcon('CloseIcon', () => import('~icons/eva/close-outline'));
export const ArrowBackIcon = makeIcon('ArrowBackIcon', () => import('~icons/eva/arrow-back-outline'));
export const TrashIcon = makeIcon('TrashIcon', () => import('~icons/eva/trash-outline'));
export const PlusIcon = makeIcon('PlusIcon', () => import('~icons/eva/plus-outline'));
export const EditIcon = makeIcon('EditIcon', () => import('~icons/eva/edit-outline'));
export const MoreHorizontalIcon = makeIcon('MoreHorizontalIcon', () => import('~icons/eva/more-horizontal-outline'));
export const CopyIcon = makeIcon('CopyIcon', () => import('~icons/eva/copy-outline'));
export const CornerUpLeftIcon = makeIcon('CornerUpLeftIcon', () => import('~icons/eva/corner-up-left-outline'));
export const CornerUpRightIcon = makeIcon('CornerUpRightIcon', () => import('~icons/eva/corner-up-right-outline'));
export const LogOutIcon = makeIcon('LogOutIcon', () => import('~icons/eva/log-out-outline'));
export const SaveIcon = makeIcon('SaveIcon', () => import('~icons/eva/save-outline'));
export const CheckmarkCircleIcon = makeIcon('CheckmarkCircleIcon', () => import('~icons/eva/checkmark-circle-outline'));
export const CheckmarkIcon = makeIcon('CheckmarkIcon', () => import('~icons/eva/checkmark-outline'));
export const LoaderIcon = makeIcon('LoaderIcon', () => import('~icons/eva/loader-outline'));
export const UnlockIcon = makeIcon('UnlockIcon', () => import('~icons/eva/unlock-outline'));
export const RepeatIcon = makeIcon('RepeatIcon', () => import('~icons/eva/repeat-outline'));
export const FileTextIcon = makeIcon('FileTextIcon', () => import('~icons/eva/file-text-outline'));
export const DropletOffIcon = makeIcon('DropletOffIcon', () => import('~icons/eva/droplet-off-outline'));
export const PersonIcon = makeIcon('PeopleIcon', () => import('~icons/eva/person-outline'));
export const PeopleIcon = makeIcon('PeopleIcon', () => import('~icons/eva/people-outline'));

export const PersonFillIcon = makeIcon('PersonFillIcon', () => import('~icons/eva/person-fill'));

export const LogoIcon = makeIcon('LogoIcon', () => import('~icons/custom/logo'));
export const ToolsBeadCircleIcon = makeIcon('ToolsBeadCircleIcon', () => import('~icons/custom/tools-bead-circle'));
export const ToolsBeadBugleIcon = makeIcon('ToolsBeadBugleIcon', () => import('~icons/custom/tools-bead-bugle'));
export const ToolsSelectionIcon = makeIcon('ToolsSelectionIcon', () => import('~icons/custom/tools-selection'));
