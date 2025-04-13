import { defineStore } from 'pinia';
import { useEditorStore } from './editorStore';

export const usePatternsStore = defineStore('schemas/editor/patterns', () => {
  const editorStore = useEditorStore();

  return { hasPatterns: false };
});
