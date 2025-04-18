import { defineStore } from 'pinia';
import { computed } from 'vue';
import { getPatternTitle, PatternType } from '@/enums';
import { newId } from '@/helpers';
import { useEditorStore } from './editorStore';

export const usePatternsStore = defineStore('schemas/editor/patterns', () => {
  const editorStore = useEditorStore();

  const patterns = computed(() => editorStore.schema.content.patterns);
  const hasPatterns = computed(() => !!patterns.value.length);

  function addPattern(type: PatternType): void {
    const { patterns } = editorStore.schema.content;

    patterns.push({
      id: newId(),
      name: `${getPatternTitle(type)} [${patterns.length + 1}]`,
      type,
    });
  }

  return { patterns, hasPatterns, addPattern };
});
