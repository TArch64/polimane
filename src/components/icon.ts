import { type AsyncComponentLoader, defineAsyncComponent, type FunctionalComponent, h } from 'vue';

export interface IIconProps {
  size?: number;
}

const makeIcon = (loader: AsyncComponentLoader): FunctionalComponent<IIconProps> => {
  const asyncComponent = defineAsyncComponent(loader);

  return (props) => h(asyncComponent, {
    width: `${props.size ?? 20}px`,
    height: `${props.size ?? 20}px`,
  });
};

export const SettingsIcon = makeIcon(() => import(`~icons/eva/settings-outline`));
