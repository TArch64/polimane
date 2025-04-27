import {
  type Component,
  Fragment,
  type FunctionalComponent,
  nextTick,
  type Slot,
  type VNode,
  type VNodeTypes,
} from 'vue';

interface ISlots {
  default: Slot;
}

function isComponent(node: VNodeTypes): node is Component {
  if (typeof node !== 'object') return false;
  return 'name' in node || '__name' in node;
}

function getContentNodes(nodes: VNode[]): VNode[] {
  return nodes.flatMap((child) => {
    if (child.type === Fragment && child.children?.length) {
      return getContentNodes(child.children as VNode[]);
    }
    return isComponent(child.type) ? child : [];
  });
}

export type CanvasStackRendererUpdate = (isInitial: boolean, keys: unknown[]) => void;

export function createCanvasStackRenderer(update: CanvasStackRendererUpdate): FunctionalComponent<{}, {}, ISlots> {
  let lastKeys: unknown[] = [];

  return (_, ctx) => {
    const children = ctx.slots.default();
    const nodes = getContentNodes(children);
    const keys = nodes.map((node, index) => node.key ?? index);

    if (lastKeys.join('-') !== keys.join('-')) {
      const isInitial = !lastKeys.length;
      nextTick(() => update(isInitial, keys));
      lastKeys = keys;
    }

    return children;
  };
}
