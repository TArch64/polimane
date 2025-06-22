import {
  inject,
  type InjectionKey,
  type MaybeRefOrGetter,
  onBeforeUnmount,
  provide,
  ref,
  type Ref,
  toValue,
} from 'vue';
import Konva from 'konva';
import {
  ContextMenuPlugin,
  type IContextMenuAction,
  type MaybeContextMenuAction,
} from '@/components/contextMenu';
import { Point } from '@/models';
import { newId } from '@/helpers';
import { useNodeListener } from './useNodeListener';

export interface INodeContextMenuOptions {
  nodeRef: MaybeRefOrGetter<Konva.Node>;
  title: MaybeRefOrGetter<string>;
  actions: MaybeRefOrGetter<MaybeContextMenuAction[]>;
}

interface INodeContextMenuItem extends INodeContextMenuOptions {
  id: string;
}

const TOKEN = Symbol('nodeContextMenu') as InjectionKey<Ref<INodeContextMenuItem[]>>;

export function provideNodeContextMenu(stage: Ref<Konva.Stage>): void {
  const plugin = ContextMenuPlugin.inject();
  const items: Ref<INodeContextMenuItem[]> = ref([]);

  provide(TOKEN, items);

  function isClosestNode(parent: Konva.Stage | Konva.Node, node: Konva.Node): boolean {
    if (parent._id === node._id) {
      return true;
    }

    return node.parent ? isClosestNode(parent, node.parent) : false;
  }

  function matchOptions(node: Konva.Node): INodeContextMenuOptions | null {
    const matched = items.value
      .filter((item) => isClosestNode(toValue(item.nodeRef), node))
      .sort((i1, i2) => toValue(i2.nodeRef).getDepth() - toValue(i1.nodeRef).getDepth());

    return matched[0] || null;
  }

  useNodeListener(stage, 'contextmenu', (event: Konva.KonvaEventObject<MouseEvent>) => {
    const options = matchOptions(event.target);

    if (!options) return;

    plugin.show({
      control: false,
      title: toValue(options.title),
      actions: toValue(options.actions).filter((action): action is IContextMenuAction => !!action),

      position: new Point({
        x: event.evt.clientX,
        y: event.evt.clientY,
      }),
    });

    addEventListener('click', () => plugin.hide(), { once: true, capture: true });
  });
}

export function useNodeContextMenu(options: INodeContextMenuOptions): void {
  const item: INodeContextMenuItem = { ...options, id: newId() };
  const items = inject(TOKEN)!;
  items.value.push(item);

  onBeforeUnmount(() => {
    items.value = items.value.filter((i) => i.id !== item.id);
  });
}
