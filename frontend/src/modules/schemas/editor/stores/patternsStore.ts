import { defineStore } from 'pinia';
import { computed } from 'vue';
import { getPatternTitle, PatternType } from '@/enums';
import { newId } from '@/helpers';
import { Collection, type ISchemaPattern } from '@/models';
import { useEditorStore } from './editorStore';

export const usePatternsStore = defineStore('schemas/editor/patterns', () => {
  const editorStore = useEditorStore();

  const patterns = Collection.fromProperty(editorStore.schema.content, 'patterns');
  const hasPatterns = computed(() => !!patterns.size);

  const addPattern = (type: PatternType) => patterns.append({
    id: newId(),
    name: `${getPatternTitle(type)} [${patterns.size + 1}]`,
    type,
  });

  function deletePattern(pattern: ISchemaPattern) {
    patterns.delete(pattern);
  }

  return { patterns, hasPatterns, addPattern, deletePattern };
});
