<template>
  <li class="sidebar-pattern" :class="classes">
    {{ title }}
  </li>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import type { ISchemaPattern } from '@/models';
import { useEditorStore } from '@/modules/schemas/editor/stores';

const props = defineProps<{
  index: number;
  pattern: ISchemaPattern;
}>();

const editorStore = useEditorStore();

const classes = computed(() => ({
  'sidebar-pattern--active': editorStore.activePattern?.id === props.pattern.id,
}));

const title = computed(() => {
  const typeTitle = props.pattern.type === 'square' ? 'Квадратна Сітка' : 'Ромбова Сітка';
  return `${typeTitle} ${props.index + 1}`;
});
</script>

<style scoped>
@layer page {
  .sidebar-pattern {
    font-size: var(--font-sm);
    padding: 8px 12px;
    transition: background-color 0.15s ease-out;
    will-change: background-color;

    &:hover:not(.sidebar-pattern--active) {
      background-color: var(--color-background-2);
    }
  }

  .sidebar-pattern--active {
    background-color: var(--color-background-3);
  }
}
</style>
