<template>
  <SidebarSection>
    <TextField
      required
      variant="control"
      label="Назва Схеми"
      placeholder="Назва Схеми"
      v-model="schemaName"
    />
  </SidebarSection>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { TextField } from '@/components/form';
import { useEditorStore } from '../../../stores';
import SidebarSection from './SidebarSection.vue';

const editorStore = useEditorStore();

const schemaName = computed({
  get: () => editorStore.schema.name,
  set: (name) => editorStore.schema.name = name,
});
</script>
