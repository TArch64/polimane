import { computed, type Ref } from 'vue';
import { useBackgroundContrast } from './useBackgroundContrast';

const EMPTY_LIGHT = 'rgba(0, 0, 0, 0.1)';
const EMPTY_DARK = 'rgba(255, 255, 255, 0.1)';

export function useBackgroundCanvasColor(): Ref<string> {
  const contrast = useBackgroundContrast(EMPTY_DARK);
  return computed(() => contrast.isAA ? EMPTY_DARK : EMPTY_LIGHT);
}
