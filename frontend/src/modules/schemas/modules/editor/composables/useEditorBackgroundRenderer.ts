import { computed, onUnmounted, watch } from 'vue';
import { useEditorStore } from '../stores';

export function useEditorBackgroundRenderer() {
  const editorStore = useEditorStore();
  const backgroundColor = computed(() => editorStore.schema.backgroundColor);

  watch(backgroundColor, (color) => {
    document.documentElement.style.setProperty('--editor-background-color', color);
  }, { immediate: true });

  onUnmounted(() => {
    document.documentElement.style.removeProperty('--editor-background-color');
  });
}
