import { computed, type Ref } from 'vue';
import { useBackgroundContrast } from './useBackgroundContrast';

export function useSelectionColor(): Ref<string> {
  const contrast = useBackgroundContrast('#000');
  return computed(() => contrast.isAA ? 'var(--color-black)' : 'var(--color-white)');
}
