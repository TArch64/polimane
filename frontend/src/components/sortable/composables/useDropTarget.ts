import { onMounted, onUnmounted, unref } from 'vue';
import type { CleanupFn } from '@atlaskit/pragmatic-drag-and-drop/types';
import { dropTargetForElements } from '@atlaskit/pragmatic-drag-and-drop/element/adapter';
import type { WrapMaybeRef } from '@/types';

export type DropTargetArgs = WrapMaybeRef<Parameters<typeof dropTargetForElements>[0], 'element'>;

export function useDropTarget(args: DropTargetArgs) {
  let cleanup: CleanupFn;

  onMounted(() => cleanup = dropTargetForElements({
    ...args,
    element: unref(args.element),
  }));

  onUnmounted(() => cleanup());
}
