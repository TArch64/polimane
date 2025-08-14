import {
  type ComponentOptionsMixin,
  type ComputedOptions,
  defineComponent,
  type DefineComponent,
  getCurrentInstance,
  onMounted,
  onUnmounted,
  onUpdated,
  type PropType,
  reactive,
  type VNode,
  watch,
} from 'vue';
import Konva from 'konva';
import type { KonvaEventObject } from 'konva/lib/Node';
import type { SafeAny } from '@/types';
import { applyNodeProps } from './applyNodeProps';
import { updatePicture } from './updatePicture';
import { findParentKonva } from './findParentKonva';
import { checkOrder } from './checkOrder';
import type { KonvaNodeConstructor } from './KonvaNodeConstructor';

const EVENTS_NAMESPACE = '.vue-konva-event';

type KonvaComponentProps<C extends KonvaNodeConstructor> = {
  config: {
    type: PropType<ConstructorParameters<C>[0]>;
    required: false;
  };
};

export interface IKonvaNodeHolder<N extends Konva.Node = Konva.Node> {
  getNode: () => N;
}

type KonvaComponentMethods<C extends KonvaNodeConstructor> = IKonvaNodeHolder<InstanceType<C>> & {
  getStage: () => Konva.Stage;
  [key: string]: SafeAny;
};

type KonvaEvent<C extends KonvaNodeConstructor, E> = KonvaEventObject<E, InstanceType<C>>;

type KonvaComponentEmits<C extends KonvaNodeConstructor> = {
  mouseover: (event: KonvaEvent<C, MouseEvent>) => void;
  mouseout: (event: KonvaEvent<C, MouseEvent>) => void;
  mousedown: (event: KonvaEvent<C, MouseEvent>) => void;
  mouseup: (event: KonvaEvent<C, MouseEvent>) => void;
  mousemove: (event: KonvaEvent<C, MouseEvent>) => void;
  contextmenu: (event: KonvaEvent<C, MouseEvent>) => void;
  click: (event: KonvaEvent<C, MouseEvent>) => void;
  wheel: (event: KonvaEvent<C, WheelEvent>) => void;
  [key: string]: SafeAny;
};

export type KonvaComponent<C extends KonvaNodeConstructor> = DefineComponent<
  KonvaComponentProps<C>,
  object,
  object,
  ComputedOptions,
  KonvaComponentMethods<C>,
  ComponentOptionsMixin,
  ComponentOptionsMixin,
  KonvaComponentEmits<C>
>;

export function createKonvaNode<NC extends KonvaNodeConstructor>(componentName: string, NodeConstructor: NC) {
  return defineComponent({
    name: componentName,

    props: {
      config: {
        type: Object as PropType<ConstructorParameters<NC>[number]>,
        default: () => ({}),
      },
    },

    setup(props, { attrs, slots, expose }) {
      const instance = getCurrentInstance();
      if (!instance) return;
      const oldProps = reactive({});

      const __konvaNode = new NodeConstructor({});
      instance.__konvaNode = __konvaNode;
      // @ts-expect-error untyped Konva instance
      instance.vnode.__konvaNode = __konvaNode;
      uploadKonva();

      function getNode() {
        return instance?.__konvaNode;
      }

      function getStage() {
        return instance?.__konvaNode;
      }

      function uploadKonva() {
        if (!instance) return;
        const events: VNode['props'] = {};
        for (const key in instance?.vnode.props) {
          if (key.slice(0, 2) === 'on') {
            events[key] = instance.vnode.props[key];
          }
        }
        const existingProps = oldProps || {};
        const newProps = {
          ...attrs,
          ...props.config,
          ...events,
        };
        applyNodeProps(instance, newProps, existingProps);
        Object.assign(oldProps, newProps);
      }

      onMounted(() => {
        const parentKonvaNode = findParentKonva(instance)?.__konvaNode;
        if (parentKonvaNode && 'add' in parentKonvaNode)
          (parentKonvaNode as { add: (node: Konva.Node) => void }).add(__konvaNode);
        updatePicture(__konvaNode);
      });

      onUnmounted(() => {
        updatePicture(__konvaNode);
        __konvaNode.destroy();
        __konvaNode.off(EVENTS_NAMESPACE);
      });

      onUpdated(() => {
        uploadKonva();
        checkOrder(instance.subTree, __konvaNode);
      });

      watch(() => props.config, uploadKonva, { deep: true });

      expose({
        getStage,
        getNode,
      });

      return () => slots.default?.();
    },
  }) as KonvaComponent<NC>;
}

declare module '@vue/runtime-core' {
  export interface ComponentInternalInstance {
    __konvaNode: Konva.Node;
  }
}
