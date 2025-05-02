import { defineStore } from 'pinia';
import { computed } from 'vue';
import { getPatternTitle, PatternType } from '@/enums';
import { newId } from '@/helpers';
import { Collection, type ISchemaPattern } from '@/models';
import { setObjectParent } from '../models';
import { useEditorStore } from './editorStore';

export const usePatternsStore = defineStore('schemas/editor/patterns', () => {
  const editorStore = useEditorStore();

  const patterns = Collection.fromParent(editorStore.schema, {
    onAdded: (parent, object) => setObjectParent(parent, object),
  });

  const hasPatterns = computed(() => !!patterns.size);

  const addPattern = (kind: PatternType) => patterns.append({
    id: newId(),
    name: `${getPatternTitle(kind)} [${patterns.size + 1}]`,
    type: kind,
    content: [],
  });

  const deletePattern = (pattern: ISchemaPattern) => patterns.delete(pattern);

  return { patterns, hasPatterns, addPattern, deletePattern };
});
