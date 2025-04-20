import { defineStore } from 'pinia';
import { computed } from 'vue';
import { getPatternTitle, PatternKind, SchemaObjectType } from '@/enums';
import { newId } from '@/helpers';
import { Collection, type ISchemaPattern } from '@/models';
import { useEditorStore } from './editorStore';

export const usePatternsStore = defineStore('schemas/editor/patterns', () => {
  const editorStore = useEditorStore();

  const patterns = Collection.fromProperty(editorStore.schema, 'content');
  const hasPatterns = computed(() => !!patterns.size);

  const addPattern = (kind: PatternKind) => patterns.append({
    id: newId(),
    type: SchemaObjectType.PATTERN,
    name: `${getPatternTitle(kind)} [${patterns.size + 1}]`,
    kind: kind,
    content: [],
  });

  const deletePattern = (pattern: ISchemaPattern) => patterns.delete(pattern);

  return { patterns, hasPatterns, addPattern, deletePattern };
});
