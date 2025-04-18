import { defineStore } from 'pinia';
import { computed } from 'vue';
import { PatternType } from '@/models';
import { newId } from '@/helpers';
import { useEditorStore } from './editorStore';

export const usePatternsStore = defineStore('schemas/editor/patterns', () => {
  const editorStore = useEditorStore();

  const patterns = computed(() => editorStore.schema.content.patterns);
  const hasPatterns = computed(() => !!patterns.value.length);

  function addPattern(type: PatternType): void {
    editorStore.schema.content.patterns.push({
      id: newId(),
      type,
    });
  }

  return { patterns, hasPatterns, addPattern };
});
