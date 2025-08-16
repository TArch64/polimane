/// <reference types="vite/client" />
/// <reference types="unplugin-icons/types/vue" />

import Konva from 'konva';
import { ShallowRef } from 'vue';

interface ImportMetaEnv {
  FRONTEND_PUBLIC_API_URL: string;
  FRONTEND_PUBLIC_SENTRY_DSN: string;
  FRONTEND_PUBLIC_CDN_HOST: string;
}

interface ImportMeta {
  readonly env: ImportMetaEnv;
}

declare global {
  // This is a workaround for the fact that `setTimeout` returns a different type in Node.js and browser environments.
  type TimeoutId = ReturnType<typeof setTimeout>;

  interface Window {
    __KONVA_STAGE_REF__: ShallowRef<Konva.Stage | null>;
  }
}

declare module 'konva/lib/Node' {
  export interface NodeEventMap extends GlobalEventHandlersEventMap {
  }
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

  export interface IKonvaNodeHolder<N extends Konva.Node = Konva.Node> {
    getNode: () => N;
  }

  interface KonvaComponentMethods<C> extends IKonvaNodeHolder<InstanceType<C>> {
    getStage: () => Konva.Stage;
  }

  type KonvaEvent<C, E> = KonvaEventObject<E, InstanceType<C>>;

  interface KonvaComponentEmits<C> {
    mouseover: (event: KonvaEvent<C, MouseEvent>) => void;
    mouseout: (event: KonvaEvent<C, MouseEvent>) => void;
    mousedown: (event: KonvaEvent<C, MouseEvent>) => void;
    mouseup: (event: KonvaEvent<C, MouseEvent>) => void;
    mousemove: (event: KonvaEvent<C, MouseEvent>) => void;
    contextmenu: (event: KonvaEvent<C, MouseEvent>) => void;
    click: (event: KonvaEvent<C, MouseEvent>) => void;
    wheel: (event: KonvaEvent<C, WheelEvent>) => void;
  }

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
  export type KonvaLine = KonvaComponent<typeof Konva.Line>;
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
    KonvaLine: K.KonvaLine;
  }
}

export {};
