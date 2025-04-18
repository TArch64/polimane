import { type AsyncComponentLoader, defineAsyncComponent, type FunctionalComponent, h } from 'vue';

export interface IIconProps {
  size?: number | string;
}

const makeIcon = (loader: AsyncComponentLoader): FunctionalComponent<IIconProps> => {
  const asyncComponent = defineAsyncComponent(loader);

  return (props) => h(asyncComponent, {
    width: `${props.size ?? 20}px`,
    height: `${props.size ?? 20}px`,
  });
};

export const SettingsIcon = makeIcon(() => import(`~icons/eva/settings-outline`));
export const CloseIcon = makeIcon(() => import(`~icons/eva/close-outline`));
export const ArrowBackIcon = makeIcon(() => import(`~icons/eva/arrow-back-outline`));
export const TrashIcon = makeIcon(() => import(`~icons/eva/trash-outline`));
export const PlusIcon = makeIcon(() => import(`~icons/eva/plus-outline`));
