import { onMounted, onUnmounted, ref, type Ref, unref, watch } from 'vue';
import { autoScrollForElements } from '@atlaskit/pragmatic-drag-and-drop-auto-scroll/element';
import type { CleanupFn, ElementDragType } from '@atlaskit/pragmatic-drag-and-drop/types';
import type {
  ElementAutoScrollArgs,
} from '@atlaskit/pragmatic-drag-and-drop-auto-scroll/dist/types/internal-types';
import type { WrapMaybeRef } from '@/types';

export type AutoScrollArgs = WrapMaybeRef<ElementAutoScrollArgs<ElementDragType>, 'element'> & {
  isEnabled?: Ref<boolean>;
};

export function useAutoScroll(args: AutoScrollArgs): void {
  const isEnabled = args.isEnabled ?? ref(true);
  let cleanup: CleanupFn;

  function mount() {
    cleanup = autoScrollForElements({
      ...args,
      element: unref(args.element),
    });
  }

  watch(isEnabled, (value) => {
    value ? mount() : cleanup?.();
  });

  onMounted(() => isEnabled.value && mount());
  onUnmounted(() => cleanup?.());
}
