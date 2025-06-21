import {
  Comment,
  computed,
  defineComponent,
  getCurrentInstance,
  h,
  type PropType,
  resolveComponent,
  type Slot,
  type SlotsType,
  type VNode,
  type VNodeProps,
} from 'vue';
import Konva from 'konva';
import type { IKonvaNodeHolder, KonvaGroup } from 'vue-konva';
import { useDebounceFn } from '@vueuse/core';
import type { InferComponentProps, MaybeArray } from '@/types';
import { newId } from '@/helpers';
import { useNodeRef } from '@/modules/schemas/editor/composables';
import { NodeRect } from '@/models';
import { getClientRect } from '@/modules/schemas/editor/helpers';

export interface IGroupLayoutEvent {
  clientRect: NodeRect;
  nodes: Konva.Node[];
}

type GroupProps = InferComponentProps<KonvaGroup> & VNodeProps & Record<string, unknown>;

const PROXY_PROPS = [
  'onMouseout',
  'onMouseover',
  'onClick',
];

export const GroupRenderer = defineComponent({
  name: 'GroupRenderer',

  props: {
    config: {
      type: Object as PropType<Partial<Konva.GroupConfig>>,
      required: false,
    },
  },

  emits: [
    'click',
    'mouseover',
    'mouseout',
    'layout',
  ],

  slots: Object as SlotsType<{
    default: Slot;
  }>,

  setup(props, ctx) {
    const KonvaGroup = resolveComponent('KonvaGroup');

    const instance = getCurrentInstance();
    const groupRef = useNodeRef<Konva.Group>();
    const groupId = props.config?.id ?? newId();
    let nodes: Konva.Node[] = [];

    const config = computed(() => ({
      ...props.config,
      id: groupId,
    }));

    function isKonvaComponent(node: VNode): node is VNode & {
      component: { exposed: IKonvaNodeHolder };
    } {
      return typeof node.component?.exposed?.getNode === 'function';
    }

    function isDebugNode(node: VNode): boolean {
      return !!node.component?.exposed?.debugNode;
    }

    function getContentNodes(nodes: VNode[]): Konva.Node[] {
      return nodes.flatMap((child): MaybeArray<Konva.Node> => {
        if (child.type === Comment || isDebugNode(child)) {
          return [];
        }
        if (isKonvaComponent(child)) {
          return child.component.exposed.getNode();
        }
        if (child.component && isKonvaComponent(child.component.subTree)) {
          return getContentNodes([child.component.subTree]);
        }
        return child.children?.length ? getContentNodes(child.children as VNode[]) : [];
      });
    }

    function getChildrenClientRect(): NodeRect {
      if (!nodes.length) {
        return NodeRect.BLANK;
      }

      let minX = Number.POSITIVE_INFINITY;
      let minY = Number.POSITIVE_INFINITY;
      let maxX = Number.NEGATIVE_INFINITY;
      let maxY = Number.NEGATIVE_INFINITY;

      for (const node of nodes) {
        const { height, width, x, y } = getClientRect(node);

        minX = Math.min(minX, x);
        minY = Math.min(minY, y);
        maxX = Math.max(maxX, x + width);
        maxY = Math.max(maxY, y + height);
      }

      return new NodeRect({
        x: minX,
        y: minY,
        width: maxX - minX,
        height: maxY - minY,
      });
    }

    const updateSize = useDebounceFn(() => {
      if (!groupRef.value.listening()) {
        return;
      }

      let clientRect = getChildrenClientRect();

      if (config.value.width) {
        clientRect = clientRect.with({ width: config.value.width });
      }

      if (config.value.height) {
        clientRect = clientRect.with({ height: config.value.height });
      }

      groupRef.value.width(clientRect.width);
      groupRef.value.height(clientRect.height);

      const event: IGroupLayoutEvent = {
        clientRect,
        nodes,
      };

      ctx.emit('layout', event);
      groupRef.value.fire('layout', { evt: event }, true);
    }, 5);

    function syncChildNames(): void {
      const groupName = `managed-by-${groupId}`;

      for (const node of nodes) {
        if (!node.id()) {
          node.id(newId());
        }

        if (!node.hasName(groupName)) {
          node.addName(groupName);
          node.on('xChange', updateSize);
          node.on('yChange', updateSize);
          node.on('widthChange', updateSize);
          node.on('heightChange', updateSize);
          node.on('layout', updateSize);
        }
      }
    }

    function getChildVNodes(vnode: VNode) {
      return vnode.component!.subTree.children as VNode[];
    }

    async function onMounted(vnode: VNode) {
      nodes = getContentNodes(getChildVNodes(vnode));
      syncChildNames();
      await updateSize();
    }

    async function onUpdated(vnode: VNode) {
      nodes = getContentNodes(getChildVNodes(vnode));
      syncChildNames();
      await updateSize();
    }

    ctx.expose({
      getNode: () => groupRef.value,
    });

    return () => {
      const groupProps: GroupProps = {
        ref: groupRef,
        config: config.value,
        onVnodeMounted: onMounted,
        onVnodeUpdated: onUpdated,
      };

      for (const proxyProp of PROXY_PROPS) {
        const prop = instance?.vnode.props?.[proxyProp];
        if (prop) groupProps[proxyProp] = prop;
      }

      return h(KonvaGroup, groupProps, ctx.slots.default);
    };
  },
});

declare module 'konva/lib/Node' {
  export interface NodeEventMap {
    layout: IGroupLayoutEvent;
  }
}
