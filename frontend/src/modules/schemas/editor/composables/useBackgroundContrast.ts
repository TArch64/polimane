import { useEditorStore } from '@editor/stores';
import { type IContrast, useContrast } from './useContrast';

export function useBackgroundContrast(foreground: string): IContrast {
  const editorStore = useEditorStore();
  return useContrast(() => editorStore.schema.backgroundColor, foreground);
}
