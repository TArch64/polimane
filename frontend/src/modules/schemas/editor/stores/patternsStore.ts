import { defineStore } from 'pinia';
import { computed } from 'vue';
import { ulid } from 'ulid';
import { PatternType } from '@/models';
import { useEditorStore } from './editorStore';

export const usePatternsStore = defineStore('schemas/editor/patterns', () => {
  const editorStore = useEditorStore();
  const patterns = computed(() => editorStore.schema.content.patterns);
  const hasPatterns = computed(() => !!patterns.value.length);

  function addPattern(type: PatternType): void {
    editorStore.schema.content.patterns.push({
      id: ulid().toLowerCase(),
      type,
    });
  }

  return { hasPatterns, addPattern };
});
