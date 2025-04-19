import { onMounted, onUnmounted } from 'vue';
import type {
  CleanupFn,
  ElementDragType,
  MonitorArgs,
} from '@atlaskit/pragmatic-drag-and-drop/types';
import { monitorForElements } from '@atlaskit/pragmatic-drag-and-drop/element/adapter';

export function useMonitor(args: MonitorArgs<ElementDragType>): void {
  let cleanup: CleanupFn;
  onMounted(() => cleanup = monitorForElements(args));
  onUnmounted(() => cleanup());
}
