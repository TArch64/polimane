import type { VNode, VNodeArrayChildren, VNodeChild, VNodeNormalizedChildren } from 'vue';

export function getChildren(instance: VNode) {
  const isVNode = (value: VNodeChild | VNodeNormalizedChildren): value is VNode =>
    !!value?.hasOwnProperty('component');
  const isVNodeArrayChildren = (
    value: VNodeChild | VNodeNormalizedChildren,
  ): value is VNodeArrayChildren => Array.isArray(value);

  const recursivelyFindChildren = (item: VNodeChild | VNodeNormalizedChildren): VNode[] => {
    if (isVNode(item)) return [item, ...recursivelyFindChildren(item.children)];
    if (isVNodeArrayChildren(item)) return item.flatMap(recursivelyFindChildren);
    return [];
  };
  return recursivelyFindChildren(instance.children);
}
