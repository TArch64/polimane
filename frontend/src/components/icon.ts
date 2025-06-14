import { type AsyncComponentLoader, defineAsyncComponent, type FunctionalComponent, h } from 'vue';

export interface IIconProps {
  size?: number | string;
}

export type IconComponent = FunctionalComponent<IIconProps>;

const makeIcon = (loader: AsyncComponentLoader): IconComponent => {
  const asyncComponent = defineAsyncComponent(loader);

  return (props) => h('span', {
    style: {
      display: 'flex',
      width: `${props.size ?? 20}px`,
      height: `${props.size ?? 20}px`,
    },
  }, [
    h(asyncComponent, {
      width: '100%',
      height: '100%',
    }),
  ]);
};

export const SettingsIcon = makeIcon(() => import('~icons/eva/settings-outline'));
export const CloseIcon = makeIcon(() => import('~icons/eva/close-outline'));
export const ArrowBackIcon = makeIcon(() => import('~icons/eva/arrow-back-outline'));
export const ArrowUpwardIcon = makeIcon(() => import('~icons/eva/arrow-upward-outline'));
export const ArrowDownwardIcon = makeIcon(() => import('~icons/eva/arrow-downward-outline'));
export const TrashIcon = makeIcon(() => import('~icons/eva/trash-outline'));
export const PlusIcon = makeIcon(() => import('~icons/eva/plus-outline'));
export const EditIcon = makeIcon(() => import('~icons/eva/edit-outline'));
export const MoreHorizontalIcon = makeIcon(() => import('~icons/eva/more-horizontal-outline'));
export const CopyIcon = makeIcon(() => import('~icons/eva/copy-outline'));
