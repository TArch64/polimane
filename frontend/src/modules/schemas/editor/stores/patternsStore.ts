import { defineStore } from 'pinia';
import { computed } from 'vue';
import { getPatternTitle, PatternType } from '@/enums';
import { newId } from '@/helpers';
import { Collection, type ISchemaPattern } from '@/models';
import { setObjectParent } from '../models';
import { useEditorStore } from './editorStore';

export const usePatternsStore = defineStore('schemas/editor/patterns', () => {
  const editorStore = useEditorStore();
  const patterns = Collection.fromParent(editorStore.schema);
  const hasPatterns = computed(() => !!patterns.size);

  function createPattern(type: PatternType) {
    const pattern: ISchemaPattern = {
      id: newId(),
      name: `${getPatternTitle(type)} [${patterns.size + 1}]`,
      type: type,
      content: [],
    };

    setObjectParent(editorStore.schema, pattern);
    return pattern;
  }

  function deletePattern(pattern: ISchemaPattern) {
    patterns.delete(pattern);
  }

  function movePattern(pattern: ISchemaPattern, shift: number): void {
    const index = patterns.indexOf(pattern);
    const newIndex = index + shift;

    if (newIndex < 0 || newIndex === patterns.size) {
      return;
    }

    patterns.move(pattern, newIndex);
  }

  return { patterns, hasPatterns, createPattern, deletePattern, movePattern };
});
