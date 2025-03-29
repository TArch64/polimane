import type { AllowedComponentProps, VNodeProps } from 'vue';
import type { ComponentProps as ComponentProps_ } from 'vue-component-type-helpers';

export type ComponentProps<T> = Omit<ComponentProps_<T>, keyof (VNodeProps & AllowedComponentProps)>;
