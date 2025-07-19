import Konva from 'konva';
import { computed, type Ref } from 'vue';
import type { KonvaNodeRef } from './useNodeRef';
import { useNodeClientRect } from './useNodeClientRect';
import { type NodePaddingInput, useNodePadding } from './useNodePadding';

interface IMinSize {
  width?: number;
  height?: number;
}

export interface INodeFillerOptions {
  minSize?: IMinSize;
  padding?: NodePaddingInput;
}

export function useNodeFiller(sourceRef: KonvaNodeRef, options: INodeFillerOptions = {}): Ref<Pick<Konva.NodeConfig, 'width' | 'height'>> {
  const sourceRect = useNodeClientRect(sourceRef);
  const padding = useNodePadding(options.padding ?? 0);

  const minSize = computed(() => ({
    width: options.minSize?.width ?? 0,
    height: options.minSize?.height ?? 0,
  }));

  return computed(() => ({
    width: Math.max(sourceRect.value.width, minSize.value.width) + padding.horizontal * 2,
    height: Math.max(sourceRect.value.height, minSize.value.height) + padding.vertical * 2,
  }));
}
