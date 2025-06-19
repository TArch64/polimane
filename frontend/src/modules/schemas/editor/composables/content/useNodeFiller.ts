import { computed, type Ref } from 'vue';
import Konva from 'konva';
import type { KonvaNodeRef } from './useNodeRef';
import { useNodeClientRect } from './useNodeClientRect';

interface IMinSize {
  width?: number;
  height?: number;
}

interface IPadding {
  vertical?: number;
  horizontal?: number;
}

export interface INodeFillerOptions {
  minSize?: IMinSize;
  padding?: IPadding;
}

export function useNodeFiller(sourceRef: KonvaNodeRef, options: INodeFillerOptions = {}): Ref<Pick<Konva.NodeConfig, 'width' | 'height'>> {
  const sourceRect = useNodeClientRect(sourceRef);

  const minSize = computed(() => ({
    width: options.minSize?.width ?? 0,
    height: options.minSize?.height ?? 0,
  }));

  const padding = computed(() => ({
    vertical: options.padding?.vertical ?? 0,
    horizontal: options.padding?.horizontal ?? 0,
  }));

  return computed(() => ({
    width: Math.max(sourceRect.value.width, minSize.value.width) + padding.value.horizontal,
    height: Math.max(sourceRect.value.height, minSize.value.height) + padding.value.vertical,
  }));
}
