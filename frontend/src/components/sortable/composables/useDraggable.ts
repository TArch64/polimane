import { onMounted, onUnmounted, unref } from 'vue';
import type { CleanupFn } from '@atlaskit/pragmatic-drag-and-drop/types';
import { draggable } from '@atlaskit/pragmatic-drag-and-drop/element/adapter';
import type { WrapMaybeRef } from '@/types';

export type DraggableArgs = WrapMaybeRef<Parameters<typeof draggable>[0], 'element' | 'dragHandle'>;

export function useDraggable(args: DraggableArgs) {
  let cleanup: CleanupFn;

  onMounted(() => cleanup = draggable({
    ...args,
    element: unref(args.element),
    dragHandle: unref(args.dragHandle),
  }));

  onUnmounted(() => cleanup());
}
