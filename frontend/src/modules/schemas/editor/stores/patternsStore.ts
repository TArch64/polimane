import { defineStore } from 'pinia';
import { computed } from 'vue';
import { getPatternTitle, PatternType } from '@/enums';
import { newId } from '@/helpers';
import { Collection, type ISchemaPattern } from '@/models';
import { setObjectParent } from '../models';
import { useEditorStore } from './editorStore';

export interface IPatternAddOptions {
  kind: PatternType;
  toIndex: number;
}

export const usePatternsStore = defineStore('schemas/editor/patterns', () => {
  const editorStore = useEditorStore();

  const patterns = Collection.fromParent(editorStore.schema, {
    onAdded: (parent, object) => setObjectParent(parent, object),
  });

  const hasPatterns = computed(() => !!patterns.size);

  const addPattern = (options: IPatternAddOptions) => patterns.insert({
    id: newId(),
    name: `${getPatternTitle(options.kind)} [${patterns.size + 1}]`,
    type: options.kind,
    content: [],
  }, {
    toIndex: options.toIndex,
  });

  function deletePattern(pattern: ISchemaPattern) {
    patterns.delete(pattern);
  }

  return { patterns, hasPatterns, addPattern, deletePattern };
});
