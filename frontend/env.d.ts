/// <reference types="vite/client" />
/// <reference types="unplugin-icons/types/vue" />

interface ImportMetaEnv {
  FRONTEND_PUBLIC_API_URL: string;
}

interface ImportMeta {
  readonly env: ImportMetaEnv;
}

declare global {
  // This is a workaround for the fact that `setTimeout` returns a different type in Node.js and browser environments.
  type TimeoutId = ReturnType<typeof setTimeout>;
}

declare module 'vue-konva' {
  import Konva from 'konva';
  import { KonvaEventObject } from 'konva/lib/Node';
  import { ComponentOptionsMixin, ComputedOptions, DefineComponent, PropType } from 'vue';

  type KonvaComponentProps<C> = {
    config: {
      type: PropType<ConstructorParameters<C>[0]>;
      required: false;
    };
  };

  type KonvaComponentMethods<C> = {
    getNode(): InstanceType<C>;
    getStage(): Konva.Stage;
  };

  type KonvaEvent<C, E> = KonvaEventObject<E, InstanceType<C>>;

  type KonvaComponentEmits<C> = {
    mouseover: (event: KonvaEvent<C, MouseEvent>) => void;
    mouseout: (event: KonvaEvent<C, MouseEvent>) => void;
    click: (event: KonvaEvent<C, MouseEvent>) => void;
    wheel: (event: KonvaEvent<C, WheelEvent>) => void;
  };

  export type KonvaComponent<C> = DefineComponent<
    KonvaComponentProps<C>,
    object,
    object,
    ComputedOptions,
    KonvaComponentMethods<C>,
    ComponentOptionsMixin,
    ComponentOptionsMixin,
    KonvaComponentEmits<C>
  >;

  export type KonvaStage = KonvaComponent<typeof Konva.Stage>;
  export type KonvaLayer = KonvaComponent<typeof Konva.Layer>;
  export type KonvaGroup = KonvaComponent<typeof Konva.Group>;
  export type KonvaText = KonvaComponent<typeof Konva.Text>;
  export type KonvaRect = KonvaComponent<typeof Konva.Rect>;
  export type KonvaLabel = KonvaComponent<typeof Konva.Label>;
  export type KonvaTag = KonvaComponent<typeof Konva.Tag>;
  export type KonvaImage = KonvaComponent<typeof Konva.Image>;
  export { KonvaEventObject };
}

declare module 'vue' {
  import * as K from 'vue-konva';

  export interface GlobalComponents {
    KonvaStage: K.KonvaStage;
    KonvaLayer: K.KonvaLayer;
    KonvaGroup: K.KonvaGroup;
    KonvaText: K.KonvaText;
    KonvaRect: K.KonvaRect;
    KonvaLabel: K.KonvaLabel;
    KonvaTag: K.KonvaTag;
    KonvaImage: K.KonvaImage;
  }
}

export {};
