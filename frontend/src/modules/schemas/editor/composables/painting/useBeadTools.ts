import { useToolsStore } from '@editor/stores';
import { computed, type Ref } from 'vue';
import { type IBeadPaintingListeners, useBeadPainting } from './useBeadPainting';
import { type IBeadSelectionListeners, useBeadSelection } from './useBeadSelection';
import type { IBeadToolsOptions } from './IBeadToolsOptions';

type BeadToolListeners = IBeadPaintingListeners & IBeadSelectionListeners;

export function useBeadTools(options: IBeadToolsOptions): Ref<BeadToolListeners> {
  const toolsStore = useToolsStore();
  const painting = useBeadPainting(options);
  const selection = useBeadSelection(options);

  return computed(() => {
    return toolsStore.isSelection ? selection.value : painting.value;
  });
}
